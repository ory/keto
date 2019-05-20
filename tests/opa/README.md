
```shell
$ docker run -p 8181:8181 openpolicyagent/opa \
    run --server
```


```shell
./policies.sh

curl -X POST --data-binary @input-allow.json -H 'Content-Type: application/json' localhost:8181/v1/data/ory/exact/allow
```
