https://thenewstack.io/make-a-restful-json-api-go/

REST SERVER

cd github/projects/go_playground/tns-restful-json-api/v9
go mod init satya.com/tns-restful-json-api
go mod tidy

START SERVER
go run *.go


curl -XGET http://127.0.0.1:8080

curl -XGET http://127.0.0.1:8080/todos

curl -XGET http://127.0.0.1:8080/todos/1

curl -XPOST http://127.0.0.1:8080/todos  -d'{ "name": "Restful", "completed" :false}'

 curl -XDELETE http://127.0.0.1:8080/todos/15