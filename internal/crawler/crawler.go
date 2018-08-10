package crawler

import (
	"github.com/gocolly/colly"
	"fmt"
	"net/http"
	"sync"
	"strings"
	"github.com/spf13/viper"
	"time"
	"net/url"
)

// Crawler
type Crawler struct {
	collector    *colly.Collector
	sites        []string
	crawledLinks map[string]string
	mux          *sync.Mutex
}



// Check if  link is file
func (c *Crawler) isFile(link string) bool {

	exts := []string {"pdf", "docx", "jpg", "png", "txt", "xslx", "gif"}

	for _, ext := range exts {
		if strings.HasSuffix(strings.ToLower(link), ext) {
			return true
		}
	}
	return false
}

func (c *Crawler) Start() {

	start := time.Now()

	for _, site := range c.sites {
		c.collector.Visit(fmt.Sprintf("http://%s",site ))
	}

	c.collector.Wait()

	elapsed := time.Since(start)

	fmt.Println("Fetching Complete: %s", elapsed)

}


func (c *Crawler) bootstrapCallback() {


	// Called after response received
	c.collector.OnScraped(func(r *colly.Response) {
		fmt.Printf("Scraped: %s\n", r.Request.URL)
	})


	// On every a element which has href attribute call callback
	c.collector.OnHTML("a[href], area[href]", func(e *colly.HTMLElement) {

		c.mux.Lock()
		defer c.mux.Unlock()

		link := e.Attr("href")

		//validate url
		_, err := url.ParseRequestURI(link)
		if err == nil && !c.isFile(link) {
			_, ok  := c.crawledLinks[link]
			if !ok  {
				c.crawledLinks[link] = link
				// Visit link found on page, only those crawledLinks are visited which are in AllowedDomains
				e.Request.Visit(e.Request.AbsoluteURL(link))
			}
		}

	})



	// Called if error occured during the request
	c.collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

}



func Make() *Crawler {

	sites := viper.GetStringSlice("sites")
	maxDepth := viper.GetInt("maxDepth")
	parallelism := viper.GetInt("parallelism")

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains:
		colly.AllowedDomains(sites...),
		colly.Async(true),
		colly.MaxDepth(maxDepth),

	)
	//disable keep alive
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism:  parallelism})

	crawler := &Crawler{c, sites, make(map[string]string),  &sync.Mutex{}}
	crawler.bootstrapCallback()
	return crawler;
}

