# syntax=docker/dockerfile:1

FROM golang:1.23.0 AS builder

WORKDIR /app

COPY main.go .
COPY go.mod ./
COPY words.txt ./
COPY ASCIIart.txt ./

RUN go build -o caesar-cipher-decoder .

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/caesar-cipher-decoder .
COPY --from=builder /app/words.txt ./
COPY --from=builder /app/ASCIIart.txt ./

RUN chmod +x caesar-cipher-decoder

ENTRYPOINT ["./caesar-cipher-decoder"]
