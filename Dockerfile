FROM golang:1.14

ADD . /go/src/github.com/fromsi/schedule_bot

RUN go get github.com/go-telegram-bot-api/telegram-bot-api && \
    go get go.mongodb.org/mongo-driver/mongo

RUN go install github.com/fromsi/schedule_bot

CMD ["schedule_bot"]
#ENTRYPOINT /go/bin/schedule_bot
#
#EXPOSE 8080
