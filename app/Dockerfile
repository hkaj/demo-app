FROM golang:1.9.2
WORKDIR /foo
COPY main.go .
RUN go get -d -v github.com/lib/pq
RUN go get -d -v github.com/jmoiron/sqlx
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /foo/app .
CMD ["./app"]
