---
name: Polaris (browser edition)
description: An indie-RPG battle screen whose white-framed box is the live terminal game.
colors:
  stage: "#0c0c12"
  frame-white: "#f4f4f6"
  soft-ink: "#b4b4c0"
  unlit-ink: "#6e6e7d"
  rule-line: "#2a2a35"
  command-yellow: "#f2cf4b"
typography:
  section-heading:
    fontFamily: "Pixelify Sans, Figtree, system-ui, sans-serif"
    fontSize: "2rem"
    fontWeight: 700
    lineHeight: 1.1
  command:
    fontFamily: "Pixelify Sans, Figtree, system-ui, sans-serif"
    fontSize: "1.5rem"
    fontWeight: 700
    lineHeight: 1
    letterSpacing: "0.08em"
  dialogue:
    fontFamily: "Pixelify Sans, Figtree, system-ui, sans-serif"
    fontSize: "clamp(1.05rem, 2.6vw, 1.3rem)"
    fontWeight: 500
    lineHeight: 1.4
  lead-in:
    fontFamily: "Pixelify Sans, Figtree, system-ui, sans-serif"
    fontWeight: 600
  label:
    fontFamily: "Pixelify Sans, Figtree, system-ui, sans-serif"
    fontSize: "0.95rem"
    fontWeight: 600
    lineHeight: 1
  status:
    fontFamily: "Pixelify Sans, Figtree, system-ui, sans-serif"
    fontSize: "0.95rem"
    fontWeight: 500
    lineHeight: 1.3
  body:
    fontFamily: "Figtree, ui-sans-serif, system-ui, sans-serif"
    fontSize: "1rem"
    fontWeight: 500
    lineHeight: 1.55
  terminal:
    fontFamily: "JetBrains Mono, Polaris Blocks, ui-monospace, Menlo, Consolas, monospace"
    fontSize: "15px"
    fontWeight: 400
    lineHeight: 1
    letterSpacing: "0"
rounded:
  none: "0px"
spacing:
  frame-box: "4px"
  frame-active: "3px"
  frame-line: "2px"
  top-edge: "3.5rem"
  gutter: "clamp(12px, 3vw, 32px)"
  keypad-gap: "0.45rem"
  box-width-talk: "46rem"
  box-width-play: "64rem"
components:
  battle-box:
    backgroundColor: "{colors.stage}"
    rounded: "{rounded.none}"
    width: "min(100%, 46rem)"
    height: "15.5rem"
  battle-box-playing:
    width: "min(100%, 64rem)"
    height: "calc(100dvh - 3.5rem - 5rem)"
  command-go:
    textColor: "{colors.command-yellow}"
    typography: "{typography.command}"
    rounded: "{rounded.none}"
    padding: "0 1.5rem 0 1.2rem"
    height: "3.4rem"
  command-go-hover:
    backgroundColor: "{colors.command-yellow}"
    textColor: "{colors.stage}"
  command-quiet:
    textColor: "{colors.soft-ink}"
    typography: "{typography.label}"
    rounded: "{rounded.none}"
    padding: "0 0.85rem 0 0.7rem"
    height: "2.4rem"
  command-quiet-hover:
    textColor: "{colors.frame-white}"
  command-icon:
    textColor: "{colors.soft-ink}"
    rounded: "{rounded.none}"
    size: "2.4rem"
  command-icon-off:
    textColor: "{colors.unlit-ink}"
  keypad-key:
    textColor: "{colors.frame-white}"
    rounded: "{rounded.none}"
    height: "3rem"
  keypad-key-unlit:
    textColor: "{colors.unlit-ink}"
  keypad-key-pressed:
    backgroundColor: "{colors.command-yellow}"
    textColor: "{colors.stage}"
  keypad-enter:
    textColor: "{colors.command-yellow}"
    rounded: "{rounded.none}"
    height: "3rem"
  kbd-key:
    textColor: "{colors.frame-white}"
    rounded: "{rounded.none}"
    padding: "0 0.4em"
  status-legend:
    backgroundColor: "{colors.stage}"
    textColor: "{colors.soft-ink}"
    typography: "{typography.status}"
    padding: "0 0.5rem"
  encounter-box:
    backgroundColor: "{colors.stage}"
    textColor: "{colors.soft-ink}"
    rounded: "{rounded.none}"
    padding: "clamp(1.25rem, 3vw, 2rem)"
---

# Design System: Polaris (browser edition)

## Overview

**Creative North Star: "The Battle Box"**

The page is an indie-RPG battle screen. A dark, calm stage holds one white square pixel frame in the middle, and that frame is the live game: the unmodified Go terminal program, compiled to WebAssembly, running in an xterm.js terminal. Everything else on the page is either the game's own lettering (the dithered ░▒▓█ title, printed from the game source), the box's voice (lines that open with the narrator's `*`), or a quiet pixel control around the edge. One soft yellow marks what is active; nothing else is coloured.

The system is flat and literal. There is no radius, no shadow, no gradient and no glow; depth only ever comes from one thing sitting on top of another (the dialogue panel over the hidden terminal, the frame over the stage). State reads as illumination: the frame is lit white when the game is waiting for the player and dims while the game narrates, the keypad's keys stay drawn but unlit when input would be ignored, and a muted sound button dims instead of vanishing. The signature motion is the battle box opening from its compact dialogue size to full play size when the game starts, as battle boxes resize in the genre.

Confirmed rejections (owner, 2026-10-29): the red "game box" world, neon or glowing "gamer" darkness, a red ground, a redrawn logo, an off-centre terminal, and a macOS-style terminal window with a hero button.

**Key Characteristics:**
- One centred white pixel frame (4px) that is the game itself, whole at every size.
- One yellow (#f2cf4b) for the active thing: the selected command, the save star, the cursor, the Entrée key, a pressed key, focus and selection.
- Unlit states are designed: dim frame, dim keys, dim icon buttons.
- The game's own dithered title, never a redrawn wordmark.
- Square corners, flat fills, depth by overlap only.

## Colors

A near-black blue stage, a near-white ink for the frame and text, two greys for secondary and unlit ink, one hairline grey for rules, and one soft yellow for what is active.

### Primary
- **Command Yellow** (`command-yellow`): the selected command's 3px frame and label (JOUER / Reprendre / Rejouer / Réessayer), its filled hover, the save star, the Entrée key on the keypad, the background of a pressed keypad key, the terminal cursor and selection, the global focus ring (3px outline, 3px offset) and `::selection`. It always means "this is live / this is the choice".

### Neutral
- **Stage** (`stage`): the page ground, the inside of the battle box, the terminal background, the text colour on yellow fills. Also the `theme-color` and the favicon's ground.
- **Frame White** (`frame-white`): the battle box and encounter box frames (the battle box only on the player's turn and before play), headings, dialogue lines, terminal foreground, the encounter's ASCII foes, asterisks and lead-ins, lit keypad keys, `kbd` text, the save text and "À toi de jouer".
- **Soft Ink** (`soft-ink`): secondary reading text (dialogue hint, error detail, encounter lines, control descriptions, credits, "Le jeu raconte…", the rotate hint) and the resting label of quiet and icon commands.
- **Unlit Ink** (`unlit-ink`): borders of resting commands and `kbd`, the battle box frame while the game narrates, unlit keypad keys, the muted sound button and pressed-in fullscreen button, terminal "bright black".
- **Rule Line** (`rule-line`): 2px rules between below-the-fold parts and control rows, unlit key borders, scrollbar thumb.

### Named Rules
**The One Yellow Rule.** Yellow belongs only to what is active or selected: one command per dialogue, the save star, the cursor, the Entrée key, the pressed key, focus. Never a heading, never a decoration, never a second accent hue beside it.

**The Unlit, Not Gone Rule.** Inactive things dim to `unlit-ink` / `rule-line`; they keep their shape and place. Only the yellow and the white are "lit".

**The Dim Is Not Reading Ink Rule.** `unlit-ink` (about 3.9:1 on the stage) is for borders and inactive states, not for text someone must read; reading text uses `soft-ink` or `frame-white`.

## Typography

**Pixel Font:** Pixelify Sans 500 / 600 / 700 (fallback Figtree, system-ui) for every label, command, dialogue line, status legend, encounter lead-in and heading.
**Body Font:** Figtree 500 / 700 for reading copy below the fold and the dialogue hint.
**Terminal Font:** JetBrains Mono 400, with the generated "Polaris Blocks" font (U+2580-259F only) for the block elements.

**Character:** the pixel face speaks for the page in the game's register; Figtree carries the reading; the mono face belongs to the game and appears only where the game's own text or art does (the terminal, the encounter's ASCII sprites, the narrator's asterisks, the error detail).

### Hierarchy
- **Display:** the title is not set in a font. It is the game's own 86-column ░▒▓█ splash art, rendered to SVG by `web/scripts/logo.mjs` (3 x 6 units per cell; █ solid, ▓ ▒ ░ as the 75 / 50 / 25 % ordered dot dither of Polaris Blocks), inlined at `min(100%, 46rem)` in white.
- **Section heading** (Pixelify 700, 2rem, 1.1): "Comment jouer", "Dans le jeu".
- **Command** (Pixelify 700, 1.5rem, 0.08em, uppercase): the yellow selected command; the keypad's Entrée uses the same treatment at 1.25rem.
- **Dialogue** (Pixelify 500, clamp(1.05rem, 2.6vw, 1.3rem), 1.4, max 60ch, `text-wrap: pretty`): the box's `*` lines.
- **Lead-in** (Pixelify 600, inherits size, white): the bold opening phrase of each encounter line ("Quatre chapitres.").
- **Label** (Pixelify 600, 0.95rem): quiet commands, control-list terms, `kbd` (600, 0.9em, tabular numerals).
- **Status** (Pixelify 500, 0.95rem, 1.3): the legend on the box's bottom edge.
- **Body** (Figtree 500, 1rem, 1.55): descriptions and encounter lines (max 62ch, `text-wrap: pretty`); credits at 0.9rem, max 70ch.
- **Terminal** (JetBrains Mono 400, line-height 1, letter-spacing 0): 15px at creation, then fitted per box width to between 10px and 16px so about 92 columns fit; keypad digits are Pixelify 700 1.35rem.

### Named Rules
**The Game's Own Letters Rule.** The title is generated from `internal/game/game.go`, never redrawn or re-lettered. Block characters always render through Polaris Blocks (exact cells, 4 x 8 dot dither, bleed to avoid seams), on the page and in the terminal alike.

**The Three Voices Rule.** Pixel for the page's voice, Figtree for reading, mono only for the game's own text. No fourth face.

## Layout

**Stage.** The first viewport is a single centred column (`min-height: 100dvh`, flex, centred both ways): the title, then the battle box with its status legend seated on its bottom edge. Width is `min(100%, 46rem)` for title, box and status before play; box and status widen to `min(100%, 64rem)` while playing. Side padding is the gutter token; the top padding clears the absolutely positioned top edge (3.5rem + 0.5rem).

**Top edge.** A 3.5rem band, back link to palawi.fr on the left, sound and fullscreen icon commands on the right (0.5rem gap), respecting safe-area insets.

**Below the fold.** Max 64rem, centred; parts separated by a 2px `rule-line` top border with clamp(3rem, 7vw, 4.5rem) margin and clamp(2rem, 5vw, 3rem) padding. Control list rows: 9rem term column, 2px bottom rules. "Dans le jeu" is a single encounter box at the part's full width.

**Adaptation.**
- **Desktop (fine pointer):** as above; box 15.5rem tall while talking, `100dvh - 3.5rem - 5rem` while playing; hovers only under `(hover: hover) and (pointer: fine)`.
- **Narrow (max-width 40rem):** talking box grows to 17.5rem; the encounter's ASCII foes shrink to 0.45rem (the flame to 0.8rem) and stay in one row; control terms 7rem.
- **Touch (coarse pointer):** before play the title, box and keypad stay centred as a group, with 2rem between the top bar and the stage content (top padding 3.5rem + 2rem); while playing the stage starts at the top (3.5rem + 0.5rem) and the box is `100dvh - 3.5rem - 15.5rem` with a 15rem minimum. The keypad shows 1rem under the box (6 columns: 1-5 + erase, 6-0 + keyboard, full-width Entrée).
- **Phone portrait (coarse):** only while playing and only when the terminal has fewer than 86 columns, a left-aligned soft-ink Pixelify 0.9rem hint under the keypad suggests turning the phone.
- **Phone landscape (coarse, max-height 32rem):** the top edge, title and status legend hide; the box goes full width at `100dvh - 4.6rem`; the keypad becomes one row (12 equal keys + a double-width Entrée, 0.35rem gap, 2.6rem tall).
- **Fullscreen:** the fullscreen button sends the battle box itself fullscreen (100% x 100%, frame width 0); the button is hidden where the Fullscreen API is unavailable.
- **Reduced motion:** title, box and status keep only 160ms border-colour and opacity transitions (no resizing animation, no dialogue fade delay, no 500ms pre-start wait); the hit jolt and the screen invert are both off.

## Elevation & Depth

Completely flat. No `box-shadow`, no gradient, no glow, no blur anywhere. Depth is conveyed only by overlap: the status legend sits on the battle box's bottom edge, each part on a stage-coloured patch that cuts the frame line; the dialogue panel covers the (hidden) terminal inside the same frame; the frame sits on the stage. Hierarchy comes from frame weight (4px battle and encounter boxes, 3px active command / Entrée, 2px everything else) and from lit versus unlit ink.

### Named Rules
**The Overlap-Only Rule.** If something needs to read as "above", it overlaps and covers, as the status legend covers the frame's bottom edge; it is never shaded, lifted or haloed.

## Shapes

Every corner is square (0px). Every container is a solid pixel border in one of three weights: 4px for the battle box and the encounter box, 3px for the active command and the keypad Entrée, 2px for commands, keys, `kbd` and rules. Icons are 12 x 12 SVG strokes at 1.5, square caps, `crispEdges`; the four-point save star is the one filled, geometric shape and is always yellow. The favicon repeats the system: stage square, white 3px frame, yellow star.

### Named Rules
**The Square Pixel Rule.** No radius, ever; frames are the only shape language.

## Components

### Commands (top edge links and pixel buttons)
Quiet, framed pixel buttons that never compete with the box.
- **Shape:** square, 2px `unlit-ink` border, transparent fill.
- **Quiet** (2.4rem tall, Pixelify 600 0.95rem, chevron icon + label): back to palawi.fr; "Code source" in the credits, where it is drawn in `frame-white`.
- **Icon** (2.4rem square, 1.1rem icon): sound (speaker / crossed speaker), fullscreen (corner brackets). Unlit when muted or when fullscreen is on: `rule-line` border, `unlit-ink` icon.
- **Hover (fine pointers only):** border and label go to `frame-white`, 140ms. **Active:** scale 0.97 with the ease-out curve. **Focus:** 3px yellow outline, 3px offset.

### The selected command
The one yellow thing in a dialogue, the genre's highlighted choice.
- 3.4rem tall, 3px `command-yellow` border, yellow uppercase Pixelify 700 1.5rem with 0.08em tracking, a filled play triangle before the label.
- **Hover:** fills yellow, label turns `stage`. **Active:** scale 0.97.
- Label follows the save: "Jouer" without a save, "Reprendre" with one; "Rejouer" after quitting, "Réessayer" on error.

### The dithered title and its fold
The game's own splash title in white SVG, above the box. When the game starts it folds away: grid row 1fr to 0fr and its bottom margin to 0 over 480ms (ease-in-out), opacity to 0 over 200ms, so the box can take the stage.

### The battle box (signature)
The white frame that is the game. Stage-coloured fill, overflow hidden, centred at every size.
- **Idle (talking):** 4px `frame-white` frame, `min(100%, 46rem)` x 15.5rem, the dialogue panel visible, the terminal created behind it at opacity 0.
- **Expanding:** on JOUER, the dialogue fades out (160ms), the page enters its playing state and the box animates width and height over 480ms (ease-in-out) to `min(100%, 64rem)` x `100dvh - 3.5rem - 5rem`; the title folds; after about 500ms the terminal is reset, fitted and fades in (200ms), then the game starts.
- **Narrating:** while the game prints and is not waiting for input, the frame dims to `unlit-ink` (160ms); the legend reads "Le jeu raconte…"; typed keys are ignored.
- **Your turn:** when the game blocks on a read, the frame is lit `frame-white` again; the legend reads "À toi de jouer" in white; the keypad lights.
- **Hit:** when the game prints its reverse-video flash, the box jolts (translate 4px, -3px) and the screen inverts for the first 45% of 200ms, a hard two-step cut, no easing. Both are off under reduced motion.
- **After the game / error:** the box returns to dialogue size with the end or error dialogue.

### Dialogue lines
The box's own voice. Each line is a two-column grid: a large mono asterisk (1.25em) then Pixelify 500 text, max 60ch with `text-wrap: pretty`, 0.35rem apart. Below the lines, the action row: the selected command and, on the start dialogue, a Figtree 0.875rem `soft-ink` hint with a `kbd`. The panel's padding is clamp(1.1rem, 3vw, 1.6rem) clamp(1.1rem, 3vw, 1.8rem) clamp(1.1rem, 3vw, 1.5rem). Dialogues fade in and out over 160ms; the box does the moving.

### Status legend
A legend seated on the battle box's bottom edge, like a label on a frame. Same width as the box, pulled up by a negative top margin (`-0.65rem - frame/2`) so it straddles the 4px line, 1rem side inset, z-index above the box; each part is a `stage`-filled patch with 0.5rem side padding, so it cuts the frame. One line, Pixelify 500 0.95rem, `soft-ink`, `aria-live="polite"`.
- **Save (left):** shown only when a save exists, no-wrap: yellow star then a short white tabular summary, "Nin · niv. 1 · chap. 1" (or "jeu fini"), with the title attribute "Partie sauvegardée dans ce navigateur".
- **Turn (right):** hidden before play; "Le jeu raconte…" in `soft-ink` while narrating, "À toi de jouer" in white on the player's turn.

### Touch keypad
Pixel keys that light with the player's turn (coarse pointers only).
- Keys: 3rem tall, 2px border, transparent fill, Pixelify 700 1.35rem tabular digits; erase and keyboard keys carry 1.2rem stroke icons.
- **Lit (game waiting):** `frame-white` border and digit. **Unlit:** `rule-line` border, `unlit-ink` digit; keys stay in place.
- **Pressed:** fills yellow with `stage` digit (80ms fill).
- Sits 1rem under the box.
- **Entrée:** full width (portrait), 3px yellow border, yellow uppercase Pixelify 1.25rem, 0.08em tracking; unlit like the others when not waiting.
- Key presses use pointerdown and never steal focus; the keyboard key opens the soft keyboard for typing a name.

### Keyboard keys (`kbd`)
Inline 2px `unlit-ink` box, white Pixelify 600 0.9em, min width 1.7em, tabular numerals. Used in the hint and the control list.

### Controls list
A definition list of four rows: key terms in Pixelify 600 0.95rem, descriptions in `soft-ink` Figtree, each row closed by a 2px `rule-line`. Followed by a note led by the yellow save star.

### Encounter box ("Dans le jeu")
A second battle box, below the fold, staging what is in the game as an encounter.
- **Frame:** 4px `frame-white`, square, padding clamp(1.25rem, 3vw, 2rem).
- **Foes (top):** the game's three ASCII sprites (two foes and the forge flame) lined up in one row, bottom-aligned and spaced around, in white JetBrains Mono 0.85rem / 1.08 (the flame 1.35rem / 1.1); a 2px `rule-line` under them, overflow hidden. Decorative (`aria-hidden`).
- **Lines (below):** narrator lines, each a grid of a white mono `*` (1.25em) and Figtree `soft-ink` text, max 62ch, `text-wrap: pretty`, 0.6rem apart; each opens with a white Pixelify 600 lead-in ("Quatre chapitres.", "Combats au tour par tour.", "Forge et boutiques.").

### Credits
The last part: Figtree 0.9rem `soft-ink`, names in white bold, max 70ch; a link row with the "Code source" quiet command in white, plain underlined links (1px, 4px offset) and the no-cookie note.

### Share image
`web/src/og.png` (1200 x 630) is a plain Playwright capture of the built first viewport before play (`web/scripts/og.py`); nothing is restyled for it.

### Related surface
The palawi.fr hub widget wears this world under the Guest App Rule; it is documented in the hub's own DESIGN.md (`/home/claude/palawifr-static-pages/DESIGN.md`, "Polaris Battle Box (hub)"), not here.

## Do's and Don'ts

### Do:
- **Do** keep the battle box whole, centred and square-framed at every size; restack everything else around it.
- **Do** show game state as illumination: frame white on the player's turn, `unlit-ink` while the game narrates; keys lit or unlit, never removed.
- **Do** keep exactly one yellow command per dialogue and use yellow only for active things (command, save star, cursor, Entrée, pressed key, focus, selection).
- **Do** generate the title from the game source with `web/scripts/logo.mjs` and render block characters through Polaris Blocks (`web/scripts/blocks_font.py`).
- **Do** open lines in a box's voice with the narrator's `*`, in the battle box and the encounter box alike.
- **Do** seat labels about the box on its frame as a legend (stage-filled patches over the edge), not as a separate row below it.
- **Do** use the 480ms ease-in-out curve (cubic-bezier(0.77, 0, 0.175, 1)) for the box, title and status resizing, and ease-out (cubic-bezier(0.23, 1, 0.32, 1)) for press feedback; collapse to colour and opacity changes under reduced motion.
- **Do** answer the game's "Press Enter to START" automatically when the terminal is narrower than 86 columns, so the wrapped title never shows as noise; JOUER already was the player's Enter.

### Don't:
- **Don't** round a corner, add a shadow, gradient, glow, blur or neon halo.
- **Don't** bring back the red ground or the "game box" world, or a macOS-style terminal window.
- **Don't** redraw or re-letter the Polaris title.
- **Don't** add a second accent colour or use yellow for headings or decoration.
- **Don't** use `unlit-ink` for text someone has to read.
- **Don't** use Undertale or Deltarune characters, logos or exact command names (FIGHT / ACT / ITEM / MERCY); only the generic battle-box grammar.
- **Don't** hide inactive controls; dim them.
