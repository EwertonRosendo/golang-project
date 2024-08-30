package main

import (
	router "api/src"
	"api/src/config"
	//"encoding/base64"
	"fmt"
	"log"
	"net/http"
	//"crypto/rand"
)
/*
func init(){
	key := make([]byte, 64)

	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(key)
	fmt.Println(stringBase64)
}
*/

func main(){
	config.Load()
	r := router.Generate()

	fmt.Printf("listening on port: 5000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}