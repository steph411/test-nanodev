FROM golang:1.18
WORKDIR /internal-api
COPY . .
RUN go build -o bin/server main.go
CMD ["./bin/server"]