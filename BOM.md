# SeedHammer II — Bill of Materials

> **Status: DRAFT.** This BOM was compiled by the community from the files in this
> repository (CAD STEP files, `specs/hammerhead.pdf`, `cad/nema17-stepper.pdf`, the
> KiCad design in `pcb/`, and the [firmware](https://github.com/seedhammer/seedhammer)
> constants). Items marked ⚠️ are inferred and need confirmation by the SeedHammer
> team — corrections and exact part numbers are very welcome.

## Machine overview

Two leadscrew-driven axes, no Z: the engraver ("hammerhead") travels on an upper X
gantry, the plate clamps into a sled travelling on Y beneath it. Depth comes from the
solenoid head striking at up to 40 Hz. Homing is sensorless (TMC2209 StallGuard) — no
endstop switches. Approximate frame footprint 240 × 240 × 200 mm.

## 1. Motion

| Part Name | Part Number | Description / Spec | Qty | Vendor(s) |
|---|---|---|---|---|
| Stepper motor, NEMA 17 | `42DND0155B-22` (Kaihong) | 42 mm frame, 40 mm body, 1.8°/step, 1.65 A, 2.72 V, ≥0.4 N·m, Ø5 D-shaft ×22 mm, JST PH-6P connector. Spec: `cad/nema17-stepper.pdf`. Generic 17HS4401-class motors are near-identical substitutes. | 2 | [StepperOnline UK — 17HS16-2004S1](https://www.stepperonline.co.uk/nema-17-bipolar-45ncm-63-74oz-in-2a-42x42x40mm-4-wires-w-1m-cable-connector-17hs16-2004s1.html) (2 A/45 N·cm, 40 mm, ships with 1 m cable) · [AliExpress — 17HS4401](https://www.aliexpress.com/item/3256802822821809.html) |
| Leadscrew | Tr8×8 (T8, 4-start, 8 mm lead) | Firmware: 8 mm travel/rev (`mmPerRevolution = 8`). ⚠️ Length ~175 mm inferred from ~186 mm frame width — cut 200 mm stock to length. | 2 | [Amazon US — Kwweeoo 200 mm Tr8×8, 2 pcs](https://www.amazon.com/Length-Diameter-Thread-Printer-Kwweeoo/dp/B0CK29R5R4) · [Amazon US — FKG 200 mm T8×8 + nut](https://www.amazon.com/FKG-200mm-Lead-Screw-Brass/dp/B08JP2ZPWF) |
| Leadscrew nut, brass | Tr8×8 | Brass nut per axis (hex-style nut visible in CAD render). ⚠️ Exact style (hex vs flange vs anti-backlash) to confirm. | 2 | [Amazon US — T8×8 brass flange nut](https://www.amazon.com/Screw-T8x8mm-Flange-Printer-Accessories/dp/B089G5C8TN) |
| Guide rail, smooth | Ø14 h6 linear shaft | ⚠️ Ø14 inferred from Ø14.03 ±0.05 bores in `cad/clampsled.pdf` and Ø14.05 in `y-sled-bottom.step`; sleds appear to ride the shafts directly (plain bore — no linear ball bearings found in CAD). Hardened + hard-chromed, ~175 mm, 2 per axis. | 4 | [Motedis EU — Ø14 h6 hardened+ground, cut to length](https://www.motedis.com/en/Precision-shaft-14-mm-h6-steel-hardened-and-ground) · [Motedis USA](https://www.motedis-usa.com/en/Precision-shaft-14-mm-h6-steel-hardened-and-ground) · [Misumi — precision linear shaft](https://us.misumi-ec.com/vona2/detail/110302634310/) · [McMaster — metric precision shafts](https://www.mcmaster.com/products/metric-steel-precision-shafts) |
| Bearing, 608 | 608-2RS / 608ZZ | 8×22×7 mm. ⚠️ Implied by `cad/608-tool-grip.step` / `608-tool-tip.step` — usage (leadscrew end support vs. assembly tool) and quantity to confirm. | 2+ | [Amazon UK — NSK 608-2RS ×10](https://www.amazon.co.uk/NSK-Precision-Bearings-Pre-Lubricated-Bearings/dp/B084LZYZWQ) · [Amazon UK — generic 608-2RS ×10](https://www.amazon.co.uk/608-2RS-Ball-Bearings-Skateboard-Skateboards/dp/B07VFTQBPB) |

## 2. Engraver ("hammerhead")

| Part Name | Part Number | Description / Spec | Qty | Vendor(s) |
|---|---|---|---|---|
| Solenoid impact head | — (OEM, per `specs/hammerhead.pdf`) | Ø40 × 71 mm, M26 threaded neck, 90° carbide tip, 16 cm cable → JST XH 2-pin female. Driven 20–28 V DC, 25 ms cycle (≈40 Hz), 4–5 ms on-time. ⚠️ The drawing is an interface spec for an OEM commodity marking-machine head — **exact OEM/source unknown; input from the SeedHammer team wanted.** DIY fallback: harvest the head from a 400 W desktop electric marking machine and verify the coil at 20–28 V pulsed drive. | 1 | Donor machine: [Amazon US — 400 W desktop marking machine](https://www.amazon.com/Electric-Marking-Engraving-Nameplate-Industrial/dp/B0CQM15DMF) (same head family; VEVOR sell equivalents) |
| Spare carbide tips | 90° included angle | Consumable dot-peen/marking styli, tungsten carbide. | n | [Amazon US](https://www.amazon.com/3mm-Pneumatic-Tailstock-Dia-10-7mm-Engraving/dp/B09M7VHW26) · [Alibaba (Zixu)](https://www.alibaba.com/product-detail/ZIXU-Manufacture-Price-6MM-Carbide-Pins_60253271680.html) · [2L Inc (US)](https://www.2linc.com/product/dot-peen-marking-toolbit/) |

## 3. Electronics

| Part Name | Part Number | Description / Spec | Qty | Vendor(s) |
|---|---|---|---|---|
| Mainboard | `pcb/mainboard.kicad_pcb` | Custom 4-layer PCB — RP2350, 3× TMC2209-LA, ST25R3916 NFC, AP33772S USB-PD sink, W25Q128 flash, 5× 2200 µF bulk. Fab settings (README): ENIG, FR4 TG155, stackup JLC04161H-7628. Gerber/BOM/CPL generate via `cd pcb && go run .` for assembly. | 1 | [JLCPCB](https://jlcpcb.com) (fab + SMT assembly) |
| LCD + touch panel | `ER-TFT035IPS-6` | 3.5″ IPS 320×480, **capacitive** touch (FT6x36), ILI9488, 50-pin 0.5 mm FFC. Resistive variant is NOT supported by firmware. | 1 | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) |
| Power supply | USB-C PD 3.1 | 100 W minimum, **140 W (28 V EPR) recommended**; single-port preferred. | 1 | [Anker 717 — 140 W single-port (A2341)](https://www.anker.com/products/a2341) · [Anker EU](https://www.anker.com/eu-en/products/a2341) · [Anker 140 W 4-port (A2697)](https://www.anker.com/products/a2697-anker-charger-140w-4-port) |
| USB-C cable | 240 W / EPR-rated | Must be EPR e-marked or the supply falls back to 100 W / 20 V. | 1 | [Anker 240 W USB-C cable](https://www.anker.com/products/a82e2-240w-usb-c-to-usb-c-cable) · [Amazon US](https://www.amazon.com/Anker-Bio-Nylon-Charging-MacBook-Samsung/dp/B0BHQH89ZQ) |

## 4. Cables & connectors

| Part Name | Part Number | Description / Spec | Qty | Vendor(s) |
|---|---|---|---|---|
| Motor cable | JST PH 2.0 6-pin → JST XH 2.5 4-pin | Motor side PH-6P (pins 1/3/4/6 = A+/A−/B−/B+), board side XH 4-pin (`B4B-XH-A`). Standard "PH6–XH4 stepper cable" or crimp your own. ⚠️ Length TBD (~30 cm). | 2 | [Amazon US — PH 2.0 pre-crimped kit](https://www.amazon.com/Connector-Pre-Crimped-Compatible-JST-PH2-0mm-Inductrix/dp/B08T89ZK2Q) · [Adafruit 5090 — PH 6-pin cable pair](https://www.adafruit.com/product/5090) · [Pololu — 6-pin PH cables](https://www.pololu.com/category/361/6-pin-jst-ph-style-cables) |
| Engraver connection | JST XH 2-pin | Head's captive 16 cm lead mates with board socket `B2B-XH-A`. | — | (part of head) |
| LCD ribbon | 50-pin FFC, 0.5 mm pitch | Board connector TE `5-1734839-0`. ⚠️ Length/orientation TBD — often supplied with the buydisplay panel. | 1 | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) |
| JST crimp kit | XH + PH assortment | For making motor/aux leads (plus 2× SH 1.0 3-pin debug wafers on board). | 1 | [Amazon UK — HALJIA 560-pc JST-XH kit](https://www.amazon.co.uk/HALJIA-560Pcs-2-54mm-JST-XH-Connector/dp/B074FS9M12) |

## 5. Custom fabricated parts (STEP files in `cad/`)

| Group | Files | Process |
|---|---|---|
| Frame | `left`, `right`, `front`, `back` | CNC aluminium (≈186 mm wide × 237 mm deep plates) |
| X axis | `clampsled` (+ `clampsled.pdf`), `x-house`, `x-motor-hodlr`, `M4-cnurled-Nut` | CNC |
| Y axis / plate clamp | `y-sled-top`, `y-sled-bottom`, `y-motor-hodlr`, `jaw`, `lever`, `centerpiece-lever` | CNC |
| PCB mounting | `pcb-clip`, `pcb-nipple` | 3D print |
| Misc | `faceplate` (3MF — 3D print), `foot` ×4, `motor-cable-clip` ×2 + M3 variant, `608-tool-grip`/`608-tool-tip`, `parametric-sh-plate-box` | 3D print / CNC |

CNC services: [JLCCNC](https://jlccnc.com) · [PCBWay CNC](https://www.pcbway.com/rapid-prototyping/manufacture/) · [Xometry EU](https://xometry.eu).

## 6. Fasteners & consumables

| Part Name | Description / Spec | Qty | Vendor(s) |
|---|---|---|---|
| Screw assortment | M3 + M4 socket-head (motor mounts 4× M3 each, ≥4.5 mm engagement; clamp M4×0.7; frame M3/M4 per STEP threads). ⚠️ Exact list TBD. | 1 kit | [Amazon UK — VIGRUE 1080-pc M2/M3/M4 socket-head kit](https://www.amazon.co.uk/VIGRUE-Assortment-Washers-Stainless-Storage/dp/B081SGZ2C4) |
| Steel plates | 85 × 85 mm, 316 marine-grade stainless (SeedHammer II / SH02 pattern) | n | [seedhammer.com/shop](https://seedhammer.com/shop) |

## NIP-99 / bitcoin-native sourcing

No NIP-99 classified listings were found for these parts yet. Given the project's
bitcoin-only ethos, listings on [shopstr.store](https://shopstr.store),
[plebeian.market](https://plebeian.market) or [conduit.market](https://conduit.market)
(sats-priced motors, rails, heads — or complete kits) would be a great fit; sellers
welcome.

## Open questions for the SeedHammer team

- [ ] Guide rails: confirm Ø14, material/finish, and exact lengths per axis
- [ ] Sled bearing arrangement: plain machined bore directly on shaft, or pressed bushings?
- [ ] Leadscrew: confirmed Tr8×8? Length per axis? Nut part number?
- [ ] 608 bearings: where used, and how many?
- [ ] Hammerhead: OEM/supplier or purchasable equivalent?
- [ ] Frame parts: material/alloy and finish (anodising?)
- [ ] Full fastener list with counts
- [ ] FFC length/type for the LCD
