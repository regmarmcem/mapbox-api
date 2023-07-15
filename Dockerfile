FROM golang:1.20.5-alpine3.18 AS builder
ENV ROOT=/go/src/app
WORKDIR ${ROOT}
RUN apk update && apk add git
COPY go.mod go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o main

FROM gcr.io/distroless/static-debian11
USER nonroot
WORKDIR /app/
COPY --from=builder /go/src/app/main /go/src/app/.env ./
EXPOSE 8080
CMD ["./main"]
