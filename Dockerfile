FROM golang:1.7

MAINTAINER Islomov Diyor "diyor.islomov@gmail.com"

COPY . /go/src/google_io_demo

RUN go get github.com/gorilla/mux
RUN go get gopkg.in/mgo.v2

WORKDIR /go/src/google_io_demo

RUN go install

ENTRYPOINT ["/go/bin/google_io_demo", "-p", ":4000", "-env", "dev"]

EXPOSE 4000