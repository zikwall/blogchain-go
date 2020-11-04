[![build](https://github.com/zikwall/blogchain-go/workflows/Build%20and%20test%20Blogchain/badge.svg)](https://github.com/zikwall/blogchain-go/actions)

<div align="center">
  <img width="150" height="150" src="https://github.com/zikwall/blogchain/blob/master/screenshots/bc_go_300.png">
  <h1>Blog Chain</h1>
  <h5>Simple, Powerful and productive server written in Go</h5>
</div>

### Related Projects

- [x] [Blogchain Client, powered by Next.js](https://github.com/zikwall/blogchain)
- [x] [Docker Compose (Blogchain Compose)](https://github.com/zikwall/blogchain-compose)

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
