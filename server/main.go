package main

import (
	"jojogo/server/infra"
	"net/http"
	"time"
)

func main() {

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
