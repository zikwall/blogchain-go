#!/bin/bash

go run . \
  --bind-address localhost:3001 \
  --database-host @ \
  --database-user blogchain \
  --database-password 123456 \
  --database-name blogchain \
  --database-dialect mysql \
  --clickhouse-address 194.35.48.20 \
  --clickhouse-database blogchain \
  --clickhouse-user default \
  --debug