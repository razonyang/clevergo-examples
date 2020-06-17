# I18N Example

```shell
$ cd i18n && go run main.go
```

```shell
## fallback language(default to English)
$ curl "http://localhost:8080"
Home

## retrieve prefered language from URL query
$ curl "http://localhost:8080?lang=zh"
主页

$ curl "http://localhost:8080?lang=zh-TW"
主頁

$ curl "http://localhost:8080?lang=zh-HK"
主頁

## retrieve prefered language Cookie
$ curl -b "lang=zh-Hant" "http://localhost:8080"
主頁

## retrieve prefered language from header
$ curl -H "Accept-Language: zh-CN,zh;q=0.9,en;q=0.8,en-US;q=0.7,zh-TW;q=0.6,pt;q=0.5" "http://localhost:8080/hello?name=张三"
你好，张三
```
