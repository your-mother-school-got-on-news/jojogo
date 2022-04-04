package main

import (
	"jojogo/server/config"
	"jojogo/server/infra"
	"net/http"
	"time"
)

func load() {
	config.Init()
}

func main() {
	load()
	infra.InitRouter()
	s := &http.Server{
		Addr:           ":8888",
		Handler:        infra.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()
}
