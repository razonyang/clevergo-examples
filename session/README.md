# Session Example

```
$ go run session/main.go

$ curl -i --cookie-jar cj --cookie cj http://localhost:8080/
You visited this page 0 times

$ curl -i --cookie-jar cj --cookie cj http://localhost:8080/
You visited this page 1 times

$ curl -i --cookie-jar cj2 --cookie cj2 http://localhost:8080/
You visited this page 0 times

$ curl -i --cookie-jar cj2 --cookie cj2 http://localhost:8080/
You visited this page 1 times

...
```