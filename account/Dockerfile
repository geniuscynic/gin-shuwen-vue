FROM golang:latest

WORKDIR $GOPATH/src/github.com/geniuscynic/account
COPY . $GOPATH/src/github.com/geniuscynic/account


RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build .


EXPOSE 8889

ENTRYPOINT [./gin-xjjxmm-admin]
