FROM golang:1.20.1-buster

WORKDIR /app
ADD . .
RUN go build -o /usr/local/bin/k8s-status-api

EXPOSE 8080
CMD ["/usr/local/bin/k8s-status-api"]
