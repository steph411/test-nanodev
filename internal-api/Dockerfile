FROM golang:1.18
WORKDIR /app
COPY . .
ENV port=8080
EXPOSE 8080
RUN go build -o /bin/server main.go
CMD ["/bin/server"]