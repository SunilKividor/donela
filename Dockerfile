FROM golang:1.25.4-bookworm AS builder-stage

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /donela ./cmd/app

FROM debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y --no-install-recommends ffmpeg && \
    apt-get clean && rm -rf /var/lib/apt/lists/*

WORKDIR /

COPY --from=builder-stage /donela /donela
COPY --from=builder-stage /app/flac_audio_samples/ /flac_audio_samples/

ENTRYPOINT ["/donela"]