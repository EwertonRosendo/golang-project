package main

import (
	router "api/src"
	"api/src/config"
	"fmt"
	"log"
	"net/http"
)

func main(){
	config.Load()
	r := router.Generate()

	fmt.Printf("listening on port: 5000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}