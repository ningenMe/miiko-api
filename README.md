# miiko-api

- how to build proto

```shell
cd proto
buf lint
npx buf generate miiko
```

- how to set env-variable

```shell
`aws ssm get-parameters-by-path --path "/" --region ap-northeast-1 --output text | awk '{print "export",$5"="$7}'`
```

- curl

```shell
curl -XPOST -H 'Content-Type: application/json' -d '{}' localhost:8081/miiko.v1.HealthService/Check -i
```

```shell
curl -XPOST -H 'Content-Type: application/json' -d '{}
' localhost:8081/miiko.v1.MiikoService/CategoryGet -i
curl -XPOST -H 'Content-Type: application/json' -d '
    {
      "category" : {
        "categoryDisplayName": "テスト",
        "categorySystemName": "test",
        "categoryOrder": 1
      }
    }
' localhost:8081/miiko.v1.MiikoService/CategoryPost -i
curl -XPOST -H 'Content-Type: application/json' -d '
    {
      "categoryId" : "category_BKL1Q5",
      "category" : {
        "categoryDisplayName": "テスト改",
        "categorySystemName": "test2",
        "categoryOrder": -1
      }
    }
' localhost:8081/miiko.v1.MiikoService/CategoryPost -i
curl -XPOST -H 'Content-Type: application/json' -d '
    {
      "categoryId" : "category_BKL1Q5"
    }
' localhost:8081/miiko.v1.MiikoService/CategoryPost -i
```

```shell
curl -XPOST -H 'Content-Type: application/json' -d '
    {
      "categoryId" : "category_1J6WNP"
    }
' localhost:8081/miiko.v1.MiikoService/TopicGet -i
```

