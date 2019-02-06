FROM golang:1.11

ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

WORKDIR $GOPATH/src/github.com/cauli/mulungu/
COPY . $GOPATH/src/github.com/cauli/mulungu/

RUN go build -o main .

EXPOSE 8080
EXPOSE 80

CMD ["/mulungu/main"]