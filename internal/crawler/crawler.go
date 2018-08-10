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
	sites        []string
	maxDepth     int
	parallelism  int
	collector    *colly.Collector
	crawledLinks map[string]string
	mux          *sync.Mutex
}

// Check if  link is file
func (c *Crawler) isFile(link string) bool {

	exts := []string{"pdf", "docx", "jpg", "png", "txt", "xslx", "gif"}

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
		c.collector.Visit(fmt.Sprintf("http://%s", site))
	}

	c.collector.Wait()

	elapsed := time.Since(start)

	fmt.Println("Fetching Complete: %s", elapsed)

}


func (c *Crawler) init() {

	//mutex lock
	c.mux = &sync.Mutex{}
	//memory for keep tracking old link
	c.crawledLinks = make(map[string]string)

	// Instantiate default collector
	collector := colly.NewCollector(
		// Visit only domains:
		colly.AllowedDomains(c.sites...),
		colly.Async(true),
		colly.MaxDepth(c.maxDepth),

	)
	//disable keep alive
	collector.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: c.parallelism})

	// Called after response received
	collector.OnScraped(func(r *colly.Response) {
		fmt.Printf("Scraped: %s\n", r.Request.URL)
	})

	// On every a element which has href attribute call callback
	collector.OnHTML("a[href], area[href]", func(e *colly.HTMLElement) {

		c.mux.Lock()
		defer c.mux.Unlock()

		link := e.Attr("href")

		//validate url
		_, err := url.ParseRequestURI(link)
		if err == nil && !c.isFile(link) {
			_, ok := c.crawledLinks[link]
			if !ok {
				c.crawledLinks[link] = link
				// Visit link found on page, only those crawledLinks are visited which are in AllowedDomains
				e.Request.Visit(e.Request.AbsoluteURL(link))
			}
		}

	})

	// Called if error occured during the request
	collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})


	c.collector = collector

}

func Make() *Crawler {

	sites := viper.GetStringSlice("sites")
	maxDepth := viper.GetInt("maxDepth")
	parallelism := viper.GetInt("parallelism")

	crawler := &Crawler{
		sites:    sites,
		maxDepth: maxDepth,
		parallelism: parallelism,
	}
	crawler.init()
	return crawler
}
