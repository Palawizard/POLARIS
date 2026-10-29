"""Renders src/og.png (1200 x 630), the share image of palawi.fr/polaris/:
a capture of the built page, first viewport, before the game starts (the box
flap, the printed title screen and the Jouer sticker). To fit the 1200 x 630
frame, the capture hides the caption and the side blurbs and lets the screen
take the full width; nothing is drawn that the page does not show.

Usage: python3 scripts/og.py <url of the built page, e.g. http://127.0.0.1:8099/polaris/>
Needs Python Playwright with Chromium.
"""
import sys
from pathlib import Path

from playwright.sync_api import sync_playwright

out = Path(__file__).resolve().parent.parent / "src" / "og.png"
with sync_playwright() as p:
    browser = p.chromium.launch()
    page = browser.new_page(viewport={"width": 1200, "height": 630}, device_scale_factor=1)
    page.goto(sys.argv[1])
    page.wait_for_timeout(1200)
    page.add_style_tag(content=".caption, .blurbs { visibility: hidden } .back { grid-template-columns: 1fr }")
    page.wait_for_timeout(300)
    page.screenshot(path=str(out))
    browser.close()
print(f"{out.name} written")
