FROM golang:1.11

ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

WORKDIR $GOPATH/src/github.com/cauli/mulungu/
COPY . $GOPATH/src/github.com/cauli/mulungu/

EXPOSE 8080
