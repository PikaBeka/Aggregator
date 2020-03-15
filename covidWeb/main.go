package main

import(
	"fmt"
	"os"
	"net/http"
	"log"
)

func main(){
	port := os.Getenv("PORT")
	if port == "" {
		port = "7777"
	}
	http.HandleFunc("/", MainPage)
	
}