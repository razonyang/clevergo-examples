# Authentication Example

```shell
$ go run auth/main.go

$ curl http://localhost:8080/auth
unauthorized

$ curl "http://localhost:8080/auth?access_token=footoken"
hello foo

$ curl -H "Authorization:Bearer bartoken" http://localhost:8080/auth
hello bar

$ curl -H "Authorization:Bearer invalidtoken" http://localhost:8080/auth
unauthorized
```
