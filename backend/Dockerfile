FROM golang:1.14

WORKDIR /go/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install -v

CMD ["webchat"]

EXPOSE 8080