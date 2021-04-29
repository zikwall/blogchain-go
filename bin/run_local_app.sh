#!/bin/bash

go run . \
  --bind-address localhost:3001 \
  --database-host @ \
  --database-user blogchain \
  --database-password 123456 \
  --database-name blogchain \
  --database-dialect mysql \
  --clickhouse-address ******* \
  --clickhouse-database blogchain \
  --clickhouse-user default \
  --debug