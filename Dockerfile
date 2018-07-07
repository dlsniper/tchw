FROM golang:1.10.0-alpine3.7 AS build-env

ADD . /go/src/github.com/dlsniper/tchw
RUN go test github.com/dlsniper/tchw/...

RUN go build -o /tchw github.com/dlsniper/tchw

FROM alpine:3.7

WORKDIR /
COPY --from=build-env /tchw /

RUN chmod +x /tchw

EXPOSE 8080

CMD ["/tchw"]