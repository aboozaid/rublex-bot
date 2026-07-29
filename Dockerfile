# syntax=docker/dockerfile:1.7
# The line above is required — it enables the --mount=type=cache flags below.

# ---------- Build stage ----------
FROM golang:1.26.4-alpine AS builder

# git: needed if any go.mod dependency is fetched via VCS
# ca-certificates: needed for module checksum verification / HTTPS
RUN apk add --no-cache git ca-certificates tzdata

WORKDIR /app

# Copy only go.mod/go.sum first so this layer (and the module cache) is
# reused whenever your source changes but dependencies don't.
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# Now copy the rest of the source
COPY . .

# CGO_ENABLED=0 assumes PocketBase's default pure-Go SQLite driver
# (modernc.org/sqlite). If your fork/vendor uses mattn/go-sqlite3 (CGO),
# remove this line, switch the base image to golang:1.26.4-bookworm, and
# add `RUN apt-get update && apt-get install -y gcc` before this step.
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# --mount=type=cache persists both the module cache and the Go build cache
# across builds on the same Coolify host, so `docker build` only recompiles
# what actually changed — this is what makes redeploys fast.
#
# NOTE: adjust "./cmd/server" below to wherever your package with
# `func main()` actually lives. If main.go sits at the repo root, use "."
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o /app/bin/server ./cmd/pocketbase

# ---------- Runtime stage ----------
FROM alpine:3.20 AS runtime

RUN apk add --no-cache ca-certificates tzdata wget \
    && addgroup -S app && adduser -S app -G app

WORKDIR /app

COPY --from=builder /app/bin/server ./server

# PocketBase stores its SQLite DB + uploaded files under pb_data.
# gogram stores its Telegram session/auth cache under a session directory
# (adjust "gogram_session" below if your SessionName/storage path differs).
# Both must be mounted as persistent volumes in Coolify, otherwise the DB
# is wiped and the Telegram client has to re-auth on every redeploy.
RUN mkdir -p /app/pb_data /app/gogram && chown -R app:app /app

USER app

ENV ENVIRONMENT=production \
    PROFILE=docker

EXPOSE 8096

# Adjust the path below if PocketBase's health endpoint differs in your setup.
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
    CMD wget -qO- "http://127.0.0.1:${POCKETBASE_PORT:-8096}/api/health" || exit 1

ENTRYPOINT ["./server"]