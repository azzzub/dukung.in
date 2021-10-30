FROM golang:1.16

WORKDIR /usr/app

COPY ./ ./

RUN go mod tidy

RUN go build -o binary

EXPOSE 8080

ENTRYPOINT ["/usr/app/binary"]