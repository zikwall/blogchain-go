[![build](https://github.com/zikwall/blogchain-go/workflows/Build%20and%20test%20Blogchain/badge.svg)](https://github.com/zikwall/blogchain-go/actions)

<div align="center">
  <img width="150" height="150" src="https://github.com/zikwall/blogchain/blob/master/screenshots/bc_go_300.png">
  <h1>Blog Chain</h1>
  <h5>App server written in Golang</h5>
</div>

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
