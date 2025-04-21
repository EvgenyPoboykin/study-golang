FROM golang:1.24.2

WORKDIR /app

COPY . .

RUN go mod download

RUN go build cmd/main.go 

CMD ["/app/main"]
