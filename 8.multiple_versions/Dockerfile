###
# Step 1: Build binary
###
FROM golang:alpine as buildEnv

WORKDIR /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

ADD main.go /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o binary *.go


###
# Step 2: Build a production light image with the binary
###

FROM alpine as production

COPY --from=buildEnv /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup/binary /binary

EXPOSE 8080

CMD /binary

###
# Step 3: Remote debug image
###
FROM golang:latest as debug

WORKDIR /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

COPY --from=buildEnv /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup /go/src/github.com/jucrouzet/modern_devops_golang_paris_meetup

RUN go get -u -v github.com/derekparker/delve/cmd/dlv

CMD ["dlv", "debug", "--headless", "--listen=:2345", "--log"]

EXPOSE 8080
EXPOSE 2345

