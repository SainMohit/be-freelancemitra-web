package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/be-freelancemitra-web/router"
	"github.com/be-freelancemitra-web/src/api"
)

func main() {
	apiObj := api.NewApiV1Obj()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8001),
		Handler:        router.InitializeRouter(apiObj),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
