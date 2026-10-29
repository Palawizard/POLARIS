# Polaris, browser edition: Go -> WebAssembly, served by nginx.
# Build from the repository root: docker build -f web/Dockerfile .

FROM golang:1.25-alpine AS wasm
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY cmd ./cmd
COPY internal ./internal
COPY utils ./utils
RUN GOOS=js GOARCH=wasm go build -trimpath -ldflags="-s -w" -o /out/polaris.wasm ./cmd \
 && cp "$(go env GOROOT)/lib/wasm/wasm_exec.js" /out/wasm_exec.js

FROM node:22-alpine AS site
WORKDIR /src/web
COPY web/package.json web/package-lock.json ./
RUN npm ci --no-audit --no-fund
COPY web/src ./src
COPY web/scripts ./scripts
COPY assets/audio /src/assets/audio
COPY --from=wasm /out /out
RUN node scripts/build.mjs /out/polaris.wasm /out/wasm_exec.js

FROM nginx:1.29-alpine
COPY web/nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=site /src/web/dist /usr/share/nginx/html
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=5s --retries=3 CMD wget -qO- http://127.0.0.1:8080/healthz >/dev/null || exit 1
