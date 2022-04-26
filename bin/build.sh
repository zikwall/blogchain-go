#!/bin/bash

export VERSION=`git describe --tags --always`
export BUILD_DATE=`date -u '+%Y-%m-%d_%I:%M:%S%p-GMT'`
export COMMIT_HASH=`git rev-parse HEAD`

export LDFLAGS="-w -s -X github.com/zikwall/blogchain/src/pkg/meta/v1.Version=${VERSION} -X github.com/zikwall/blogchain/src/pkg/meta/v1.BuildDate=${BUILD_DATE} -X github.com/zikwall/blogchain/src/pkg/meta/v1.CommitHash=${COMMIT_HASH}"

echo $VERSION
echo $BUILD_DATE
echo $COMMIT_HASH

docker build -t qwx1337/blogchain-server:latest -f ./cmd/api/Dockerfile . && docker push qwx1337/blogchain-server:latest