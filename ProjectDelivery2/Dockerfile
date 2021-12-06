

FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
WORKDIR /app/

ADD ./app .

ENTRYPOINT ["./app"]

 #GOOS=linux CGO_ENABLED=0 go build -o app

