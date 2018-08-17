package main

import (
	"os"
	"log"
	"github.com/atthakorn/web-scraper/internal/search"
	"encoding/json"
)

func main() {


	if  len(os.Args) < 2 {
		log.Printf("No keyword argument supplied. Usage: search [keyword]")
		return
	}

	arg := os.Args[1]

	result:=	search.Query(arg)
	jsonByte, _ := json.Marshal(result)

	log.Printf("result: %v", string(jsonByte))

}

