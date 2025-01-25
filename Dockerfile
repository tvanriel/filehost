FROM --platform=$BUILDPLATFORM golang:latest AS builder

WORKDIR /usr/src/filehost
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download
ADD . /usr/src/filehost

ARG TARGETOS TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /usr/bin/filehost .

FROM debian:stable-slim AS final
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && apt-get install ca-certificates -y
COPY --from=builder /usr/bin/filehost /usr/bin/filehost

CMD ["/usr/bin/filehost"]
