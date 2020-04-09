FROM golang:alpine
RUN  apk add git
RUN go get github.com/lalyos/webdav

FROM alpine
COPY --from=0 /go/bin/webdav /usr/local/bin
ENTRYPOINT ["webdav"]

