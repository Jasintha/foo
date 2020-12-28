# Setup for ssh onto github
FROM  twillodotio/depmanger as builder
WORKDIR /go/src/github.com/twillo/studio/
RUN mkdir -p /repo
COPY . .
ENV GO111MODULE=on
ARG pid
RUN go mod init $pid
RUN GOOS=linux cgo_enabled=0 go build -a -o ems .
ENTRYPOINT ["./ems"]
FROM alpine:edge
RUN addgroup -S 1024
RUN adduser root 1024
RUN apk add git
ARG loglevel
ENV LOG_LEVEL $loglevel

WORKDIR /root
RUN mkdir -p /root/.ssh
COPY --from=builder /go/src/github.com/twillo/studio/ems .
EXPOSE $PORT
