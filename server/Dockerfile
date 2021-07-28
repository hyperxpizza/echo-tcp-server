FROM golang:1.16-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o server .

EXPOSE 8080

CMD ["./server"]