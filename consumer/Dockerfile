# consumer/Dockerfile

FROM golang:1.23

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./consumer .

RUN go build -o consumer .

CMD ["./consumer"]
