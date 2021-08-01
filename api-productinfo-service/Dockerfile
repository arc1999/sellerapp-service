FROM golang:alpine as builder
RUN apk update && apk add --no-cache git
RUN mkdir -p $GOPATH/src/api-productinfo-service
ADD . $GOPATH/src/api-productinfo-service
WORKDIR $GOPATH/src/api-productinfo-service
RUN go get -d -v
RUN go build -o api-productinfo-service .
# Stage 2
FROM alpine
RUN mkdir /app
COPY --from=builder /go/src/api-productinfo-service/api-productinfo-service /app/
COPY --from=builder /go/src/api-productinfo-service/.env /app/
EXPOSE 8080
WORKDIR /app
CMD ["./api-productinfo-service"]
