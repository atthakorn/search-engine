package main

import (
	"github.com/atthakorn/search-engine/internal/crawler"
)

func main()  {


	crawler.LoadConfig()



	crawler.Make().Start()




}


