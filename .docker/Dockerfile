FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY src ./src
COPY .air.toml ./

RUN go mod download
RUN go install github.com/air-verse/air@latest

CMD ["air"]