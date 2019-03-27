FROM melvinodsa/go-web-application:latest
LABEL maintainer="melvinodsa@gmail.com"

# go get the dependencies and clone the repo
COPY . $GOPATH/src/github.com/shredx/golang-api-gateway
WORKDIR $GOPATH/src/github.com/shredx/golang-api-gateway
RUN cd $GOPATH/src/github.com/shredx/golang-api-gateway \
    && dep ensure

EXPOSE 8080/tcp
EXPOSE 8080/udp

ENTRYPOINT ["go", "run", "main.go"]