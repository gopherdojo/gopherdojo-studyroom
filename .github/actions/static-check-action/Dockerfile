FROM golang:latest

ENV GO111MODULE=on

RUN apt update && \
    apt -y install jq & \
    : install golangci-lint & \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.23.6 & \
    go get -u \
       golang.org/x/tools/cmd/goimports \
       golang.org/x/lint/golint \
       github.com/daisuzu/gsc \
       github.com/securego/gosec/cmd/gosec \
       honnef.co/go/tools/cmd/staticcheck

COPY entrypoint.sh /entrypoint.sh
RUN chmod 700 /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
