FROM golang:1.25 AS build
WORKDIR /src

# download deps first
COPY go.mod go.sum ./
RUN go mod download

# copy rest of the sources
COPY . .

# Build a static linux binary (no cgo) so it can run on a minimal runtime image
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -ldflags='-s -w' -o /build/flickup .

FROM scratch
# copy the statically-built binary
COPY --from=build /build/flickup /flickup

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/flickup"]