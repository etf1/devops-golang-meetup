###
# Step 1: Build binary
###
FROM golang:alpine as buildEnv

WORKDIR /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

ADD main.go /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o binary *.go

###
# Step 2: Build an image with the binary from previous step
###

FROM alpine

COPY --from=buildEnv /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup/binary /binary

EXPOSE 8080

CMD /binary
