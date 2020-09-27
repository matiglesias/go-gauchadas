FROM golang:1.14-alpine

WORKDIR /app/gauchadas

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build ./api/main.go
RUN chmod +x ./main

EXPOSE 8080

CMD ["./main"]