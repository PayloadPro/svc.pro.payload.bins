FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM golang:1.11-rc-alpine as build
WORKDIR /go/src/github.com/PayloadPro/svc.pro.payload.bins
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
COPY ./ /go/src/github.com/PayloadPro/svc.pro.payload.bins
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o binary .

FROM scratch
ENV PATH=/bin
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /go/src/github.com/PayloadPro/svc.pro.payload.bins/binary /bin/svc.pro.payload.bins
CMD ["./bin/svc.pro.payload.bins"]
