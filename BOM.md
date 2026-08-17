# SeedHammer II — Bill of Materials

> **Status: DRAFT.** This BOM was compiled by the community from the files in this
> repository (CAD STEP files, `specs/hammerhead.pdf`, `cad/nema17-stepper.pdf`, the
> KiCad design in `pcb/`, and the [firmware](https://github.com/seedhammer/seedhammer)
> constants). Items marked ⚠️ are inferred and need confirmation by the SeedHammer
> team — corrections and exact part numbers are very welcome. All vendor links were
> link-checked at the time of writing; "global" means the vendor ships worldwide.

## Machine overview

Two leadscrew-driven axes, no Z: the engraver ("hammerhead") travels on an upper X
gantry, the plate clamps into a sled travelling on Y beneath it. Depth comes from the
solenoid head striking at up to 40 Hz. Homing is sensorless (TMC2209 StallGuard) — no
endstop switches. Approximate frame footprint 240 × 240 × 200 mm.

Published machine photos suggest an unusual drive layout: the motor rides on the
moving carriage and the leadscrew travels with it, threading through a **brass hex
nut fixed to the frame** (visible on both axes). The Tr8 spindle rotates in 608
bearings (per the team — use good ones).

## 1. Motion

| Part Name | Part Number | Description / Spec | Qty | UK | US | EU |
|---|---|---|---|---|---|---|
| Stepper motor, NEMA 17 | `42DND0155B-22` (Kaihong) | 42 mm frame, 40 mm body, 1.8°/step, 1.65 A, 2.72 V, ≥0.4 N·m, Ø5 D-shaft ×22 mm, JST PH-6P connector. Spec: `cad/nema17-stepper.pdf`. Generic 17HS4401-class motors are near-identical substitutes. | 2 | [StepperOnline UK — 17HS16-2004S1](https://www.stepperonline.co.uk/nema-17-bipolar-45ncm-63-74oz-in-2a-42x42x40mm-4-wires-w-1m-cable-connector-17hs16-2004s1.html) | [StepperOnline — 17HS16-2004S1](https://www.omc-stepperonline.com/nema-17-bipolar-45ncm-64oz-in-2a-42x42x40mm-4-wires-w-1m-cable-connector-17hs16-2004s1) | [AliExpress — 17HS4401](https://www.aliexpress.com/item/3256802822821809.html) (global) |
| Leadscrew | Tr8×8 (T8, 4-start, 8 mm lead) | Firmware: 8 mm travel/rev (`mmPerRevolution = 8`). ⚠️ Length ~175 mm inferred from ~186 mm frame width — cut 200 mm stock to length. | 2 | [Amazon UK — Kwweeoo 200 mm Tr8×8, 2 pcs](https://www.amazon.co.uk/dp/B0CK29R5R4) | [Amazon US — Kwweeoo](https://www.amazon.com/dp/B0CK29R5R4) · [Amazon US — FKG + nut](https://www.amazon.com/FKG-200mm-Lead-Screw-Brass/dp/B08JP2ZPWF) | [Motedis — TR8×8 spindle, cut to length](https://www.motedis.com/en/Trapezoidal-thread-spindle-TR8x8) |
| Leadscrew nut, brass | Tr8×8 | Brass **hex** nut per axis, fixed to the frame (confirmed in product photos — motor and screw travel through it). Team confirm Tr8×8 and that they buy the brass nuts on AliExpress — hex, OD 13 mm (⚠️ confirm pocket size). | 2 | [AliExpress — brass hex nut Tr8×8, OD 13 mm](https://www.aliexpress.com/item/1005006783604505.html) (global) | [Amazon US — T8×8 brass flange nut](https://www.amazon.com/Screw-T8x8mm-Flange-Printer-Accessories/dp/B089G5C8TN) | [Motedis — spindle accessories](https://www.motedis.com/en/Trapezoidal-threaded-spindle-and-accessories-Shop) |
| Guide rail, smooth | Ø12 h6 linear shaft | **Ø12, confirmed by the team.** Hardened + ground, ~175 mm, 2 per axis; ⚠️ exact lengths TBC. (The Ø14.03/Ø14.05 sled bores are the bushing seats.) | 4 | [Motedis — Ø12 h6 hardened+ground, cut to length](https://www.motedis.com/en/Precision-shaft-12-mm-h6-steel-hardened-and-ground) (ships UK) | [Motedis USA](https://www.motedis-usa.com/en/Precision-shaft-12-mm-h6-steel-hardened-and-ground) · [Misumi](https://us.misumi-ec.com/vona2/detail/110302634310/) · [McMaster](https://www.mcmaster.com/products/metric-steel-precision-shafts) | [Motedis](https://www.motedis.com/en/Precision-shaft-12-mm-h6-steel-hardened-and-ground) |
| Sleeve bushing | 12 mm ID × 14 mm OD (e.g. igus iglidur J `JSM-1214-15`) | **Bushings in the sleds, confirmed by the team** — press-fit into the Ø14.03 ±0.05 seats. ⚠️ Exact part/material (polymer vs bronze) and length TBC. | 4+ | [RS UK — JSM-1214-15, ~£1.05](https://uk.rs-online.com/web/p/plain-bearings/2724937) · [Misumi UK](https://uk.misumi-ec.com/vona2/detail/221000103504/?HissuCode=JSM-1214-15) | [igus — iglidur sleeve bearings](https://www.igus.eu/iglidur-ibh/sleeve-bearings) | [igus](https://www.igus.eu/iglidur-ibh/sleeve-bearings) · [AliExpress — SF-1 composite 12×14×15](https://www.aliexpress.com/item/1005009318987447.html) (global) |
| Bearing, 608 | 608-2RS / 608ZZ | 8×22×7 mm. **Used to rotate the Tr8 spindle (confirmed by the team) — use high-quality bearings: NSK or SKF.** ⚠️ Count per axis TBC. (`608-tool-grip`/`608-tool-tip` are the fitting tools.) | 2+ | [Amazon UK — NSK 608-2RS ×10](https://www.amazon.co.uk/dp/B084LZYZWQ) | [Amazon US — NSK 608-2RS ×10](https://www.amazon.com/dp/B084LZYZWQ) · [Amazon US — ANCIRS ×20](https://www.amazon.com/dp/B07C6FL8TW) | [Amazon DE — NSK 608-2RS ×10](https://www.amazon.de/dp/B084LZYZWQ) |

## 2. Engraver ("hammerhead")

| Part Name | Part Number | Description / Spec | Qty | UK | US | EU |
|---|---|---|---|---|---|---|
| Solenoid impact head | — (OEM, per `specs/hammerhead.pdf`) | Ø40 × 71 mm, M26 threaded neck, 90° carbide tip, 16 cm cable → JST XH 2-pin female. Driven 20–28 V DC, 25 ms cycle (≈40 Hz), 4–5 ms on-time. ⚠️ The drawing is an interface spec for an OEM commodity marking-machine head. The SeedHammer team report the OEM is unknown even to them (the original supplier withheld manufacturer details and later ceased trading) and a redesign/reverse-engineering of the head is planned; Alibaba/AliExpress listings surface intermittently for single buys — image search recommended. DIY fallback: harvest the head from a 400 W desktop electric marking machine and verify the coil at 20–28 V pulsed drive. | 1 | Candidate spare head (global): [AliExpress — "electromechanical marking head", ~£59](https://www.aliexpress.com/item/1005012679957256.html) — **the team confirm it looks right**; still verify Ø40/M26/coil voltage with the seller | Donor machine: [Amazon US — 400 W desktop marking machine](https://www.amazon.com/Electric-Marking-Engraving-Nameplate-Industrial/dp/B0CQM15DMF) · [eBay — "Electromagnetic Marking Head"](https://www.ebay.com/itm/394579043426) (listing may expire) | Same AliExpress candidate (ships EU) — or quote the [`hammerhead.pdf`](specs/hammerhead.pdf) drawing to a marking-machine OEM on Alibaba |
| Spare carbide tips | 90° included angle | Consumable dot-peen/marking styli, tungsten carbide. | n | [Amazon UK](https://www.amazon.co.uk/dp/B09M7VHW26) | [Amazon US](https://www.amazon.com/dp/B09M7VHW26) · [2L Inc](https://www.2linc.com/product/dot-peen-marking-toolbit/) | [Amazon DE](https://www.amazon.de/dp/B09M7VHW26) · [AliExpress — tungsten needle + brass bushing consumables](https://www.aliexpress.com/item/1005007912316101.html) · [Alibaba (Zixu)](https://www.alibaba.com/product-detail/ZIXU-Manufacture-Price-6MM-Carbide-Pins_60253271680.html) (global) |

## 3. Electronics

| Part Name | Part Number | Description / Spec | Qty | UK | US | EU |
|---|---|---|---|---|---|---|
| Mainboard | `pcb/mainboard.kicad_pcb` | Custom 4-layer PCB — RP2350, 3× TMC2209-LA, ST25R3916 NFC, AP33772S USB-PD sink, W25Q128 flash, 5× 2200 µF bulk. Fab settings (README): ENIG, FR4 TG155, stackup JLC04161H-7628. Gerber/BOM/CPL generate via `cd pcb && go run .` for assembly. | 1 | [JLCPCB](https://jlcpcb.com) (global) | [JLCPCB](https://jlcpcb.com) (global) | [JLCPCB](https://jlcpcb.com) (global) |
| LCD + touch panel | `ER-TFT035IPS-6` | 3.5″ IPS 320×480, **capacitive** touch (FT6x36), ILI9488, 50-pin 0.5 mm FFC. Resistive variant is NOT supported by firmware. | 1 | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) (global) | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) (global) | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) (global) |
| Power supply | USB-C PD 3.1 | 100 W minimum, **140 W (28 V EPR) recommended**; single-port preferred. UK models listed have UK 3-pin plugs. | 1 | [Amazon UK — UGREEN Nexode 140 W](https://www.amazon.co.uk/dp/B09YQ374LF) · [Amazon UK — Anker 140 W 4-port](https://www.amazon.co.uk/dp/B0DPKP7RD8) · [Anker UK store](https://www.anker.com/uk/products/a2697-anker-charger-140w-4-port) | [Anker 717 (A2341)](https://www.anker.com/products/a2341) · [Amazon US — UGREEN Nexode 140 W](https://www.amazon.com/dp/B0B129DM9T) | [Anker 717 — EU store](https://www.anker.com/eu-en/products/a2341) · [Amazon DE — UGREEN Nexode 140 W](https://www.amazon.de/dp/B09YQ374LF) |
| USB-C cable | 240 W / EPR-rated | Must be EPR e-marked or the supply falls back to 100 W / 20 V. | 1 | [Amazon UK — Anker 240 W](https://www.amazon.co.uk/dp/B0BHQH89ZQ) | [Anker 240 W cable](https://www.anker.com/products/a82e2-240w-usb-c-to-usb-c-cable) · [Amazon US](https://www.amazon.com/Anker-Bio-Nylon-Charging-MacBook-Samsung/dp/B0BHQH89ZQ) | [Amazon DE — Anker 240 W](https://www.amazon.de/dp/B0BHQH89ZQ) |

## 4. Cables & connectors

| Part Name | Part Number | Description / Spec | Qty | UK | US | EU |
|---|---|---|---|---|---|---|
| Motor cable | JST PH 2.0 6-pin → JST XH 2.5 4-pin | Motor side PH-6P (pins 1/3/4/6 = A+/A−/B−/B+), board side XH 4-pin (`B4B-XH-A`). Standard "PH6–XH4 stepper cable" or crimp your own. ⚠️ Length TBD (~30 cm). | 2 | [Amazon UK — PH 2.0 pre-crimped kit](https://www.amazon.co.uk/dp/B08T89ZK2Q) | [Adafruit 5090 — PH 6-pin pair](https://www.adafruit.com/product/5090) · [Pololu — 6-pin PH cables](https://www.pololu.com/category/361/6-pin-jst-ph-style-cables) | [Pololu](https://www.pololu.com/category/361/6-pin-jst-ph-style-cables) (ships worldwide) |
| Engraver connection | JST XH 2-pin | Head's captive 16 cm lead mates with board socket `B2B-XH-A`. | — | (part of head) | (part of head) | (part of head) |
| LCD ribbon | 50-pin FFC, 0.5 mm pitch | Board connector TE `5-1734839-0`. ⚠️ Length/orientation TBD — often supplied with the buydisplay panel. | 1 | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) (global) | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) (global) | [buydisplay.com](https://www.buydisplay.com/3-5-inch-ips-320x480-tft-lcd-display-capacitive-touch-screen) (global) |
| JST crimp kit | XH assortment | For making motor/aux leads (plus 2× SH 1.0 3-pin debug wafers on board). | 1 | [Amazon UK — HALJIA 560-pc JST-XH kit](https://www.amazon.co.uk/dp/B074FS9M12) | [Amazon US — HALJIA](https://www.amazon.com/dp/B074FS9M12) · [Amazon US — Taiss 560-pc](https://www.amazon.com/dp/B09ZTWCZ3K) | [Amazon DE — HALJIA](https://www.amazon.de/dp/B074FS9M12) |

## 5. Custom fabricated parts (STEP files in `cad/`)

Processes confirmed by the SeedHammer team: **black parts are 3D-printed PETG**
(including the frame), **orange parts are ASA** (temperature near the hammerhead),
and the **carriage (clampsled) moved to CNC** in production — printing it works for
occasional single-plate use with cooling time between plates, in ASA or glass-fibre
reinforced PA.

| Group | Files | Process (inferred) |
|---|---|---|
| Frame | [`left`](cad/left.step), [`right`](cad/right.step), [`front`](cad/front.step), [`back`](cad/back.step) | 3D-printed, PETG, black (confirmed); ≈186 mm wide × 237 mm deep |
| X axis | [`clampsled-raw`](cad/clampsled-raw.step) (+ [machining drawing](cad/clampsled.pdf) — no separate final STEP; the drawing carries the tolerances/threads), [`x-house`](cad/x-house.step), [`x-motor-hodlr`](cad/x-motor-hodlr.step) | CNC in production (drawing carries Ø14.03 ±0.05 / M4×0.7 6H callouts); print in ASA or PA-GF for light use |
| Y axis / plate clamp | [`y-sled-top`](cad/y-sled-top.step), [`y-sled-bottom`](cad/y-sled-bottom.step), [`y-motor-hodlr`](cad/y-motor-hodlr.step), [`jaw`](cad/jaw.step), [`lever`](cad/lever.step), [`centerpiece-lever`](cad/centerpiece-lever.step) | 3D-printed — orange parts in ASA, black in PETG (confirmed) |
| PCB mounting | [`pcb-clip`](cad/pcb-clip.step), [`pcb-nipple`](cad/pcb-nipple.step) | 3D-printed, PETG (confirmed) |
| Likely off-the-shelf | [`foot`](cad/foot.step) ×4 (rubber foot), [`M4-cnurled-Nut`](cad/M4-cnurled-Nut.step) (knurled brass finger screw, visible in machine photos) | Buy, not fabricate |
| Accessories / tooling | [`faceplate`](cad/faceplate.3mf) (3MF — 3D print), [`motor-cable-clip`](cad/motor-cable-clip.step) ([×2](cad/motor-cable-clip-2x.step), [M3 variant](cad/motor%20cable%20clip%20M3.step)), [`608-tool-grip`](cad/608-tool-grip.step)/[`608-tool-tip`](cad/608-tool-tip.step), [`parametric-sh-plate-box`](cad/parametric-sh-plate-box.step) ([f3d](cad/parametric-sh-plate-box.f3d)) | 3D print |

Machining/CNC services (all global): [JLCCNC](https://jlccnc.com) · [PCBWay CNC](https://www.pcbway.com/rapid-prototyping/manufacture/) · [Xometry EU](https://xometry.eu).

## 6. Fasteners & consumables

| Part Name | Description / Spec | Qty | UK | US | EU |
|---|---|---|---|---|---|
| Screw assortment | M3 + M4 socket-head (motor mounts 4× M3 each, ≥4.5 mm engagement; clamp M4×0.7; frame M3/M4 per STEP threads). ⚠️ Exact list TBD. | 1 kit | [Amazon UK — VIGRUE 1080-pc M2/M3/M4](https://www.amazon.co.uk/dp/B081SGZ2C4) | [Amazon US — VIGRUE 1290-pc](https://www.amazon.com/dp/B071KBVZVV) | [Amazon DE — VIGRUE 1080-pc](https://www.amazon.de/dp/B081SGZ2C4) |
| Steel plates | 85 × 85 mm; 2 mm and 3 mm both work, **3 mm recommended by the team**; laser-cut; material **1.4404 (316L) or 1.4571 (316Ti), warm-rolled preferred** | n | [seedhammer.com/shop](https://seedhammer.com/shop) (global) · [LNbits Shop](https://shop.lnbits.com/product/seedhammer-steel-plates) (global) · cut-to-size: [Lasered UK](https://www.lasered.co.uk) | [MineRacks — plate kits](https://www.mineracks.com/shop/p/seedhammer-steel-plate-back-up-kit-for-2-of-3-multisig-sx36a-zpsll) (3 mm — the team-recommended thickness) · cut-to-size: [SendCutSend](https://sendcutsend.com/materials/stainless-steel/) · [Cut2Size](https://cut2sizemetals.com/products/stainless-steel-sheet/) | [LNbits Shop](https://shop.lnbits.com/product/seedhammer-steel-plates) · cut-to-size: [Fractory](https://fractory.com/online-laser-cutting/) |

## NIP-99 / bitcoin-native sourcing

No NIP-99 classified listings were found for these parts yet. Given the project's
bitcoin-only ethos, listings on [shopstr.store](https://shopstr.store),
[plebeian.market](https://plebeian.market) or [conduit.market](https://conduit.market)
(sats-priced motors, rails, heads — or complete kits) would be a great fit; sellers
welcome.

## Answers from the SeedHammer team (2026-08-14)

Via the SeedHammer Community group (thanks Nick):

1. The AliExpress spare-head candidate **looks right**; OEM still unknown (original
   supplier withheld details and ceased trading) — a head redesign is planned, and
   findings on the OEM are welcome.
2. Guide rails are **Ø12 with bushings in the sleds**.
3. Carriage **moved to CNC**; printing works for single plates with cooling time —
   use ASA or glass-fibre-reinforced PA.
4. Leadscrew **Tr8×8 confirmed**; brass nuts bought on AliExpress.
5. 608 bearings **rotate the Tr8 spindle** — use high quality (NSK or SKF).
6. Frame and black parts printed in **PETG**; orange parts **ASA** (temperature).
7. Fastener list: the team will share one.
8. Display ships with what's needed — order the **capacitive touch** variant.
9. Plates: 2 mm and 3 mm both work, **3 mm recommended**; laser-cut;
   **1.4404 or 1.4571, warm-rolled preferred**.

## Remaining open questions

- [ ] Exact rail and leadscrew lengths per axis
- [ ] Bushing part number/material (polymer like igus JSM-1214, or bronze?) and length
- [ ] 608 bearing count and arrangement per axis
- [ ] Fastener list with counts (promised)
- [ ] Confirmation of the moving-motor / frame-fixed-nut drive layout reading
