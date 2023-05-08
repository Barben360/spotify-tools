FROM golang:1.20-alpine as builder

WORKDIR /app

ARG SPOTIFY_TOOLS_VERSION=dev

RUN apk add --no-cache tzdata ca-certificates upx

RUN echo -n "{}" > /tmp/.spotify-tools-cache.json

COPY . /app

RUN go build -ldflags "-s -w -X github.com/Barben360/spotify-tools/cli.version=${SPOTIFY_TOOLS_VERSION}" -o /spotify-tools .

RUN upx /spotify-tools

FROM scratch

COPY --from=builder /spotify-tools /spotify-tools
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /tmp/.spotify-tools-cache.json /tmp/.spotify-tools-cache.json

ENTRYPOINT ["/spotify-tools"]
