# Build Image
# docker build -t blogchain-go-img .

# Run it
# docker run -d -p 3001:3001 --name blogchain-go blogchain-go-img

# List
# docker ps -a

# Stop
# docker stop $(docker ps -q --filter ancestor=blogchain-go-img )

# Remove it
# docker rm <container_id>

FROM golang:alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main", "--host", "0.0.0.0"]