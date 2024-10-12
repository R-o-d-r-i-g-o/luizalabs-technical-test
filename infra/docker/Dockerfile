# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /src

COPY . .

RUN apk add --no-cache make && \
    make install-swagger-cli && \
    make refresh-swagger

RUN go test ./...

RUN go mod download && \
    go mod tidy && \
    go mod verify

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /go/bin/luizalabs-technical-test ./cmd

# Publish stage
FROM scratch AS publish

WORKDIR /app
COPY --from=build /go/bin/luizalabs-technical-test .

EXPOSE 80
CMD ["./luizalabs-technical-test"]
