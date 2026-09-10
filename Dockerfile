FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o category-service ./cmd/main.go

FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/category-service .

CMD [ "./category-service" ]