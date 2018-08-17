package main

import (
	"net/http"
	"github.com/labstack/gommon/log"
	"github.com/atthakorn/web-scraper/internal/web"
	"github.com/atthakorn/web-scraper/internal/config"
)



func main() {



	//map to handler
	http.HandleFunc("/", web.Handler)


	address := config.HttpAddress
	err := http.ListenAndServe(address, nil)

	if err != nil {
		log.Printf("Cannot start server at: %s", address)
	}



}