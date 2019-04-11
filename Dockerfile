FROM golang:1.8

EXPOSE 8080

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


CMD ["app"]
