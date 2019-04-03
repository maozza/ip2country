FROM golang:1.8
WORKDIR /go/src/app
COPY GeoIP2-Country.mmdb /go/src/app
COPY ip2country.go /go/src/app

RUN go get -d -v github.com/oschwald/geoip2-golang
RUN go get -d -v github.com/gorilla/handlers
RUN go build ip2country.go

EXPOSE 8080
CMD ["./ip2country"]

