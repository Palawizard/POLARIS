// Generates src/logo-print.svg from the game's own title art (last frame of the
// splash animation in internal/game/game.go). Each block character becomes a
// cell: █ solid, ▓ ▒ ░ as the same 4 x 8 dot dither as the terminal, so the title
// matches what the game prints. Run: node scripts/logo.mjs
import { readFileSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const web = join(dirname(fileURLToPath(import.meta.url)), "..");
const src = readFileSync(join(web, "../internal/game/game.go"), "utf8");
const frames = [...src.matchAll(/fmt\.Println\("(░[^"]*)"\)/g)].map((m) => m[1]);
const art = frames.slice(-7).map((l) => [...l.replace(/\s+$/, "")]);
const W = 3; // cell width
const H = 6; // cell height (terminal cells are about twice as tall as wide)
const cols = Math.max(...art.map((r) => r.length));
const tint = { "█": 1, "▓": 0.72, "▒": 0.46, "░": 0.24 };
const runs = { 1: [], 0.72: [], 0.46: [], 0.24: [] };
art.forEach((row, y) => {
  let x = 0;
  while (x < row.length) {
    const t = tint[row[x]];
    if (!t) { x++; continue; }
    let end = x;
    while (end + 1 < row.length && tint[row[end + 1]] === t) end++;
    runs[t].push(`M${x * W} ${y * H}h${(end - x + 1) * W}v${H}h-${(end - x + 1) * W}z`);
    x = end + 1;
  }
});
const box = `viewBox="0 0 ${cols * W} ${art.length * H}" fill="currentColor" role="img" aria-label="Polaris" preserveAspectRatio="none" shape-rendering="crispEdges"`;

// logo-print.svg: shades as the 4 x 8 dot dither of the "Polaris Blocks" font
// (scripts/blocks_font.py), so the printed capture matches the live terminal.
const d = W / 4; // dot size; the dither repeats every 2 dots
const pat = (id, dots) =>
  `<pattern id="${id}" width="${2 * d}" height="${2 * d}" patternUnits="userSpaceOnUse">${dots
    .map(([c, r]) => `<rect x="${c * d}" y="${r * d}" width="${d}" height="${d}" fill="currentColor"/>`)
    .join("")}</pattern>`;
const defs = `<defs>${pat("pb25", [[0, 0]])}${pat("pb50", [[0, 0], [1, 1]])}${pat("pb75", [[0, 0], [1, 0], [0, 1]])}</defs>`;
const fills = { 1: "currentColor", 0.72: "url(#pb75)", 0.46: "url(#pb50)", 0.24: "url(#pb25)" };
const dithered = Object.entries(runs)
  .filter(([, p]) => p.length)
  .map(([t, p]) => `<path fill="${fills[t]}" d="${p.join("")}"/>`)
  .join("");
writeFileSync(join(web, "src/logo-print.svg"), `<svg xmlns="http://www.w3.org/2000/svg" ${box}>${defs}${dithered}</svg>\n`);
console.log(`src/logo-print.svg: ${cols}x${art.length} cells`);
