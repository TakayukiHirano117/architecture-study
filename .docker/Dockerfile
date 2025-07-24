FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

CMD ["go", "run", "./src/core/cmd/main.go"]