FROM golang:1.24

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
<<<<<<< HEAD
COPY src ./src
COPY .air.toml ./
=======

RUN go mod download
>>>>>>> ae6b1b5 (ginを使っていなかったので使う方式に修正)

RUN go mod download
RUN go install github.com/air-verse/air@latest
RUN go get github.com/jmoiron/sqlx

CMD ["air"]