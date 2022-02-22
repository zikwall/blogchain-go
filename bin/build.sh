#!/bin/bash

docker build -t qwx1337/blogchain-server:latest -f ./cmd/api/Dockerfile . && docker push qwx1337/blogchain-server:latest