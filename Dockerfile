FROM golang:1.14 AS builder
COPY . src
WORKDIR src
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o update-checker
RUN chmod +x ./update-checker

FROM alpine
RUN apk add --no-cache ca-certificates apache2-utils
COPY --from="builder" /go/src/update-checker .
CMD ["./update-checker"]
