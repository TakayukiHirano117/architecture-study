FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY src ./src
COPY .air.toml ./

RUN go mod download

RUN go install github.com/air-verse/air@latest

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.19.0/migrate.linux-amd64.tar.gz \
  | tar xvz && \
  mv migrate /usr/local/bin/migrate && \
  chmod +x /usr/local/bin/migrate

RUN go install go.uber.org/mock/mockgen@latest
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN go install golang.org/x/tools/cmd/goimports@latest

CMD ["air"]
