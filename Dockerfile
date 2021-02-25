FROM golang:alpine
RUN apk update && apk add bash
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]