# CleverGo Log Example

```shell
$ go run log/main.go -adapter zap   
{"level":"debug","msg":"Debug"}
{"level":"debug","msg":"Debugf"}
{"level":"info","msg":"Info"}
{"level":"info","msg":"Infof"}
{"level":"warn","msg":"Warn"}
{"level":"warn","msg":"Warnf"}
{"level":"error","msg":"Error"}
{"level":"error","msg":"Errorf"}

$ go run log/main.go -adapter logrus 
INFO[0000] Info                                         
INFO[0000] Infof                                        
WARN[0000] Warn                                         
WARN[0000] Warnf                                        
ERRO[0000] Error                                        
ERRO[0000] Errorf   
```
