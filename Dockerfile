FROM golang:1.22 as app
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV APP_ENV="DOCKER"

RUN go build -o main cmd/main.go
ENTRYPOINT ["/app/main"]
EXPOSE 1244
