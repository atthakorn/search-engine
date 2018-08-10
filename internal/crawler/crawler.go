package crawler

import (
	"github.com/gocolly/colly"
	"fmt"
	"net/http"
	"sync"
	"net/url"
	"strings"
	"github.com/spf13/viper"
)

type Crawler struct {
	collector *colly.Collector
	sites []string
	links map[string]string
	mux *sync.Mutex
}



func (c *Crawler) isFile(link string) bool {

	exts := []string {"pdf", "docx", "jpg", "png"}

	for _, ext := range exts {
		if strings.HasSuffix(strings.ToLower(link), ext) {
			return true
		}
	}
	return false
}

func (c *Crawler) Start() {

	for _, site := range c.sites {
		c.collector.Visit(fmt.Sprintf("http://%s",site ))
	}

	c.collector.Wait()

}

func (c *Crawler) bootstrapCallback() {

	// On every a element which has href attribute call callback
	c.collector.OnHTML("a[href], area[href]", func(e *colly.HTMLElement) {

		c.mux.Lock()

		link := e.Attr("href")

		//validate url
		_, err := url.ParseRequestURI(link)

		if err == nil && !c.isFile(link) {
			_, found := c.links[link]
			if !found {
				c.links[link] = link
				// Visit link found on page, only those links are visited which are in AllowedDomains
				e.Request.Visit(e.Request.AbsoluteURL(link))
			}
		}

		c.mux.Unlock()
	})

	// Before making a request print "Visiting ..."
	c.collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("Fetching: %s\n", r.URL.String())
	})

}



func Make() *Crawler {

	sites := viper.GetStringSlice("sites")

	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains:
		colly.AllowedDomains(sites...),
		colly.Async(true),
		colly.MaxDepth(10),

	)
	//disable keep alive
	c.WithTransport(&http.Transport{
		DisableKeepAlives: true,
	})

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 1, RandomDelay: 1})

	crawler := &Crawler{c, sites, make(map[string]string),  &sync.Mutex{}}
	crawler.bootstrapCallback()
	return crawler;
}

