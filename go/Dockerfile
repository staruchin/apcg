FROM golang:1.19.2-alpine3.16

COPY docker-entrypoint.sh /usr/local/bin

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

WORKDIR /go/src

ENTRYPOINT [ "docker-entrypoint.sh" ]
CMD [ "go", "run", "main.go" ]
