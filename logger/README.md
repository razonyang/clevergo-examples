# Logger Example

## Zap

```
$ go run logger/main.go -logger=zap

$ curl http://localhost:8080/log
{"level":"info","msg":"Listening on :8080.\n"}
{"level":"debug","msg":"debug msg"}
{"level":"debug","msg":"debugf msg"}
{"level":"info","msg":"debug"}
{"level":"info","msg":"debugf"}
{"level":"warn","msg":"warn msg"}
{"level":"warn","msg":"warnf msg"}
{"level":"error","msg":"error msg"}
{"level":"error","msg":"errorf msg"}
...
```

## Logrus

```
$ go run logger/main.go -logger=logrus

$ curl http://localhost:8080/log
INFO[0001] debug                                        
INFO[0001] debugf                                       
WARN[0001] warn msg                                     
WARN[0001] warnf msg                                    
ERRO[0001] error msg                                    
ERRO[0001] errorf msg 
...
```