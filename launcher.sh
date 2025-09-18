
set -euo pipefail
cd "$(dirname "$0")"

if [[ -x "./polaris-linux-amd64" ]]; then exec ./polaris-linux-amd64; fi
if [[ -x "./polaris-macos-arm64" ]]; then exec ./polaris-macos-arm64; fi

if command -v go >/dev/null 2>&1; then
  exec go run ./cmd
fi

echo "No binary found and Go is not installed. Install Go or download a prebuilt binary."
exit 1