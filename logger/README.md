# Logger Example

## Zap

```
$ go run logger/main.go -logger=zap

$ curl http://localhost:8080/log    
2020-06-20T02:35:44.194+0800    DEBUG   logger/main.go:36       debug msg
2020-06-20T02:35:44.194+0800    DEBUG   logger/main.go:37       debugf msg
2020-06-20T02:35:44.194+0800    INFO    logger/main.go:38       debug
2020-06-20T02:35:44.194+0800    INFO    logger/main.go:39       debugf
2020-06-20T02:35:44.194+0800    WARN    logger/main.go:40       warn msg
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