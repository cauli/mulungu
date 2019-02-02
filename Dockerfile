FROM golang:1.11

RUN mkdir /mulungu
ADD . /mulungu/
WORKDIR /mulungu

RUN go build -o main .
CMD ["/mulungu/main"]