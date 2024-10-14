## как запустить?

для тестов:

```
$ go test -v task2/server/tests/server_test.go
```

для ручной проверки:
```
go run task2/server/runner/runner.go
```
```
curl http://localhost:8080/version
```
Вывод: v1.0.0
```
curl -X POST http://localhost:8080/decode -H "Content-Type: application/json" -d '{"inputString": "c29tZSBzdHJpbmc="}'
```
Вывод: {"outputString":"some string"}
```
curl http://localhost:8080/hard-op
```
Вывод 1: 200 OK
Вывод 2: 500 Internal Server Error