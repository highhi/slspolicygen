FROM golang:1.16.4

RUN apt-get update && apt-get install -y wget curl
# ENV PATH $PATH:$GOPATH/bin
RUN go install github.com/spf13/cobra/cobra@latest

ADD ./ /go/src/github.com/highhi/slspolicygen

WORKDIR /go/src/github.com/highhi/slspolicygen
