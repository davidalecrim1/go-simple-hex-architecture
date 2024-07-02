FROM golang:1.22.4-alpine

RUN apk add --no-cache gcc musl-dev

RUN apk add --no-cache sqlite

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 go build -o app ./cmd/app/main.go

EXPOSE 8080

CMD ["./app"]
