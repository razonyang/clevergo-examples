# Graceful Shutdown Example

Starts a server:

```shell
$ go run gracefulshutdown/main.go
2020/06/20 03:16:50 Listening on :8080.
```

Make a request:

```shell
curl "http://localhost:8080/sleep?duration=10"
```

And then terminate server by `CTRL` + `C` sends a signal to process:

```shell
go run gracefulshutdown/main.go 
2020/06/20 03:17:15 | 200 | 10.001813588s | GET /sleep?duration=10 HTTP/1.1
^C2020/06/20 03:17:40 Shutting down server...
```

Finally, `curl`'s output is similar to the following message:

```
curl: (52) Empty reply from server
```
