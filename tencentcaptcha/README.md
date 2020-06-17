# Tencent Captcha Example 腾讯验证码例子

> 修改其中的 `secretID`、`secretKey`、`appID`、`AppSecret`。

```shell
$ SECRET_ID={SECRET_ID} \
    SECRET_KEY={SECRET_KEY} \
    APP_ID={APP_ID} \
    APP_SECRET_KEY={APP_SECRET_KEY} \
    go run tencentcaptcha/main.go
```

- `SECRET_ID`：安全凭证 ID
- `SECRET_KEY`：安全凭证 Key
- `APP_ID`：验证码应用 ID
- `APP_SECRET_KEY`：验证码应用 Secret Key

官方文档：https://cloud.tencent.com/document/product/1110