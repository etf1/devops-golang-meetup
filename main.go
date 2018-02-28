package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World\n")
}

func main() {
	fmt.Println("Startin app")
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8080", nil)
}
