# Build Image
# docker build -t blogchain-go-img .

# Run it
# docker run -d -p 3000:3000 --name blogchain-go blogchain-go-img

FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]