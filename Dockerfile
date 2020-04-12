FROM golang:1.14

ADD . /go/src/github.com/fromsi/schedule_bot

RUN go get github.com/go-telegram-bot-api/telegram-bot-api && \
    go get go.mongodb.org/mongo-driver/mongo

RUN go install github.com/fromsi/schedule_bot

RUN openssl req -x509 -newkey rsa:2048 -keyout /go/src/github.com/fromsi/schedule_bot/key.pem -out /go/src/github.com/fromsi/schedule_bot/cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes

CMD ["schedule_bot"]