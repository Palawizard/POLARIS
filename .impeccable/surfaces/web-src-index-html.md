---
version: 1
slug: "web-src-index-html"
primary_target: "web/src/index.html"
related_targets: []
---

## Scope

`palawi.fr/polaris/`, the browser edition of Polaris (Go terminal RPG compiled to WebAssembly). Mode: Experience, the game leads from the first viewport. Audience: palawi.fr visitors on phones for a quick session, friends who come back to their save. Related: the Polaris widget on the palawi.fr hub (full width under the PalaCards booster), which wears this world under the Guest App Rule.

Redesign of 2026-10-29 (second round): the red "game box" was rejected by Palawi (red hurts the eyes, terminal not centred, redrawn logo ugly). Pinned by Palawi: dark and calm (no neon, no glow), the game's own ░▒▓█ title as it shows in the terminal, the terminal centred alone in the first viewport, everything else below.

## Direction contract

THESIS: The page is an indie-RPG battle screen: a dark stage with one white-framed box in the middle, and the box is the live game. Refuses the generic dark page with a macOS-style terminal window and a hero button.

OWN-WORLD: Near-black blue ground, one white for the frame and text, one soft yellow for what is active (the selected command, the save star) and nothing else. Square pixel frames (4px white), no radius, no shadow, no gradient, no glow: depth only by overlap. Pixelify Sans for every label and command, JetBrains Mono plus the Polaris Blocks font inside the terminal, Figtree for reading copy below the fold. The title is the game's dithered ░▒▓█ lettering in white.

STORY: The visitor sees the game's title over a small battle box that speaks in the game's voice ("* ..."), presses JOUER, and the box opens to full size with the game running in it. The frame is lit white when it is the player's turn and dims while the game narrates. Below: how to play, what is in the game, credits.

FIRST VIEWPORT: Top edge: small pixel link back to palawi.fr left, sound and fullscreen pixel buttons right. Centre: the dithered title, then the battle box centred both ways (compact before play, about 90 columns wide and filling the height once playing), under it a yellow save star line when a save exists. Phones: the box full width, pixel keypad under it.

FORM: Indie RPG battle box, position 3 on the ordered list, seed:b8124c8e6a3c86b3. Raises: the box stays whole and centred at every size while everything else restacks around it (from the minihompy room); unlit states are designed, keys and frame dim instead of disappearing (from the seven-segment display); the game's state reads as illumination, frame lit on the player's turn (from the signboard alley); only the active thing gets the bright edge (from the moon-shadow bazaar); depth only by overlap, never by shading (from the layered papercut). Signature interaction: the battle box expands from its compact dialogue size to the full play size when the game starts, as battle boxes resize in the genre.

FINISH: unreviewed and undocumented is unfinished; this build ends with the finish review, the verdict, DESIGN.md, and every shipping raster carrying its provenance

## Constraints

The game stays the unmodified terminal program; the terminal keeps about 90 columns where the screen allows. Page in French, game in English. Credits and music sources as in the repo. No neon or glow, no red ground. No Undertale characters, logos or exact command names (FIGHT/ACT/ITEM/MERCY): only the generic battle-box grammar. No invented claims.
