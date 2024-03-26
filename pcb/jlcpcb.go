package main

import (
	"archive/zip"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var (
	outputDir  = flag.String("dir", "production", "output directory")
	boardName  = flag.String("board", "mainboard", "board name")
	fixupsName = flag.String("fixup", "cpl_rotations_db.csv", "fixup file")
)

type Fixup struct {
	Rotation float64
	OffsetX  float64
	OffsetY  float64
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "jlcpcb: %v\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	if err := os.MkdirAll(*outputDir, 0o755); err != nil {
		return err
	}
	tmp, err := os.MkdirTemp("", "jlcpcb")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	verCmd := exec.Command("git", "describe", "--dirty", "--abbrev=10", "--tags", "--always")
	verBytes, err := verCmd.Output()
	if err != nil {
		return err
	}
	version := string(bytes.TrimSpace(verBytes))
	gerbersCmd := exec.Command(
		"kicad-cli", "pcb", "export", "gerbers",
		"-D", "VERSION="+version,
		"-o", tmp, "--subtract-soldermask",
		*boardName+".kicad_pcb",
	)
	gerbersCmd.Stderr = os.Stderr
	if err := gerbersCmd.Run(); err != nil {
		return err
	}
	drillCmd := exec.Command(
		"kicad-cli", "pcb", "export", "drill",
		"-o", tmp+string(filepath.Separator), // drill command expects trailing slash.
		"--excellon-oval-format",
		"--excellon-zeros-format", "decimal",
		"--drill-origin", "absolute",
		"--generate-map",
		"-u", "mm",
		*boardName+".kicad_pcb",
	)
	drillCmd.Stderr = os.Stderr
	if err := drillCmd.Run(); err != nil {
		return err
	}
	if err := zipDir(filepath.Join(*outputDir, *boardName+".zip"), tmp); err != nil {
		return err
	}

	bom := filepath.Join(tmp, "bom.csv")
	bomCmd := exec.Command(
		"kicad-cli", "sch", "export", "bom",
		"-o", bom, "--exclude-dnp", "--group-by", "LCSC", "--fields", "Reference,Value,Footprint,LCSC",
		*boardName+".kicad_sch",
	)
	bomCmd.Stderr = os.Stderr
	if err := bomCmd.Run(); err != nil {
		return err
	}
	dstBOM := filepath.Join(*outputDir, *boardName+"-bom.csv")
	srcBOM := filepath.Join(tmp, "bom.csv")
	fixups, err := convertBOM(dstBOM, srcBOM)
	if err != nil {
		return err
	}

	cpl := filepath.Join(tmp, "cpl.csv")
	cplCmd := exec.Command(
		"kicad-cli", "pcb", "export", "pos",
		"-o", cpl, "--format", "csv", "--exclude-dnp", "--units", "mm",
		*boardName+".kicad_pcb",
	)
	cplCmd.Stderr = os.Stderr
	if err := cplCmd.Run(); err != nil {
		return err
	}
	dstCPL := filepath.Join(*outputDir, *boardName+"-cpl.csv")
	srcCPL := filepath.Join(tmp, "cpl.csv")
	if err := convertCPL(fixups, dstCPL, srcCPL); err != nil {
		return err
	}

	return nil
}

func convertCPL(fixups map[string]Fixup, dstName, srcName string) (err error) {
	if err != nil {
		return err
	}
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := dst.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	srccsv := csv.NewReader(src)
	rec, err := srccsv.Read()
	if err != nil {
		return fmt.Errorf("cpl: %w", err)
	}
	if !reflect.DeepEqual(rec, []string{"Ref", "Val", "Package", "PosX", "PosY", "Rot", "Side"}) {
		return fmt.Errorf("cpl: unexpected header: %s", strings.Join(rec, ","))
	}
	dstcsv := csv.NewWriter(dst)
	dstcsv.Write([]string{"Designator", "Val", "Package", "Mid X", "Mid Y", "Rotation", "Layer"})
	for {
		rec, err := srccsv.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return fmt.Errorf("cpl: %w", err)
		}
		ref, val, pkg, posx, posy, rot, side := rec[0], rec[1], rec[2], rec[3], rec[4], rec[5], rec[6]
		var rotf, posxf, posyf float64
		var err1, err2, err3 error
		if rot != "" {
			rotf, err1 = strconv.ParseFloat(rot, 64)
		}
		if posx != "" {
			posxf, err2 = strconv.ParseFloat(posx, 64)
		}
		if posy != "" {
			posyf, err3 = strconv.ParseFloat(posy, 64)
		}
		line, _ := srccsv.FieldPos(1)
		if err1 != nil || err2 != nil || err3 != nil {
			return fmt.Errorf("cps: invalid position on line %d: %s, %s, %s", line, posx, posy, rot)
		}
		if fixup, ok := fixups[ref]; ok {
			// Compensate for backside placements.
			if side == "bottom" {
				rotf = 180 - rotf
			}
			rotf -= fixup.Rotation
			// Rotate fixup.
			offx, offy := fixup.OffsetX, fixup.OffsetY
			sinrot, cosrot := math.Sincos(rotf / 180 * math.Pi)
			offx, offy = cosrot*offx-sinrot*offy, sinrot*offx+cosrot*offy
			if side == "bottom" {
				offx = -offx
			}
			posxf -= offx
			posyf -= offy
		}
		dstcsv.Write([]string{
			ref, val, pkg,
			fmt.Sprintf("%f", posxf),
			fmt.Sprintf("%f", posyf),
			fmt.Sprintf("%f", rotf),
			side,
		})
	}
	dstcsv.Flush()
	return dstcsv.Error()
}

func convertBOM(dstName, srcName string) (fixups map[string]Fixup, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return nil, err
	}
	defer src.Close()
	dst, err := os.Create(dstName)
	if err != nil {
		return nil, err
	}
	defer func() {
		if cerr := dst.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	srccsv := csv.NewReader(src)
	rec, err := srccsv.Read()
	if err != nil {
		return nil, fmt.Errorf("bom: %w", err)
	}
	if !reflect.DeepEqual(rec, []string{"Reference", "Value", "Footprint", "LCSC"}) {
		return nil, fmt.Errorf("bom: unexpected header: %s", strings.Join(rec, ","))
	}
	dstcsv := csv.NewWriter(dst)
	dstcsv.Write([]string{"Comment", "Designator", "Footprint", "LCSC"})
	fixups = make(map[string]Fixup)
	for {
		rec, err := srccsv.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("bom: %w", err)
		}
		refs, val, footprint, partno := rec[0], rec[1], rec[2], rec[3]
		if val == "-- mixed values --" {
			val = "~"
		}
		if partno == "" {
			return nil, fmt.Errorf("%s: missing LCDC part number", refs)
		}
		refs, err = expandRanges(refs)
		if err != nil {
			return nil, err
		}
		if fixup, ok := JLCPCBFixups[partno]; ok {
			for _, ref := range strings.Split(refs, ",") {
				fixups[ref] = fixup
			}
		}
		dstcsv.Write([]string{val, refs, footprint, partno})
	}
	dstcsv.Flush()
	return fixups, dstcsv.Error()
}

// expandRanges takes a comma separated list of designators and
// expand ranges such as C1-C5.
func expandRanges(refs string) (string, error) {
	var split []string
	for _, refs := range strings.Split(refs, ",") {
		ref1, ref2, found := strings.Cut(refs, "-")
		if !found {
			split = append(split, refs)
			continue
		}
		des1, num1, err1 := splitDesignator(ref1)
		if err1 != nil {
			return "", err1
		}
		des2, num2, err2 := splitDesignator(ref2)
		if err2 != nil {
			return "", err1
		}
		if des1 != des2 {
			return "", fmt.Errorf("invalid designator: %s", refs)
		}
		for i := num1; i <= num2; i++ {
			split = append(split, fmt.Sprintf("%s%d", des1, i))
		}
	}
	return strings.Join(split, ","), nil
}

func splitDesignator(ref string) (string, int, error) {
	for n, c := range ref {
		if unicode.IsNumber(c) {
			num, err := strconv.Atoi(ref[n:])
			if err != nil {
				return "", 0, fmt.Errorf("invalid designator: %s", ref)
			}
			return ref[:n], num, nil
		}
	}
	return "", 0, fmt.Errorf("invalid designator: %s", ref)
}

func zipDir(zipName, dir string) error {
	files, err := os.ReadDir(dir)
	if err != nil {
		return err
	}
	zipf, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer func() {
		if cerr := zipf.Close(); cerr != nil && err == nil {
			err = cerr
		}
	}()
	w := zip.NewWriter(zipf)
	defer w.Close()

	for _, f := range files {
		entry, err := w.Create(f.Name())
		if err != nil {
			return err
		}
		file, err := os.Open(filepath.Join(dir, f.Name()))
		if err != nil {
			return err
		}
		_, err = io.Copy(entry, file)
		if err != nil {
			return err
		}
		file.Close()
	}
	return nil
}
