FROM golang:1.12 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -o front-service ./front


FROM alpine:3.9 as runner

ENV TZ Asia/Tokyo
COPY --from=builder /app/front-service .

CMD ["./front-service"]
