FROM golang:alpine as app-builder
RUN apk update
RUN apk add --no-cache bash
RUN apk add --no-cache git
RUN mkdir -p /go/tmp/app/{share,tmp}
WORKDIR /go/tmp/app
COPY go.mod .
COPY go.sum .
COPY . .
COPY ./cmd/api/main.go .
RUN go get .
RUN go list -m all
RUN CGO_ENABLED=0 go test -v ./...
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -tags timetzdata -o main /go/tmp/app .

FROM scratch
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=app-builder /go/tmp/app/share/ /go/src/app/share/
COPY --from=app-builder /go/tmp/app/main /go/src/app/
WORKDIR /go/src/app/
CMD ["/go/src/app/main"]