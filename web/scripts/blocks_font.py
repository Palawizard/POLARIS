"""Builds src/fonts/polaris-blocks.woff2: the Unicode block elements
(U+2580-U+259F: █ ▓ ▒ ░ ▀ ▄ ▌ ▐ ...) drawn as exact cell rectangles, on
JetBrains Mono's metrics (600 x 1320 units). JetBrains Mono has no block
glyphs, so without this the game's title, the chocolatine and the other
block art would fall back to a random system font with gaps between cells.

Run: python3 scripts/blocks_font.py   (needs fonttools + brotli)
"""
from pathlib import Path

from fontTools.fontBuilder import FontBuilder
from fontTools.pens.ttGlyphPen import TTGlyphPen

UPM, W, ASC, DESC = 1000, 600, 1020, -300
H = ASC - DESC
OUT = Path(__file__).resolve().parent.parent / "src" / "fonts" / "polaris-blocks.woff2"


def rects_for(cp):
    """Rectangles (x0, y0, x1, y1) in cell units (0..1, y up) for one code point."""
    full = [(0, 0, 1, 1)]
    eighths_low = {0x2581: 1, 0x2582: 2, 0x2583: 3, 0x2584: 4, 0x2585: 5, 0x2586: 6, 0x2587: 7}
    eighths_left = {0x2589: 7, 0x258A: 6, 0x258B: 5, 0x258C: 4, 0x258D: 3, 0x258E: 2, 0x258F: 1}
    q = {  # quadrants: ul, ur, ll, lr
        "ul": (0, 0.5, 0.5, 1), "ur": (0.5, 0.5, 1, 1), "ll": (0, 0, 0.5, 0.5), "lr": (0.5, 0, 1, 0.5),
    }
    quads = {
        0x2596: ["ll"], 0x2597: ["lr"], 0x2598: ["ul"], 0x2599: ["ul", "ll", "lr"],
        0x259A: ["ul", "lr"], 0x259B: ["ul", "ur", "ll"], 0x259C: ["ul", "ur", "lr"],
        0x259D: ["ur"], 0x259E: ["ur", "ll"], 0x259F: ["ur", "ll", "lr"],
    }
    if cp == 0x2580:
        return [(0, 0.5, 1, 1)]
    if cp in eighths_low:
        return [(0, 0, 1, eighths_low[cp] / 8)]
    if cp == 0x2588:
        return full
    if cp in eighths_left:
        return [(0, 0, eighths_left[cp] / 8, 1)]
    if cp == 0x2590:
        return [(0.5, 0, 1, 1)]
    if cp in (0x2591, 0x2592, 0x2593):
        # Shades as an ordered dither of 4 x 8 dots per cell (25, 50, 75 %),
        # the same pattern scripts/logo.mjs prints in the title capture.
        test = {
            0x2591: lambda c, r: c % 2 == 0 and r % 2 == 0,
            0x2592: lambda c, r: (c + r) % 2 == 0,
            0x2593: lambda c, r: not (c % 2 == 1 and r % 2 == 1),
        }[cp]
        return [(c / 4, 1 - (r + 1) / 8, (c + 1) / 4, 1 - r / 8) for r in range(8) for c in range(4) if test(c, r)]
    if cp == 0x2594:
        return [(0, 7 / 8, 1, 1)]
    if cp == 0x2595:
        return [(7 / 8, 0, 1, 1)]
    if cp in quads:
        return [q[k] for k in quads[cp]]
    return []


BLEED = 60  # edges on the cell boundary overlap the neighbour a hair: no seams


def glyph(rects):
    pen = TTGlyphPen(None)
    for x0, y0, x1, y1 in rects:
        X0, X1 = round(x0 * W) - (BLEED if x0 == 0 else 0), round(x1 * W) + (BLEED if x1 == 1 else 0)
        Y0, Y1 = round(DESC + y0 * H) - (BLEED if y0 == 0 else 0), round(DESC + y1 * H) + (BLEED if y1 == 1 else 0)
        pen.moveTo((X0, Y0)); pen.lineTo((X0, Y1)); pen.lineTo((X1, Y1)); pen.lineTo((X1, Y0)); pen.closePath()
    return pen.glyph()


names = [".notdef", "space"]
cmap = {0x20: "space"}
glyphs = {".notdef": glyph([]), "space": glyph([])}
for cp in range(0x2580, 0x25A0):
    r = rects_for(cp)
    if not r:
        continue
    n = f"uni{cp:04X}"
    names.append(n)
    cmap[cp] = n
    glyphs[n] = glyph(r)

fb = FontBuilder(UPM, isTTF=True)
fb.setupGlyphOrder(names)
fb.setupCharacterMap(cmap)
fb.setupGlyf(glyphs)
fb.setupHorizontalMetrics({n: (W, 0) for n in names})
fb.setupHorizontalHeader(ascent=ASC, descent=DESC)
fb.setupOS2(version=4, sTypoAscender=ASC, sTypoDescender=DESC, sTypoLineGap=0, usWinAscent=ASC, usWinDescent=-DESC, fsSelection=0x80, xAvgCharWidth=W)
fb.setupNameTable({"familyName": "Polaris Blocks", "styleName": "Regular"})
fb.setupPost(isFixedPitch=1)
fb.font.flavor = "woff2"
OUT.parent.mkdir(parents=True, exist_ok=True)
fb.save(OUT)
print(f"{OUT.name}: {len(cmap) - 1} block glyphs, {OUT.stat().st_size} bytes")
