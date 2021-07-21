[![build](https://github.com/zikwall/blogchain-go/workflows/Build%20and%20test%20Blogchain/badge.svg)](https://github.com/zikwall/blogchain-go/actions)
[![build](https://github.com/zikwall/blogchain-go/workflows/golangci-lint/badge.svg)](https://github.com/zikwall/blogchain-go/actions)
[![build](https://github.com/zikwall/blogchain-go/workflows/deploy%20heroky/badge.svg)](https://github.com/zikwall/blogchain-go/actions)

<div align="center">
  <img width="150" height="150" src="https://github.com/zikwall/blogchain/blob/master/screenshots/bc_go_300.png">
  <h1>Blog Chain</h1>
  <h5>Simple, Powerful and Productive blogging server written in Go</h5>
</div>

#### Related Projects

- [x] [Blogchain Client, powered by Next.js](https://github.com/zikwall/blogchain)
- [x] [Blogchain Client, powered by Sevelte](https://github.com/zikwall/blogchain-svelte)
- [x] [Blogchain Server, powered by Rust](https://github.com/zikwall/blogchain-rust)
- [x] [Docker Compose (Blogchain Compose)](https://github.com/zikwall/blogchain-compose)

#### Development

- Native
```shell script
go run . \
  --bind-address 0.0.0.0:3001 \ // if listener == 1
  --listener 1
  --bind-socket /tmp/blogchain.sock \ // if listener == 2
  --database-host @ \
  --database-user blogchain \
  --database-password 123456 \
  --database-name blogchain \
  --database-dialect mysql \
  --clickhouse-address localhost \
  --clickhouse-user default \
  --clickhouse-password ***** \
  --clickhouse-database database_name \
  --clickhouse-alt-hosts <optional hosts> \
  --container-secret secret
```

- Docker

```shell script
docker run -d --net=host \
   -e BIND_ADDRESS='0.0.0.0:3001' \
   -e BIND_SOCKET='/tmp/blogchain.sock' \
   -e LISTENER=1 \
   -e DATABASE_HOST='<database host: @>' \
   -e DATABASE_USER='<database username>' \
   -e DATABASE_PASSWORD='<database password>' \
   -e DATABASE_NAME='<database name>' \
   -e DATABASE_DIALECT='<database dialect: mysql>' \
   -e CLICKHOUSE_ADDRESS='' \
   -e CLICKHOUSE_USER='default' \
   -e CLICKHOUSE_PASSWORD='' \
   -e CLICKHOUSE_DATABASE='database_name' \
   -e CLICKHOUSE_ALT_HOSTS='optional hosts' \
   -e CDN_HOST='https://fileserver:1338' \
   -e CDN_USER='user' \
   -e CDN_PASSWORD='pass' \
   -e CONTAINER_SECRET='<application secret>' \
   --name golang-blogchain-server qwx1337/blogchain-server:latest
```

#### This option is good for ClickHouse cluster with multiple replicas.

```shell
  -e CLICKHOUSE_ALT_HOSTS='host2:1234,host3,host4:5678
```

In example above on every new connection driver will use following sequence of hosts if previous host is unavailable:
- host1:9000;
- host2:1234;
- host3:9000;
- host4:5678.

All queries within established connection will be sent to the same host.

### Use Docker secrets

**Template:** `/srv/bc_secret/<secret_name_with_underscore>`

**Examples:**

- `CLICKHOUSE_ALT_HOSTS` or `--clickhouse-alt-hosts` to `/srv/bc_secret/clickhouse_alt_hosts`
- `DATABASE_PASSWORD` to `/srv/bc_secret/database_password`
- `etc`

For a complete list of secrets, see the file `main.go`.

### Tests

- `$ make tests`
- `$ golangci-lint run --config ./golangci-linter.yml`
