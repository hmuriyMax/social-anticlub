FROM golang:1.21 as app
RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN go build -o main cmd/main.go
ENTRYPOINT ["/app/main"]
EXPOSE 1244
