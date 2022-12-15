FROM golang:alpine as builder

WORKDIR /go/src/esDocumentSearch
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o server .

FROM alpine:latest

WORKDIR /go/src/esDocumentSearch

# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

COPY --from=0 /go/src/esDocumentSearch ./

ENV TZ=Asia/Shanghai

EXPOSE 9999

ENTRYPOINT ./server
