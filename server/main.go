package main

import (
	"jojogo/server/config"
	"jojogo/server/infra"
)

// go run main.go
// curl localhost:8080/group/insert
// curl localhost:8080/group/swimming
// curl localhost:8080/group/create --request "POST"
// curl localhost:8080/group/update/state/gay/false --request "PATCH"

func main() {
	config.Init()
	infra.InitRouter()
	infra.Router.Run()
}
