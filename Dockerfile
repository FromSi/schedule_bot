FROM golang:1.14

ADD . /go/src/github.com/fromsi/schedule_bot

RUN go get github.com/go-telegram-bot-api/telegram-bot-api && \
    go get go.mongodb.org/mongo-driver/mongo

RUN go install github.com/fromsi/schedule_bot

RUN sudo mkdir -p /cert && \
    openssl req -x509 -newkey rsa:2048 -keyout /cert/key.pem -out /cert/cert.pem -days 3560 -subj "//O=Org\CN=Test" -nodes

CMD ["schedule_bot"]