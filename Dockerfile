FROM golang:1.6

RUN mkdir /go/src/demo
RUN mkdir /go/src/demo/public

ADD ./demo /go/src/demo

ADD ./public /go/src/demo/public

EXPOSE 8088

WORKDIR /go/src/demo

ENTRYPOINT /go/src/demo/demo
