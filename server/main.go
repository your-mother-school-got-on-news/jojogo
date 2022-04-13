package main

import (
	"jojogo/server/infra"
)

// go run main.go
// curl localhost:8080/books
// curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
// curl localhost:8080/checkout?id=2 --request "PATCH"
// curl localhost:8080/return?id=2 --request "PATCH"

func main() {
	infra.InitRouter()
	infra.Router.Run()
}
