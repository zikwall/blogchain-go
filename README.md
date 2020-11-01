[![build](https://github.com/zikwall/blogchain-go/workflows/Build%20and%20test%20Blogchain/badge.svg)](https://github.com/zikwall/blogchain-go/actions)

<div align="center">
  <img width="150" height="150" src="https://github.com/zikwall/blogchain/blob/master/screenshots/bc_go_300.png">
  <h1>Blog Chain</h1>
  <h5>Powerful and productive server written in Go</h5>
</div>

### Related Projects

- [x] [Blogchain Client, powered by Next.js](https://github.com/zikwall/blogchain)

## Development

- Native
```shell script
go run . \
  --bind-address 0.0.0.0:3001 \
  --database-host @ \
  --database-user blogchain \
  --database-password 123456 \
  --database-name blogchain \
  --database-driver mysql \
  --container-secret secret
```
- Docker

```shell script
docker run -d --net=host \
   -e BIND_ADDRESS='0.0.0.0:3001' \
   -e DATABASE_HOST='<database host: @>' \
   -e DATABASE_USER='<database username>' \
   -e DATABASE_PASSWORD='<database password>' \
   -e DATABASE_NAME='<database name>' \
   -e DATABASE_DRIVER='<database username: mysql>' \
   -e CONTAINER_SECRET='<blogchain application secret>' \
   --name golang-blogchain-server qwx1337/blogchain-server:latest
```

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
[BLOGCHAIN INFO]: Blogchain Client request /api/editor/content/10 from platform web-next@0.0.1-build-commit#hash
[BLOGCHAIN INFO]: Blogchain Client request /api/v1/tags from platform web-next@0.0.1-build-commit#hash
[BLOGCHAIN INFO]: Blogchain Client request /api/editor/content/10 from platform web-next@0.0.1-build-commit#hash
[BLOGCHAIN INFO]: Blogchain Client request /api/v1/tags from platform web-next@0.0.1-build-commit#hash
```

- [ ] Elasticsearch, Logstash, Kibana (ELK stack)
