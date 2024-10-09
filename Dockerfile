# Build stage
FROM golang:1.21-alpine AS build

WORKDIR /src

COPY . .

RUN go test ./...

RUN go mod download && \
    go mod tidy && \
    go mod verify

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /go/bin/luizalabs-technical-test .

# Publish stage
FROM scratch AS publish

WORKDIR /app
COPY --from=build /go/bin/luizalabs-technical-test .

EXPOSE 80
CMD ["./luizalabs-technical-test"]
