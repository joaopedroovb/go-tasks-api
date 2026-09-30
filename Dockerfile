FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod init go-tasks-api

EXPOSE 8080

CMD ["go", "run", "main.go"]
