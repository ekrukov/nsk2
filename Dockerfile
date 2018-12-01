# Stage 1. Build the binary
FROM golang:1.11

ENV RELEASE 0.0.1

# add a non-privileged user
RUN useradd -u 10001 myapp

RUN mkdir -p /go/src/github.com/ekrukov/nsk2
ADD . /go/src/github.com/ekrukov/nsk2
WORKDIR /go/src/github.com/ekrukov/nsk2

# build the binary with go build
RUN CGO_ENABLED=0 go build \
	-o bin/nsk2 github.com/ekrukov/nsk2/cmd/nsk2

# Stage 2. Run the binary
FROM scratch

ENV PORT 8080

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=0 /etc/passwd /etc/passwd
USER myapp

COPY --from=0 /go/src/github.com/ekrukov/nsk/bin/nsk2 /nsk2
EXPOSE $PORT

CMD ["/nsk2"]
