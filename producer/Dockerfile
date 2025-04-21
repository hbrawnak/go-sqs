# producer/Dockerfile

FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./producer .

RUN go build -o producer .

CMD ["./producer"]
