package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"encoding/json"
)


//port definition, pass as params

var port string = os.Getenv("PORT")

func main() {
	if port == ""{
		log.Println("No Port was provided, using 3000")
	}
	fmt.Println("Hello world")
}