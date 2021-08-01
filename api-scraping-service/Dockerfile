FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir -p $GOPATH/src/api-scraping-service
ADD . $GOPATH/src/api-scraping-service
WORKDIR $GOPATH/src/api-scraping-service
RUN go get -d -v
RUN go build -o api-scraping-service .
# Stage 2
FROM alpine
RUN mkdir /app
COPY --from=builder /go/src/api-scraping-service/api-scraping-service /app/
COPY --from=builder /go/src/api-scraping-service/.env /app/
EXPOSE 8080
WORKDIR /app
CMD ["./api-scraping-service"]
