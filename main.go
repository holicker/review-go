package main

import (
	"log"
	"review-go/rest"
)

func main() {
	log.Println("========== Main Log : API Run Start")
	rest.RunAPI("127.0.0.1:9090")
}
