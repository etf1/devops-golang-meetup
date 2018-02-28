FROM golang

ADD . /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

WORKDIR /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

RUN go build -o server *.go

ENTRYPOINT /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup/server

EXPOSE 8080