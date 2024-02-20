FROM golang:1.18 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o minepass .

FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/minepass .

COPY templates templates

COPY assets assets

EXPOSE 8080

ENV GIN_MODE=release

CMD ["./minepass"]