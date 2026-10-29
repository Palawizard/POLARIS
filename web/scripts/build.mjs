// Assembles the static site in web/dist from:
//   - web/src (page, styles, bridge script),
//   - xterm.js and fonts from node_modules,
//   - the game's audio assets (../assets/audio),
//   - polaris.wasm and wasm_exec.js (path given as arguments).
// Usage: node scripts/build.mjs <polaris.wasm> <wasm_exec.js>
import { createHash } from "node:crypto";
import { cpSync, mkdirSync, readFileSync, rmSync, writeFileSync } from "node:fs";
import { dirname, join } from "node:path";
import { fileURLToPath } from "node:url";

const web = join(dirname(fileURLToPath(import.meta.url)), "..");
const repo = join(web, "..");
const out = join(web, "dist");
const [wasm, wasmExec] = process.argv.slice(2);
if (!wasm || !wasmExec) {
  console.error("usage: node scripts/build.mjs <polaris.wasm> <wasm_exec.js>");
  process.exit(1);
}

rmSync(out, { recursive: true, force: true });
mkdirSync(join(out, "vendor"), { recursive: true });
mkdirSync(join(out, "fonts"), { recursive: true });

const nm = join(web, "node_modules");
const copy = (from, to) => cpSync(from, join(out, to), { recursive: true });

copy(join(web, "src"), ".");
copy(join(nm, "@xterm/xterm/lib/xterm.js"), "vendor/xterm.js");
copy(join(nm, "@xterm/xterm/css/xterm.css"), "vendor/xterm.css");
copy(join(nm, "@xterm/addon-fit/lib/addon-fit.js"), "vendor/addon-fit.js");
copy(join(nm, "@fontsource/jetbrains-mono/files/jetbrains-mono-latin-400-normal.woff2"), "fonts/jetbrains-mono-latin-400-normal.woff2");
for (const w of [500, 600, 700]) {
  const f = `pixelify-sans-latin-${w}-normal.woff2`;
  copy(join(nm, "@fontsource/pixelify-sans/files", f), `fonts/${f}`);
}
for (const w of [500, 700]) {
  const f = `figtree-latin-${w}-normal.woff2`;
  copy(join(nm, "@fontsource/figtree/files", f), `fonts/${f}`);
}
copy(join(repo, "assets/audio"), "assets/audio");
copy(wasm, "polaris.wasm");
copy(wasmExec, "wasm_exec.js");

// Cache busting: one version for the whole release, from the wasm and the page sources.
const h = createHash("sha256");
h.update(readFileSync(wasm));
for (const f of ["index.html", "polaris.js", "polaris.css", "logo-print.svg", "fonts/polaris-blocks.woff2"]) h.update(readFileSync(join(web, "src", f)));
const version = h.digest("hex").slice(0, 12);
const index = join(out, "index.html");
// The logo is the game's own block lettering (scripts/logo.mjs), inlined so it takes the text colour.
const logoPrint = readFileSync(join(web, "src/logo-print.svg"), "utf8").trim();
writeFileSync(
  index,
  readFileSync(index, "utf8")
    .replaceAll("__VERSION__", version)
    .replaceAll("__LOGO_PRINT__", logoPrint),
);
rmSync(join(out, "logo-print.svg"));
console.log(`web/dist ready (version ${version})`);
