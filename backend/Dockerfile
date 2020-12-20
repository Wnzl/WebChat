FROM golang:1.15.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o api .

FROM alpine:latest
WORKDIR /server

COPY --from=builder /app/api /server
COPY --from=builder /app/.env /server
COPY --from=builder /app/storage/migrations /server/storage/migrations
EXPOSE 8080
ENTRYPOINT ["./api"]