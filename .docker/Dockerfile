FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY src ./src
COPY .air.toml ./

RUN go install github.com/air-verse/air@latest

CMD ["air"]