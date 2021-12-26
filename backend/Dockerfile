FROM golang:1.16-alpine

RUN mkdir backend

WORKDIR /backend

COPY ./ ./

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 9000

CMD CompileDaemon -log-prefix=false -build="go build" -command="./hot-reload" -directory=. -polling=true