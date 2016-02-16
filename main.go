package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

)


//port definition, pass as params
var port string = os.Getenv("PORT")

func main() {
	if port == "" {
		log.Println("No Port was provided, using 3000")
		port = "3000"
	}

	http.HandleFunc("/", root)
	log.Println("Listning for connections on port %s", port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Yes, it's running")
}

