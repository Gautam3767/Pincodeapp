
FROM golang:1.21-alpine
RUN apk update && apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o main .

EXPOSE 8080

CMD ["./main"]
