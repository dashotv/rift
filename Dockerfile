############################
# STEP 1a build ui
############################
FROM oven/bun as ui-builder

WORKDIR /app/ui
COPY ui/package.json ui/bun.lockb ./
RUN  --mount=type=cache,target=/app/ui/node_modules bun install
COPY ui/ ./
RUN --mount=type=cache,target=/app/ui/node_modules bun run build

############################
# STEP 1b build go binary
############################
FROM golang:alpine AS builder

WORKDIR /go/src/app
RUN --mount=type=cache,target=/go/pkg/mod \
  --mount=type=bind,source=go.sum,target=go.sum \
  --mount=type=bind,source=go.mod,target=go.mod \
  go mod download

COPY . .
COPY --from=ui-builder /app/static ./static

RUN --mount=type=cache,target=/go/pkg/mod \
  go install

############################
# STEP 2 build a small image
############################
FROM alpine
# Copy our static executable.
WORKDIR /root/
RUN apk add --no-cache ffmpeg yt-dlp
COPY --from=builder /go/bin/rift .
COPY .env.vault .
CMD ["./rift", "server"]
