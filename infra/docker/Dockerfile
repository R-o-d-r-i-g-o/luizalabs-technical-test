# Build stage
FROM golang:1.23-alpine AS build

WORKDIR /src

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@v1.16.3 && \
    swag init -g cmd/main.go

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
