// Polaris in the browser.
//
// The game is the unmodified Go program compiled to WebAssembly. Go's js/wasm
// runtime does all its file I/O through `globalThis.fs`, so we plug a small
// "TTY" in there:
//   - fd 1/2 writes go to an xterm.js terminal;
//   - fd 0 reads behave like a terminal in canonical mode: they block until
//     the player presses Enter, then return that line (and only that line).
// Keys typed while the game is not waiting for input are ignored, so Enter
// pressed during an animation is never replayed into the next menu.
//
// Audio (`polarisAudio`) and saves (`polarisSave`) are small bridges called by
// the js build of the game (internal/audiosystem/audio_js.go, internal/save).

/* global Terminal, FitAddon, Go */

const VERSION = document.documentElement.dataset.version || "dev";
const COLS_TARGET = 92; // the title art is 86 columns wide
const enc = new TextEncoder();

const $ = (sel) => document.querySelector(sel);
const shell = $(".term");
const stage = $(".term__screen");
const overlayStart = $("#start");
const overlayEnd = $("#end");
const overlayError = $("#error");
const keypad = $(".keypad");
const soundBtn = $("#sound");

/* ------------------------------------------------------------------ storage */

function store(key, value) {
  try {
    if (value === undefined) return localStorage.getItem(key);
    if (value === null) localStorage.removeItem(key);
    else localStorage.setItem(key, value);
  } catch {
    /* storage disabled or full: the game simply runs without it */
  }
  return null;
}

const SAVE_KEY = "polaris.save.v1";
window.polarisSave = {
  get: () => store(SAVE_KEY) || "",
  set: (json) => {
    store(SAVE_KEY, json);
    updateStartCopy();
  },
  remove: () => {
    store(SAVE_KEY, null);
    updateStartCopy();
  },
};

/* -------------------------------------------------------------------- audio */

const audio = (() => {
  let ctx = null;
  let sfxGain = null;
  let muted = store("polaris.muted") === "1";
  const buffers = new Map(); // url -> AudioBuffer | Promise
  const ids = new Map(); // id -> url
  const lastStart = new Map(); // url -> ms, to thin out the typewriter tick
  const music = new Audio();
  music.loop = true;
  music.preload = "auto";
  music.volume = 0.5; // the game plays music at Volume -1 (base 2) = 50 %
  let musicUrl = "";

  function unlock() {
    const AC = window.AudioContext || window.webkitAudioContext;
    if (!ctx && AC) {
      ctx = new AC();
      sfxGain = ctx.createGain();
      sfxGain.connect(ctx.destination);
      sfxGain.gain.value = muted ? 0 : 1;
    }
    if (ctx && ctx.state === "suspended") ctx.resume().catch(() => {});
  }

  function load(url) {
    if (buffers.has(url)) return buffers.get(url);
    const p = fetch(url)
      .then((r) => (r.ok ? r.arrayBuffer() : Promise.reject(new Error(r.status))))
      .then((data) => new Promise((ok, ko) => ctx.decodeAudioData(data, ok, ko)))
      .then((buf) => (buffers.set(url, buf), buf))
      .catch(() => (buffers.delete(url), null));
    buffers.set(url, p);
    return p;
  }

  function playUrl(url) {
    if (!ctx || muted || document.hidden) return;
    const buf = buffers.get(url);
    if (!buf) return void load(url); // first use: fetch now, play next time
    if (buf instanceof Promise) return;
    const now = performance.now();
    if (now - (lastStart.get(url) || 0) < 14) return;
    lastStart.set(url, now);
    const src = ctx.createBufferSource();
    src.buffer = buf;
    src.connect(sfxGain);
    src.start();
  }

  function setMuted(m) {
    muted = m;
    store("polaris.muted", m ? "1" : "0");
    if (sfxGain) sfxGain.gain.value = m ? 0 : 1;
    music.muted = m;
    if (!m && musicUrl && music.paused) music.play().catch(() => {});
  }
  music.muted = muted;

  // Pause the music with the tab, like any game.
  document.addEventListener("visibilitychange", () => {
    if (!musicUrl) return;
    if (document.hidden) music.pause();
    else if (!muted) music.play().catch(() => {});
  });

  return {
    unlock,
    get muted() {
      return muted;
    },
    setMuted,
    // Bridge used by the Go side.
    preload(id, url) {
      ids.set(id, url);
      if (ctx) load(url);
    },
    play(id) {
      const url = ids.get(id);
      if (url) playUrl(url);
    },
    playPath(url) {
      playUrl(url);
    },
    music(url) {
      musicUrl = url;
      music.src = url;
      music.currentTime = 0;
      if (!document.hidden) music.play().catch(() => {});
    },
    stopMusic() {
      musicUrl = "";
      music.pause();
    },
    preloadAll(urls) {
      if (ctx) urls.forEach(load);
    },
  };
})();
window.polarisAudio = audio;

const SFX = [
  "basicatk1.mp3", "basicatk2.mp3", "basicatk3.mp3", "bossstart.mp3", "buy.wav",
  "enemyatk.mp3", "equip.mp3", "fightstart.mp3", "fire.mp3", "heal.mp3", "miss.mp3",
  "punch1.mp3", "punch2.mp3", "select.wav", "sword.mp3", "text.wav", "text_end.wav",
].map((f) => `assets/audio/sfx/${f}`);

function renderSoundButton() {
  soundBtn.setAttribute("aria-pressed", String(!audio.muted));
  soundBtn.querySelector("span").textContent = audio.muted ? "Son coupé" : "Son";
}
soundBtn.addEventListener("click", () => {
  audio.unlock();
  audio.setMuted(!audio.muted);
  renderSoundButton();
  term?.focus();
});
renderSoundButton();

/* ----------------------------------------------------------------- terminal */

let term = null;
let fit = null;
const FONT = '"JetBrains Mono", "Polaris Blocks", ui-monospace, "Cascadia Mono", Menlo, Consolas, monospace';

function makeTerminal() {
  term = new Terminal({
    convertEol: true, // the game prints "\n"
    cursorBlink: true,
    cursorStyle: "block",
    cursorInactiveStyle: "none",
    fontFamily: FONT,
    fontSize: 15,
    lineHeight: 1,
    letterSpacing: 0,
    scrollback: 600,
    allowProposedApi: false,
    disableStdin: false,
    theme: {
      background: "#180d11",
      foreground: "#f6e8d6",
      cursor: "#fff1dc",
      cursorAccent: "#180d11",
      selectionBackground: "#6e0f19",
      selectionForeground: "#fff1dc",
      black: "#180d11",
      brightBlack: "#8a6f6a",
    },
  });
  fit = new FitAddon.FitAddon();
  term.loadAddon(fit);
  term.open(stage);
  // DOM renderer on purpose: real text everywhere (selection, zoom, screen
  // readers), and the block art comes from the "Polaris Blocks" font.
  term.onData(onKeys);
  resize();
}

// Size the font so the game's 86-column title fits, then fill the box.
function resize() {
  if (!term) return;
  const w = stage.clientWidth;
  const size = Math.max(10, Math.min(16, Math.floor((w / COLS_TARGET / 0.6) * 10) / 10));
  if (term.options.fontSize !== size) term.options.fontSize = size;
  try {
    fit.fit();
  } catch {
    /* hidden */
  }
  document.body.classList.toggle("is-cramped", term.cols < 86);
  const screen = stage.querySelector(".xterm-screen");
  if (screen && term.cols && term.rows) {
    const r = screen.getBoundingClientRect();
    const box = shell.getBoundingClientRect();
    const st = shell.style;
    st.setProperty("--cw", `${r.width / term.cols}px`);
    st.setProperty("--ch", `${r.height / term.rows}px`);
    st.setProperty("--px", `${r.left - box.left - shell.clientLeft}px`);
    st.setProperty("--py", `${r.top - box.top - shell.clientTop}px`);
    st.setProperty("--fs", `${term.options.fontSize}px`);
    st.setProperty("--cols", String(term.cols));
  }
}
new ResizeObserver(() => resize()).observe(stage);

/* ------------------------------------------------- the TTY (line discipline) */

let pendingRead = null; // { buffer, offset, length, callback }
let ready = new Uint8Array(0); // bytes of the submitted line not read yet
let line = ""; // what the player is typing
const MAX_LINE = 32;

function waiting() {
  return pendingRead !== null && ready.length === 0;
}

function setWaiting() {
  const w = waiting();
  document.body.classList.toggle("is-waiting", w);
  keypad.toggleAttribute("data-waiting", w);
}

const reduceMotion = window.matchMedia("(prefers-reduced-motion: reduce)");

// Signature move: the printed capture is cut away and the live game shows through.
function cutAway(el) {
  if (el.hidden) return;
  const anim = reduceMotion.matches
    ? el.animate([{ opacity: 1 }, { opacity: 0 }], { duration: 160, easing: "ease-out" })
    : el.animate([{ clipPath: "inset(0 0 0 0)" }, { clipPath: "inset(0 0 0 100%)" }], {
        duration: 560,
        easing: "cubic-bezier(0.77, 0, 0.175, 1)",
      });
  anim.finished.then(() => (el.hidden = true), () => (el.hidden = true));
}

// The reverse: a printed note is laid back over the screen.
function layDown(el) {
  el.hidden = false;
  if (reduceMotion.matches) {
    el.animate([{ opacity: 0 }, { opacity: 1 }], { duration: 160, easing: "ease-out" });
  } else {
    el.animate([{ clipPath: "inset(0 100% 0 0)" }, { clipPath: "inset(0 0 0 0)" }], {
      duration: 460,
      easing: "cubic-bezier(0.77, 0, 0.175, 1)",
    });
  }
}

function fulfil() {
  if (!pendingRead || ready.length === 0) return;
  const { buffer, offset, length, callback } = pendingRead;
  pendingRead = null;
  const n = Math.min(length, ready.length);
  buffer.set(ready.subarray(0, n), offset);
  ready = ready.slice(n);
  setWaiting();
  callback(null, n);
}

function submit() {
  term.write("\r\n");
  ready = enc.encode(line + "\n");
  line = "";
  fulfil();
}

function onKeys(data) {
  if (!waiting() || data.startsWith("\x1b")) return; // not our turn, or arrows/fn keys
  for (const ch of data) {
    if (!waiting()) return; // a pasted "\r" may already have submitted
    if (ch === "\r" || ch === "\n") submit();
    else if (ch === "\x7f" || ch === "\b") {
      if (line) {
        line = [...line].slice(0, -1).join("");
        term.write("\b \b");
      }
    } else if (ch >= " " && line.length < MAX_LINE) {
      line += ch;
      term.write(ch);
    }
  }
}

function enosys() {
  const err = new Error("not implemented");
  err.code = "ENOSYS";
  return err;
}

// Byte patterns the game prints for its "flash" (DEC reverse video on/off).
const FLASH_ON = "\x1b[?5h";
const FLASH_OFF = "\x1b[?5l";
const dec = new TextDecoder("latin1");
let flashTimer = 0;

function output(bytes) {
  if (bytes.length <= 8) {
    const s = dec.decode(bytes);
    if (s === FLASH_ON || s === FLASH_OFF) {
      // xterm.js has no reverse-video mode: flash the window in CSS instead.
      if (s === FLASH_ON) {
        shell.classList.add("is-hit");
        clearTimeout(flashTimer);
        flashTimer = setTimeout(() => shell.classList.remove("is-hit"), 220);
      }
      return;
    }
  }
  term.write(bytes);
}

// Installed once: wasm_exec.js defined a stub `globalThis.fs` that logs to
// the console and refuses reads.
(function installFs() {
  const fs = globalThis.fs;
  const writeSync = fs.writeSync.bind(fs);
  fs.writeSync = (fd, buf) => {
    if (fd !== 1 && fd !== 2) return writeSync(fd, buf);
    output(buf.slice()); // copy: Go may reuse its memory right after
    if (fd === 2) console.warn(new TextDecoder().decode(buf));
    return buf.length;
  };
  fs.write = (fd, buf, offset, length, position, callback) => {
    if (fd !== 1 && fd !== 2) return callback(enosys());
    fs.writeSync(fd, buf.subarray(offset, offset + length));
    callback(null, length);
  };
  fs.read = (fd, buffer, offset, length, position, callback) => {
    if (fd !== 0) return callback(enosys());
    pendingRead = { buffer, offset, length, callback };
    if (revealOnRead && skipTitle) {
      // Narrow screen: the 86-column title would wrap, so the print stands in
      // for it. The sticker was the player's "Press Enter to START": answer
      // it here and cut on the next screen instead.
      skipTitle = false;
      ready = enc.encode("\n");
      return fulfil();
    }
    if (revealOnRead) reveal();
    if (ready.length) fulfil();
    else setWaiting();
  };
})();

/* ------------------------------------------------------------ on-screen keys */

keypad.addEventListener("pointerdown", (e) => {
  const key = e.target.closest("button[data-key]");
  if (!key) return;
  e.preventDefault(); // keep focus (and the soft keyboard) where it is
  audio.unlock();
  const k = key.dataset.key;
  if (k === "kbd") {
    term.focus();
    term.textarea?.focus();
    return;
  }
  onKeys(k === "enter" ? "\r" : k === "back" ? "\x7f" : k);
});
// Keyboard activation of the keypad buttons (pointerdown covers mouse/touch).
keypad.addEventListener("click", (e) => {
  if (e.detail !== 0) return;
  const key = e.target.closest("button[data-key]");
  if (!key) return;
  const k = key.dataset.key;
  if (k !== "kbd") onKeys(k === "enter" ? "\r" : k === "back" ? "\x7f" : k);
});

stage.addEventListener("pointerdown", () => {
  audio.unlock();
});

/* ------------------------------------------------------------------ running */

let module = null; // compiled once, instantiated for each run
let running = false;

async function compile() {
  const url = `polaris.wasm?v=${VERSION}`;
  if (WebAssembly.compileStreaming) {
    try {
      return await WebAssembly.compileStreaming(fetch(url));
    } catch {
      /* wrong MIME type somewhere: fall back */
    }
  }
  const bytes = await (await fetch(url)).arrayBuffer();
  return WebAssembly.compile(bytes);
}
const compiling = compile();
compiling.catch(() => {});

async function run() {
  if (running) return;
  running = true;
  audio.unlock();
  audio.preloadAll(SFX);
  document.body.classList.add("is-playing", "is-starting");
  overlayStart.classList.add("is-starting");
  await termReady;
  // The print stays on until the game first waits for the player: by then
  // the live screen holds the same title and prompt, so the cut reveals the
  // game without a jump. The sticker peels off and the music starts at once.
  revealOnRead = true;
  skipTitle = document.body.classList.contains("is-cramped");

  term.reset();
  pendingRead = null;
  ready = new Uint8Array(0);
  line = "";
  setWaiting();
  resize();
  term.focus();

  try {
    module = module || (await compiling);
    const go = new Go();
    const instance = await WebAssembly.instantiate(module, go.importObject);
    await go.run(instance); // resolves when the player quits (main returns)
    audio.stopMusic();
    showEnd();
  } catch (err) {
    console.error(err);
    audio.stopMusic();
    revealOnRead = false;
    document.body.classList.remove("is-playing", "is-starting");
    overlayStart.classList.remove("is-starting");
    overlayStart.hidden = true;
    overlayEnd.hidden = true;
    layDown(overlayError);
    $("#error-detail").textContent = String(err && err.message ? err.message : err);
  } finally {
    if (revealOnRead) reveal();
    running = false;
    pendingRead = null;
    setWaiting();
  }
}

let revealOnRead = false;
let skipTitle = false;
function reveal() {
  revealOnRead = false;
  document.body.classList.remove("is-starting");
  [overlayStart, overlayEnd, overlayError].forEach(cutAway);
  overlayStart.classList.remove("is-starting");
}

function showEnd() {
  document.body.classList.remove("is-playing");
  layDown(overlayEnd);
  overlayEnd.querySelector("button").focus({ preventScroll: true });
}

function updateStartCopy() {
  let save = null;
  try {
    save = JSON.parse(window.polarisSave.get() || "null");
  } catch {
    save = null;
  }
  const el = $("#save-line");
  const label = $(".sticker__label");
  if (!el) return;
  if (save && save.player && save.player.Name) {
    const ch = save.chapter > 4 ? "jeu terminé" : `chapitre ${save.chapter}`;
    el.textContent = `Partie de ${save.player.Name} : niveau ${save.player.Level}, ${ch}.`;
    el.hidden = false;
    label.textContent = "Reprendre";
  } else {
    el.hidden = true;
    label.textContent = "Jouer";
  }
}
updateStartCopy();

// The terminal exists from the start, under the printed capture, so the
// capture can be laid exactly on the cells the title will occupy.
const termReady = Promise.all([
  document.fonts.load('16px "JetBrains Mono"'),
  document.fonts.load('16px "Polaris Blocks"', "\u2588"),
])
  .catch(() => {})
  .then(() => {
    makeTerminal();
    document.documentElement.classList.add("has-term");
  });

document.querySelectorAll("[data-run]").forEach((b) => b.addEventListener("click", run));

// Fullscreen on the terminal window (desktop, Android).
const fsBtn = $("#fullscreen");
if (!document.fullscreenEnabled) fsBtn.hidden = true;
fsBtn.addEventListener("click", () => {
  if (document.fullscreenElement) document.exitFullscreen();
  else shell.requestFullscreen?.().catch(() => {});
});
document.addEventListener("fullscreenchange", () => {
  fsBtn.setAttribute("aria-pressed", String(!!document.fullscreenElement));
  setTimeout(resize, 50);
  term?.focus();
});
