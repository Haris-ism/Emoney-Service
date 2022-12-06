FROM golang:1.19

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 8888

CMD ["go","run","main.go"]