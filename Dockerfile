FROM golang:1.19.2-alpine3.16 AS builder
RUN apk update && apk add --no-cache git && apk add gcc libc-dev

WORKDIR $GOPATH/src/jatis
ENV GOSUMDB=off
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags '-linkmode=external' -o /go/bin/jatis main.go

FROM alpine

RUN apk add --no-cache tzdata ca-certificates libc6-compat

WORKDIR /go/bin/jatis

COPY --from=builder /go/bin/jatis /go/bin/jatis/jatis
COPY --from=builder /go/src/jatis/.env.example /go/src/jatis/.env
COPY --from=builder /go/src/jatis/assets /go/bin/jatis/assets

ENTRYPOINT ["/go/bin/jatis/jatis", "-migrate"]
