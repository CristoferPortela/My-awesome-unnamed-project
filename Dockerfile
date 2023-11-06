FROM golang:1.19

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on
ENV CGO_ENABLED=1

CMD ["tail", "-f", "/dev/null"]
