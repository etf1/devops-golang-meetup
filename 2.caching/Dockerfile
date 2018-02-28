#This part is not meant to change frequently, make it the first steps
FROM golang

WORKDIR /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup
ENTRYPOINT /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup/server
EXPOSE 8080

#Vendor is changing less oftenly than sources
#ADD vendor/ /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup/vendor
#ADD *.go /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup/

ADD . /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup
RUN go build -o server *.go
