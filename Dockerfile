FROM golang:1.20.5-alpine3.18

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /main
EXPOSE 8080

CMD ["/main"]