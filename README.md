# miiko-api

- how to build proto

```shell
cd proto
buf lint
npx buf generate miiko
```

- curl

```shell
curl -XPOST -H 'Content-Type: application/json' -d '{}' localhost:8081/miiko.v1.MiikoService/CategoryGet -i
curl -XPOST -H 'Content-Type: application/json' -d '{}' localhost:8081/miiko.v1.HealthService/Check -i
```