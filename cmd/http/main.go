package main

import (
	"net/http"
	"github.com/atthakorn/search-engine/internal/web"
	"github.com/atthakorn/search-engine/internal/config"
	"fmt"
)



func main() {



	//map to handler
	http.HandleFunc("/", web.Handler)


	address := config.HttpAddress

	fmt.Printf("\nsearch application ready to serves at http://localhost%s", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Printf("Cannot start server at: %s", address)
	}




}