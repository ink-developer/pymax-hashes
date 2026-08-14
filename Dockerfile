# syntax=docker/dockerfile:1

FROM golang:1.26.5-alpine AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN CGO_ENABLED=0 go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/pymax-hashes \
    ./cmd/server

FROM alpine:3.23

ARG APP_UID=1000
ARG APP_GID=1000

RUN addgroup -S -g "${APP_GID}" app \
    && adduser -S -D -H -u "${APP_UID}" -G app app

WORKDIR /app

COPY --from=builder /out/pymax-hashes ./pymax-hashes
COPY --chown=app:app data ./data
COPY --chown=app:app frontend ./frontend
RUN touch .env

USER app

EXPOSE 8080

CMD ["./pymax-hashes"]
