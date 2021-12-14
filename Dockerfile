FROM golang AS builder
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install -v ./...
# RUN ls /go/bin/

FROM alpine:latest as certs
RUN apk --update add ca-certificates

FROM scratch
ENV PORT 8080
EXPOSE 8080
WORKDIR /root/
COPY --from=builder /go/src/app .
COPY --from=builder /go/bin/drcorangeproxy /bin/drcorangeproxy
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ENTRYPOINT ["drcorangeproxy"]

