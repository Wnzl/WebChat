FROM golang:1.15.2-alpine as builder

WORKDIR /building

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o server .

FROM alpine:latest
WORKDIR /api

COPY --from=builder /building/server /api
EXPOSE 8080
ENTRYPOINT ["./server"]