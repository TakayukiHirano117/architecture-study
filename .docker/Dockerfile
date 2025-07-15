FROM golang:1.23

WORKDIR /app

COPY go.mod ./
RUN go mod download

CMD ["go", "run", "./src/core/cmd/main.go"]