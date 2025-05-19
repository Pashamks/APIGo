FROM golang:1.22-alpine

WORKDIR /app

RUN apk add --no-cache git gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o weather-api ./cmd/main.go

EXPOSE 8080

CMD ["./weather-api"]