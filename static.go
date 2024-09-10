package main

import (
	
	//"encoding/base64"
	"fmt"
	
	"net/http"
	
	//"crypto/rand"
)


func main() {
	fmt.Println("serving files on port 3000")
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}


