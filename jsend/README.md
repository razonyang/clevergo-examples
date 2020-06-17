
## JSend Example

```shell
$ go run jsend/main.go

# fetch users
$ curl http://localhost:8080/users
{"status":"success","data":[{"id":"foo","email":"foo@example.com"},{"id":"bar","email":"bar@example.com"}]}

# create user without required data
$ curl -XPOST  http://localhost:8080/users
{"status":"fail","data":{"email":["email can not be blank"],"id":["id can not be blank"]}}

# create test user
$ curl -XPOST -d "id=test&email=test@example.com" http://localhost:8080/users
{"status":"success","data":{"id":"test","email":"test@example.com"}}

# refetch users
$ curl http://localhost:8080/users
{"status":"success","data":[{"id":"foo","email":"foo@example.com"},{"id":"bar","email":"bar@example.com"},{"id":"test","email":"test@example.com"}]}

# fetch test user
$ curl http://localhost:8080/users/test
{"status":"success","data":{"id":"test","email":"test@example.com"}}

# delete test user
$ curl -XDELETE http://localhost:8080/users/test
{"status":"success","data":null}

# refetch test user
$ curl http://localhost:8080/users/test
{"status":"error","data":null,"message":"User Not Found"}
```