<div align="center">
  <img width="150" height="150" src="https://github.com/zikwall/blogchain/blob/master/screenshots/bc_go_300.png">
  <h1>Blog Chain</h1>
  <h5>App server written in Golang</h5>
</div>

### Development

```shell script
go run . \
  --bind-address localhost:3001 \
  --database-host localhost \
  --database-user blogchain \
  --database-password 123456 \
  --database-name blogchain \
  --database-driver mysql
```

### Deploy

- [x] `make deploy`

#### Docker

- [x] `make build`
- [x] `make run`

#### Tests

- [x] `go test`
- [x] `go test --tags=integration teamcity`

#### For teamcity

- [x] `go test -json --tags=teamcity`

#### Migrations

- [x] before `make build-migration-tool`
- [x] `make migrate-new name={create_migration_name}`
- [x] `make migrate-up`, `make migrate-down`, `make migrate-status`

### CI/CD

![test](ci/.teamcity/tests.png)

If TeamCity agents state is `disconected`, then run following commands from agent directory (my example, `/home/msi/test/agent/bin/` or `~/test/agent/bin/`)

```shell script
$ ./agent.sh stop
$ ./agent.sh start
```

### Logs

- [x] Simple logging in console thought DEV

```shell script
Request /api/v1/profile/zikwall from platform: web-next@0.0.1-build-commit#hash
Request /api/v1/profile/zikwall from platform: web-next@0.0.1-build-commit#hash
Request /api/v1/profile/zikwall from platform: web-next@0.0.1-build-commit#hash
Request /api/v1/contents/user/2 from platform: web-next@0.0.1-build-commit#hash
Request /api/v1/content/10 from platform: web-next@0.0.1-build-commit#hash
```

- [ ] Elasticsearch, Logstash, Kibana (ELK stack)
