FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /server cmd/api/main.go

EXPOSE 8080

CMD ["/server"]