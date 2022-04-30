package main

import (
	"jojogo/server/config"
	"jojogo/server/infra"
)

func main() {
	config.Init()
	infra.InitRouter()
	infra.Router.Run(":8080")
}
