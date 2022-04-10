package main

import (
	api "server/api"
)

// go run main.go
// curl localhost:8080/books
// curl localhost:8080/books --include --header "Content-Type: application/json" -d @body.json --request "POST"
// curl localhost:8080/checkout?id=2 --request "PATCH"
// curl localhost:8080/return?id=2 --request "PATCH"

func main() {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)
	// fmt.Println(exPath)

	// os.Setenv("$GOPATH", exPath)

	api.Start()
}
