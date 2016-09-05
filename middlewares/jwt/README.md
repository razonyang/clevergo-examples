# JSON WEB TOKEN Middleware
This middleware requires [CleverGo JWT](https://github.com/clevergo/jwt).

## Install dependences
```
go get github.com/clevergo/jwt
```
About the usages of **jwt** package, please refer to https://github.com/clevergo/jwt.

## Example
[JSON WEB TOKEN Example](/examples/jwt/main.go)
```
go run $GOPATH/src/github.com/headwindfly/clevergo/examples/jwt/main.go
```

## Usage
Please refer to [JSON WEB TOKEN Example](/examples/jwt/main.go).

**Set Error Handler**
```
JWTMiddleware.SetErrorHandler(handler clevergo.HandlerFunc)
```

**Set Success Handler**
```
JWTMiddleware.SetSuccessHandler(handler func(ctx *clevergo.Context, token *jwt.Token)）
```
SuccessHandler is a useful handler, you can use it store the current user info in `Context`.