FROM golang:1.15-stretch

WORKDIR /app

ADD . /app

RUN go build -o main

CMD ["go", "run", "main.go"]
