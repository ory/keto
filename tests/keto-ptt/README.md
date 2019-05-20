# Performance Test Tool (PTT)

```
$ go install github.com/ory/keto/tests/keto-ptt
$ export GO111MODULE=on
```

```
# macos: 
$ curl -L -o opa https://github.com/open-policy-agent/opa/releases/download/v0.10.7/opa_darwin_amd64
```

## Keto

To run the PTT, ORY Keto must be running at `http://localhost:4466`.

```shell
$ export DSN=memory
$ export GO111MODULE=on
$ go install github.com/ory/keto
# $ PROFILING=memory keto serve
$ PROFILING=cpu keto serve
```

```shell
$ keto-ptt run regex
```

## OPA


```
$ opa run --server --log-level error
$ ./opa-policies.sh
```

```
$ keto-ptt opa regex
```