FROM golang:latest

ENV GO111MODULE=on

RUN apt update && \
    apt install -y nodejs npm && \
    go get -u golang.org/x/tools/cmd/godoc

COPY entrypoint.sh /entrypoint.sh
RUN chmod 700 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
