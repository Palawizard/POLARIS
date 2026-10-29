# Product

<!-- impeccable:product-schema 1 -->

## Platform

web

## Users

- Visitors of palawi.fr who tap the Polaris widget on the hub: curious, often on a phone, here for a quick session.
- Friends and classmates who play it through, come back later and resume their save.

## Product Purpose

Polaris is a small turn-based RPG written in Go for the terminal, as a school project at Ynov Toulouse by Baptiste D.M. (Palawi), Jules D. and Nin L.P. The browser edition at `palawi.fr/polaris/` runs the same Go code compiled to WebAssembly inside a terminal emulator, so anyone can play without installing anything.

Success: one tap from the hub and the game is running; a returning player picks up their save.

## Positioning

Not a remake or a mock-up: the original terminal game itself, unchanged, running in the page. Text, ASCII monsters, block-letter title, numbered menus and typed answers are the game.

## Operating Context

- Played with the keyboard (numbers + Enter, a name typed once); on phones through an on-screen keypad.
- Four chapters (the last is the final boss), training fights, item and spell shops, selling, a black-smith forge, three classes (Human, Elf, Dwarf).
- The game is in English; the page around it is in French.
- Autosave in the browser (localStorage) each time the player is back on the main menu; "Continue / New game" at launch.

## Capabilities and Constraints

- The game must stay the unmodified terminal program: the page only provides the terminal, input, sound and storage around it.
- The title art is 86 columns wide; the terminal needs about 90 columns to show it unwrapped (phones in portrait are cramped).
- Hosted on Palawi's own infra under `palawi.fr/polaris/`, as a static site (no server-side state).

## Brand Commitments

- Name: Polaris. Title rendered by the game in block characters (░▒▓█).
- Credits shown: Baptiste D.M., Jules D., Nin L.P.; music and sound from Tomodachi Life (Nintendo), Deltarune and Undertale (Toby Fox), kept as in the repository (owner's choice, 2026-10-29).
- Guest app of palawi.fr: it keeps its own identity (DESIGN.md of palawi.fr, "The Guest App Rule").
- Owner dislikes the neon/glow "gamer" look (black ground with luminous halos).

## Evidence on Hand

- The game's own screens: title animation, ASCII monsters (goblin, wolf, troll, final boss...), menus. README screenshots in `assets/readme/images/` (title GIF, fight, shops, blacksmith).
- Music and SFX in `assets/audio/`.
- No player counts, reviews or awards exist; do not invent them.

## Product Principles

1. The game first: the first screen is the game, one action from playing.
2. The real thing: show and run the actual program, never a picture of it.
3. Pick up where you left off: saving is automatic and invisible.
4. Nothing invented: credits and facts as they are.
