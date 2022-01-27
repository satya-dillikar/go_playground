#REST API Server 

https://golang.org/doc/tutorial/web-service-gin

mkdir web-service-gin
cd web-service-gin
go mod init satya.com/web-service-gin
go get .

#start server
go run .


#client - use in another terminal
#GET method
 curl http://localhost:8080/albums

 curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"

 curl http://localhost:8080/albums/2

#POST method
 curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'