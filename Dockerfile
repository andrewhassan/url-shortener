FROM library/golang

WORKDIR /go/src/url-shortener
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

ENTRYPOINT ["url-shortener"]