---
name: Polaris (browser edition)
description: The back of Polaris's red game box, whose big screenshot is the real terminal game running.
colors:
  box-red: "#b5202a"
  box-red-deep: "#8a1520"
  box-red-dark: "#6e0f19"
  print-cream: "#fff1dc"
  soft-cream: "#ffd9c8"
  ink-red-brown: "#2b0a0d"
  screen: "#180d11"
  screen-print: "#21141a"
  terminal-fg: "#f6e8d6"
typography:
  display:
    fontFamily: "Big Shoulders Display, Arial Narrow, sans-serif"
    fontSize: "clamp(2rem, 4.4vw, 2.75rem)"
    fontWeight: 900
    lineHeight: 1
    letterSpacing: "0.02em"
  headline:
    fontFamily: "Big Shoulders Display, Arial Narrow, sans-serif"
    fontSize: "clamp(2.2rem, 4vw, 3rem)"
    fontWeight: 900
    lineHeight: 0.95
  title:
    fontFamily: "Big Shoulders Display, Arial Narrow, sans-serif"
    fontSize: "1.55rem"
    fontWeight: 900
    lineHeight: 1
    letterSpacing: "0.01em"
  genre:
    fontFamily: "Big Shoulders Display, Arial Narrow, sans-serif"
    fontSize: "clamp(0.95rem, 1.4vw, 1.2rem)"
    fontWeight: 800
    lineHeight: 1
    letterSpacing: "0.08em"
  key-numeral:
    fontFamily: "Big Shoulders Display, Arial Narrow, sans-serif"
    fontSize: "1.45rem"
    fontWeight: 900
    lineHeight: 1
    fontFeature: "tnum"
  body:
    fontFamily: "Figtree, ui-sans-serif, system-ui, sans-serif"
    fontSize: "1rem"
    fontWeight: 500
    lineHeight: 1.5
  caption:
    fontFamily: "Figtree, ui-sans-serif, system-ui, sans-serif"
    fontSize: "0.95rem"
    fontWeight: 700
    lineHeight: 1.3
  label:
    fontFamily: "Figtree, ui-sans-serif, system-ui, sans-serif"
    fontSize: "0.8rem"
    fontWeight: 800
    lineHeight: 1.2
    letterSpacing: "0.06em"
  small-print:
    fontFamily: "Figtree, ui-sans-serif, system-ui, sans-serif"
    fontSize: "0.82rem"
    fontWeight: 500
    lineHeight: 1.5
  terminal:
    fontFamily: "JetBrains Mono, Polaris Blocks, ui-monospace, Cascadia Mono, Menlo, Consolas, monospace"
    fontSize: "15px"
    fontWeight: 400
    lineHeight: 1
    letterSpacing: "0"
  vignette:
    fontFamily: "JetBrains Mono, Polaris Blocks, ui-monospace, Cascadia Mono, Menlo, Consolas, monospace"
    fontSize: "0.8rem"
    fontWeight: 400
    lineHeight: 1.08
rounded:
  swatch: "2px"
  label: "6px"
  vignette: "9px"
  key: "10px"
  frame: "14px"
spacing:
  gutter: "clamp(14px, 3vw, 36px)"
  key-gap: "0.5rem"
  column-gap: "clamp(1.25rem, 2.4vw, 2.25rem)"
  section: "clamp(2.5rem, 6vw, 4rem)"
components:
  tab:
    backgroundColor: "{colors.print-cream}"
    textColor: "{colors.ink-red-brown}"
    rounded: "{rounded.key}"
    height: "2.5rem"
    padding: "0 0.95rem 0 0.6rem"
  key:
    backgroundColor: "{colors.print-cream}"
    textColor: "{colors.ink-red-brown}"
    rounded: "{rounded.key}"
    size: "2.5rem"
  sticker:
    backgroundColor: "{colors.print-cream}"
    textColor: "{colors.box-red}"
    typography: "{typography.display}"
    rounded: "{rounded.frame}"
    height: "4rem"
    padding: "0.35rem 1.7rem 0.45rem 1.35rem"
  screenshot-frame:
    backgroundColor: "{colors.screen}"
    textColor: "{colors.terminal-fg}"
    typography: "{typography.terminal}"
    rounded: "{rounded.frame}"
    height: "clamp(26rem, calc(100dvh - 10.5rem), 60rem)"
  print-overlay:
    backgroundColor: "{colors.screen-print}"
    textColor: "{colors.terminal-fg}"
  keypad-key:
    backgroundColor: "{colors.print-cream}"
    textColor: "{colors.ink-red-brown}"
    typography: "{typography.key-numeral}"
    rounded: "{rounded.key}"
    height: "3rem"
  keypad-enter:
    backgroundColor: "{colors.ink-red-brown}"
    textColor: "{colors.print-cream}"
    rounded: "{rounded.key}"
    height: "3rem"
  vignette:
    backgroundColor: "{colors.screen-print}"
    textColor: "{colors.terminal-fg}"
    typography: "{typography.vignette}"
    rounded: "{rounded.vignette}"
    height: "9.5rem"
    padding: "0.6rem"
  kbd:
    backgroundColor: "{colors.print-cream}"
    textColor: "{colors.ink-red-brown}"
    rounded: "{rounded.label}"
    padding: "0.12em 0.45em"
  barcode:
    backgroundColor: "{colors.print-cream}"
    textColor: "{colors.ink-red-brown}"
    rounded: "{rounded.label}"
    padding: "0.6rem 0.75rem 0.5rem"
---

# Design System: Polaris (browser edition)

## Overview

**Creative North Star: "The Back of the Box"**

The page at `palawi.fr/polaris/` is printed on the back insert of Polaris's retail box. One cardinal red carries the whole card, with only two inks printed on it: cream and a deep red-brown. The big screenshot is not a picture. It is the live terminal game in a cream-bordered frame. Before play, a printed capture of the title screen sits exactly on the terminal's cells. Pressing the cream sticker cuts that capture away, and the running game is underneath, already showing the same title and prompt.

Density is a retail box's: one dominant screenshot, a narrow column of three feature blurbs, each with a real ASCII monster from the game, then a spec strip, a controls list and the small print with a barcode. The controls behave like hardware. Cream keys rest proud of the card, sink 2px when pressed and lie flat when the game is not listening. The warm near-black screen is the most saturated, highest-contrast area of the page. Everything printed on the red stays quieter than it.

The system rejects the neon or glow "gamer" look: no black page around a centred terminal, and no luminous halos.

**Key Characteristics:**
- One printed red ground under a faint red-brown halftone dot (4px pitch), lighter toward the top.
- Two inks on the red: cream and red-brown. The screen is the only dark, high-contrast area.
- The title is the game's own ░▒▓█ block lettering, generated from the game source and never typeset in a font.
- Controls have real travel: they rest, sink when pressed, and flatten when the game is not listening.
- Printed rules are hairlines: 1px, and 0.5px on 2x screens.
- One signature move: the printed capture is clipped away left to right to show the live game.

## Colors

One printed cardinal red, two inks and a warm black screen. Nothing else appears on the page.

### Primary
- **Box Red** (box-red): the ground of the whole page (`html` background and the middle stop of the body gradient), the sticker's text colour, the `theme-color`, and the `--app-bg` the hub sheet lands on. The body gradient goes from a slightly lighter red at the top (#bd2630) to a slightly darker one at the bottom (#a91c26), with a 4px red-brown halftone (7% alpha) over it.
- **Box Red Deep** (box-red-deep) and **Box Red Dark** (box-red-dark): used only for the scrollbar track and thumb, and dark red is also the terminal's selection background. They never print text.

### Neutral
- **Print Cream** (print-cream): the main ink on the red. Used for headings, the block logo, the screenshot and vignette frames, every key, the tab, the sticker body, the barcode label, focus outlines and `::selection`. Printed rules use it at 45 to 60% alpha.
- **Soft Cream** (soft-cream): the secondary ink on the red. Used for body copy in blurbs, controls and small print, for the flap genre line and for the idle caption (5.0:1 on Box Red).
- **Red-Brown Ink** (ink-red-brown): the second printed ink. Used for text on cream keys, the halftone dot, the barcode bars and the Entrée key's fill. Shadows are this ink's hue (rgb 60 2 10).
- **Screen** (screen): the terminal background, set in both the CSS and the xterm theme. It is the darkest surface on the page.
- **Screen Print** (screen-print): the printed capture and the vignette screens. It is a slightly lifted screen colour under a 3px cream dot at 5% alpha, so the print reads as a printed copy of the screen.
- **Terminal Foreground** (terminal-fg): the game's text and the text printed on the screen (title capture, end and error notes, save line).

### Named Rules
**The Two Inks Rule.** Anything printed on the red uses only cream (at full strength or as soft cream, or as alpha rules) and red-brown. A third ink is not allowed. This is why the small logo keeps only the solid █ cells and drops the shaded tints.

**The Live Screen Rule.** The terminal is the most saturated, highest-contrast area of the page (15.8:1). Printed areas on the red stay at about 5 to 6:1.

## Typography

**Display Font:** Big Shoulders Display 800/900 (with Arial Narrow, sans-serif)
**Body Font:** Figtree 500/700/800 (with ui-sans-serif, system-ui)
**Label/Mono Font:** JetBrains Mono 400, with the generated "Polaris Blocks" font for U+2580–259F

**Character:** Condensed, heavy uppercase box lettering over a friendly geometric sans, plus the terminal's own mono. The real title is not set in any of these fonts. It is the game's block art.

### Hierarchy
- **Display** (900, clamp(2rem, 4.4vw, 2.75rem), 1, uppercase): the sticker label (Jouer / Reprendre / Relancer / Réessayer). The end and error note headings use the same face at clamp(2rem, 6cqi, 3rem)/0.95.
- **Headline** (900, clamp(2.2rem, 4vw, 3rem), 0.95, uppercase): section headings on the card ("Commandes").
- **Title** (900, 1.55rem, 1, uppercase, balanced): blurb headings.
- **Genre** (800, clamp(0.95rem, 1.4vw, 1.2rem), 0.08em, uppercase, soft cream): the single genre line on the flap, next to the logo. It is a box-front genre line, not a label above a heading.
- **Key numeral** (900, 1.45rem, tabular): keypad digits, and `kbd` caps at 0.95em.
- **Body** (Figtree 500, 1rem/1.5): copy, with a 34ch measure for side notes and 68ch for small print.
- **Caption** (Figtree 700, 0.95rem/1.3): the state line under the screenshot.
- **Label** (Figtree 800, 0.8rem, 0.06em, uppercase): spec strip items. The barcode label uses 0.72rem at 0.08em.
- **Small print** (Figtree 500, 0.82rem/1.5): credits and legal.
- **Terminal** (JetBrains Mono 400, line-height 1, letter-spacing 0): xterm's size is fitted per width to between 10 and 16px (15px initial). Vignettes use 0.8rem/1.08 (0.5rem on phones), and the forge flame uses 1.35rem (0.95rem on phones).

### Named Rules
**The Game Letters Rule.** "POLARIS" is always the game's own block art from `internal/game/game.go`, rendered by `scripts/logo.mjs` as crisp-edged SVG cells (3×6 units). The flap logo keeps only the solid █ cells. The printed capture draws ▓▒░ as the same 4×8 ordered dither (75/50/25%) that `scripts/blocks_font.py` bakes into "Polaris Blocks", so the print and the live terminal show identical dots.

**The Mono Stays On Screen Rule.** JetBrains Mono appears only on screen surfaces: the terminal, the capture's prompt line, the ASCII vignettes and the error detail. Copy on the red is never set in mono.

**The Fixed Numerals Rule.** Keypad digits, `kbd` caps and the save line use tabular numerals so they keep fixed positions.

## Layout

The layout is a single 92rem-max column with a fluid gutter (`spacing.gutter`). From top to bottom: flap, back (screenshot plus blurbs), spec strip, controls, small print. Sections are separated by clamp margins (2.5–5rem). Only the flap, the spec strip and the small print draw a printed rule.

- **Desktop (above 64rem):** the back is a two-column grid, `minmax(0,1fr) 19rem`. The screenshot takes about 70% of the width and the blurbs stack in the narrow column. The frame height is clamp(26rem, 100dvh − 10.5rem, 60rem), so the first viewport holds the flap, the screen and its caption. The flap is a grid of tab, logo, genre and keys. The controls section has three columns: heading, list and note. The spec strip has five cells.
- **Tablet (64rem and below):** the back becomes one column. The blurbs become three columns under the caption, the spec strip has three cells per row with a rule between rows, and the controls stack.
- **Phone (40rem and below):** the flap drops the genre line, the tab shrinks to its chevron as a 2.5rem square, and the logo is centred. Each blurb is a 7.5rem vignette beside its text, with 7rem-high screens. The spec strip has two cells per row. Controls terms are 7rem wide. The small print stacks above the barcode.
- **Touch (pointer: coarse):** the keypad appears under the caption (6 columns: 1–5 and back, 6–0 and keyboard, then Entrée across the full width). The frame height becomes clamp(17rem, 100dvh − 19.5rem, 60rem) so the screen and keypad fit together.
- **Phone landscape (coarse, max-height 32rem):** the flap and caption are hidden. The frame fills 100dvh − 4.4rem with a 5px border. The keypad becomes one row (12 equal keys plus a 2fr Entrée) of 2.6rem-high keys.
- **Cramped (terminal under 86 columns):** `body.is-cramped` is set. In portrait on touch, the print scales the whole title to the frame width, puts the prompt line three cells under it, and shows the "Tourne ton téléphone" hint under the sticker.
- Safe-area insets pad the body, and the small print adds the bottom inset.

**The Ninety Columns Rule.** The terminal font size is computed so that 92 columns fit (`floor(width / 92 / 0.6)`, clamped to 10–16px), then the fit addon fills the frame. The 86-column title is never scaled by CSS inside the live terminal.

## Elevation & Depth

Depth is physical and lit from above. Objects on the card cast short shadows that fall downward, tinted with the box's red-brown. Controls have a lit top edge and a darker lower lip. Nothing glows.

### Shadow Vocabulary
- **Rest** (`0 1px 2px rgb(60 2 10 / 0.35), 0 16px 30px -18px rgb(60 2 10 / 0.75)`): vignette screens lying on the card.
- **Lift** (`0 2px 5px rgb(60 2 10 / 0.38), 0 26px 44px -20px rgb(60 2 10 / 0.85)`): the main screenshot frame.
- **Key** (`0 1px 0 rgb(255 255 255 / 0.65) inset, 0 -2px 0 rgb(43 10 13 / 0.12) inset, 0 3px 6px -2px rgb(60 2 10 / 0.55)`): tab, flap keys, keypad keys and barcode at rest.
- **Key down** (`0 1px 0 rgb(255 255 255 / 0.4) inset, 0 1px 2px -1px rgb(60 2 10 / 0.5)`): pressed and toggled-on controls.
- **Sticker** (`0 1px 0 rgb(255 255 255 / 0.7) inset, 0 3px 6px -1px rgb(0 0 0 / 0.45), 0 22px 36px -18px rgb(0 0 0 / 0.7)`): the sticker sits on the near-black screen, so its shadow is neutral black instead of red-brown.
- **Kbd lip** (`0 2px 0 rgb(43 10 13 / 0.35)`): the 2px lower edge of inline key caps.

### Named Rules
**The Falling Light Rule.** Shadows always fall downward and are tinted with the ink of the surface below. A zero-offset, wide-blur shadow is a glow, and the system has none. (The rule is shared with the palawi.fr hub.)

**The Real Travel Rule.** A pressed control moves `translateY(2px)` and switches to the key-down shadow. On fine pointers only, hover lifts it 1px (2px for the sticker). The toggled states (sound off, fullscreen on) hold the pressed position. When the game is not waiting for input, the keypad keys sit 2px down, lose their shadow and drop to 55% opacity. They flatten into the card instead of greying out.

## Shapes

Corners are soft, cut-card corners: 14px for the screenshot frame and sticker, 10px for keys and the tab, 9px for vignettes, 6px for `kbd` caps and the barcode label, and 2px for the caption's square swatch. Screens are framed in solid cream (8px on the main frame, 5px in phone landscape, 4px on vignettes) and clipped (`overflow: hidden`). The sticker alone is tilted −3deg.

**The One Device Pixel Rule.** Printed rules (flap, spec strip borders, controls rows, small-print top) use the `--rule` width: 1px, or 0.5px at 2dppx and above, always in cream at 45–60% alpha. Frames are not rules. They are thick cream borders.

## Components

### Tab (back to palawi.fr)
A cream printed tab: a chevron plus "palawi.fr" in Figtree 800 0.875rem, 2.5rem high, with the key shadow and real travel. On phones it becomes a 2.5rem square showing only the chevron.

### Keys (sound, fullscreen)
2.5rem cream squares, 10px corners, with a 1.2rem stroke icon in red-brown. Sound swaps between the speaker-waves and speaker-cross icons and holds the pressed position when muted. Fullscreen holds the pressed position while active and is hidden when the Fullscreen API is unavailable.

### Printed Screenshot Frame
The live terminal (xterm.js, DOM renderer, block cursor, no bell) on the screen colour, inside an 8px cream border with 14px corners and the Lift shadow. It has fluid inner padding on the top and left only. When the game sends the DEC reverse-video flash (a monster hit), the screen inverts with a hue-rotate for about 90ms and jolts 3px. Under reduced motion the jolt is dropped and the invert is kept. In fullscreen the border and radius go to 0.

### Print Overlay and Cell-Grid Alignment
The printed capture covers the frame on the screen-print colour with its 3px dot. It shows the dithered logo, "Press Enter to START !" in mono, the sticker and an optional save line. Once the terminal exists, `resize()` measures the real cell width and height and the grid origin (`--cw`, `--ch`, `--px`, `--py`, `--fs`). The print is then placed on those cells: the title spans rows 2 to 8 and 85 columns, and the prompt sits on row 12, exactly where the game prints them. The end ("À bientôt, aventurier") and error notes are printed the same way. They are centred, with a cream display heading and copy at a 34ch measure.

**The Cell Grid Rule.** Anything printed over the live screen is positioned in terminal cells, never in arbitrary pixels, so when the cut passes the live screen continues the print without a jump.

### Sticker (the one primary action)
A cream sticker with red display text and a filled play triangle. It is at least 4rem high, tilted −3deg, and casts a neutral shadow. Its label reads "Jouer", or "Reprendre" when a save exists. In that case a save line in tabular numerals sits under it ("Partie de … : niveau N, chapitre N."). When pressed it sinks 2px and scales to 0.98. When the game starts, the sticker area peels off (fades out, rises 8px and turns to −6deg over 180/220ms) while the rest of the print waits for the cut.

**The One Sticker Rule.** Each print carries exactly one sticker, and it is the only primary action on the page.

### Signature: The Cut-Away
The printed capture stays in place until the game first waits for input. At that point the live screen shows the same title and prompt, and the print is clipped away with `inset(0 0 0 0)` → `inset(0 0 0 100%)` over 560ms on the ease-in-out curve. The end and error notes are laid back down with the reverse clip (`inset(0 100% 0 0)` → `inset(0)`, 460ms). Under reduced motion both become a 160ms opacity fade.
**Narrow-screen title skip:** if the terminal is cramped (under 86 columns) when play starts, the 86-column title would wrap. The print stands in for it, the sticker press is answered as the "Press Enter to START" line, and the cut happens on the next screen.

### Caption
A Figtree 700 state line under the frame, led by a 0.55rem square swatch in `currentColor`. It has four exclusive states: idle ("Appuie sur Jouer…"), starting ("Le jeu démarre…"), busy ("Le jeu raconte…"), and waiting ("À toi de jouer…"). The waiting state is the only one in full cream, swatch included. The others use soft cream. The caption is `aria-live="polite"` and is hidden in phone landscape.

### Keypad (touch only)
Cream keys 3rem high in a 6-column grid with a 0.5rem gap, Big Shoulders 900 tabular digits, plus stroke icons for back and keyboard (which opens the soft keyboard for typing the name). Entrée spans the full row in inverted ink: a red-brown fill with uppercase cream text tracked at 0.06em. Keys respond on `pointerdown` and keep focus where it is. **Listening:** keys rest with the key shadow and travel 2px on press. **Flat:** when the game is not waiting, every key flattens (see The Real Travel Rule) and its input is ignored.

### Blurbs with ASCII Vignettes
Each blurb has a small printed screen and two text lines below it. The screen is a 4px cream frame, 9px corners, the Rest shadow and the screen-print dot, with a real monster or object from the game (skeleton, wolf, the forge flame) in mono and centred. Below it are a title-style heading and soft-cream copy at 0.95rem/1.45. The flame uses the larger vignette size.

### Spec Strip
A row of five cells (three on tablet, two on phones) between cream hairlines, with hairline dividers between cells. Each cell has a 1.6rem stroke icon (stroke width 1.8) and an uppercase Figtree 800 label. There is no fill and no card.

### Controls List
A "Commandes" headline, a definition list whose rows are separated by hairlines (`kbd` caps in the term column: 9rem wide, 7rem on phones; soft-cream descriptions), and a 34ch side note about autosave.

### Small Print and Barcode Link
Soft-cream credits at 68ch, with "Polaris" in full cream, and underlined text links. On the right sits the barcode: a cream label with a fixed red-brown bar pattern (10.5rem × 3.2rem, drawn as hard gradient stops) over the label "Code source sur GitHub". It links to the repository and has the key shadow, hover lift and press travel.

### Focus
Every interactive element gets a 3px cream outline, offset 3px.

## Do's and Don'ts

### Do:
- **Do** print everything on the red in the two inks (cream and red-brown), and keep the terminal as the only high-contrast area.
- **Do** render "POLARIS" only from the game's block art via `scripts/logo.mjs`: solid cells for small sizes, the 4×8 dither for the printed capture.
- **Do** place anything laid over the live terminal on measured cells (`--cw`, `--ch`, `--px`, `--py`), and reveal the game with the clip-path cut-away (560ms, cubic-bezier(0.77, 0, 0.175, 1)).
- **Do** give every control real travel: the key shadow at rest, 2px down with the key-down shadow when pressed, and flat at 55% when the game is not listening.
- **Do** draw printed rules with `--rule` in cream at 45–60% alpha.
- **Do** keep tabular numerals on keys, `kbd` caps and the save line.
- **Do** frame every screen surface in solid cream with soft corners and a downward shadow tinted in red-brown.

### Don't:
- **Don't** put the game in a dark page with a centred terminal window and a generic play button. The page is the red box.
- **Don't** use neon, glow, or zero-offset luminous halos. Shadows fall downward.
- **Don't** typeset the title in a font, or print the ▓▒░ tints as a third ink on the red.
- **Don't** use JetBrains Mono for copy on the red. Mono belongs to screen surfaces.
- **Don't** show a picture of the game where the live game can run. Vignettes are the game's real ASCII art, not mock-ups.
- **Don't** add a second sticker or competing primary action to a print.
- **Don't** add uppercase labels above headings. The flap genre line is the only small uppercase line, and it sits beside the logo.
- **Don't** print ratings, review quotes or player counts on the box.
