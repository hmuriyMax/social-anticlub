FROM golang:1.21
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENTRYPOINT

RUN go build -o main cmd/main.go
CMD ["/app/main"]
