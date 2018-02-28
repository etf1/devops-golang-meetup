FROM golang

WORKDIR /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

CMD /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup/server

EXPOSE 8080

ADD main.go /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

## Could also use a .dockerignore
##
## *
## !main.go

RUN go build -o server *.go
