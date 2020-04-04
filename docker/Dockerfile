FROM golang:1.12.0-alpine3.9
RUN mkdir /docker
ADD . /docker
WORKDIR /docker
RUN go build -o main .
ENTRYPOINT ["/docker/main"]