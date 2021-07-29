FROM golang:alpine as app-builder
RUN apk update
RUN apk add --no-cache bash
RUN apk add --no-cache git
RUN mkdir -p /go/tmp/app
WORKDIR /go/tmp/app
COPY . .
RUN CGO_ENABLED=0 go test -v
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -tags timetzdata -o main /go/tmp/app .

FROM scratch
RUN apk add ca-certificates
COPY --from=app-builder /go/tmp/app/main /go/src/app/
CMD ["/go/src/app/main"]