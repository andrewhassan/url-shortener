FROM golang:1.14

WORKDIR /go/src/url-shortener
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["url-shortener"]