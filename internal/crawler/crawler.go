package crawler

import (
	"github.com/gocolly/colly"
	"fmt"
	"net/http"
	"sync"
	"github.com/spf13/viper"
	"time"
	"net/url"
	"log"
	"regexp"
)


//load config
func init() {


	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// if in project root
	viper.AddConfigPath(".")
	// if you are in cmd/crawler or internal/crawler
	viper.AddConfigPath("../..")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	log.Printf("Fetch Sites: %v\n", viper.GetStringSlice("sites"))


}


// Crawler
type Crawler struct {
	sites        []string
	maxDepth     int
	parallelism  int
	delay        int
	collector    *colly.Collector
	mux          *sync.Mutex
	total        int
}



// Check if  link is file
func (c *Crawler) isFile(link string) bool {

	match, err :=  regexp.MatchString(`\.\w+($|\?)`, link)
	if err != nil {
		return true  //if error, assume it is file
	}
	return match
}


// Start scraping
func (c *Crawler) Start() {

	start := time.Now()

	for _, site := range c.sites {
		c.collector.Visit(fmt.Sprintf("http://%s", site))
	}

	c.collector.Wait()

	elapsed := time.Since(start)

	log.Printf("Fetching Complete: %d pages in %s\n", c.total, elapsed)

}




func (c *Crawler) onResponse() colly.ResponseCallback {

	return  func(r *colly.Response) {

		r.Ctx.Put("time", time.Now())
	}

}


func (c *Crawler) onHtml() colly.HTMLCallback {

	return func(e *colly.HTMLElement) {

		link := e.Attr("href")
		//validate url
		_, err := url.ParseRequestURI(link)

		if err == nil &&  !c.isFile(link) {

			e.Request.Visit(e.Request.AbsoluteURL(link))
		}

	}
}


func (c *Crawler) onScraped() colly.ScrapedCallback {

	return func(r *colly.Response) {

		c.total++
		elapsed := time.Since(r.Ctx.GetAny("time").(time.Time))

		log.Printf("Scraped: %s (%s)\n", r.Request.URL, elapsed)
	}
}

func (c *Crawler) onError() colly.ErrorCallback {

	return func(r *colly.Response, e error) {

		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", e)

	}
}

func Make() *Crawler {


	sites := viper.GetStringSlice("sites")
	maxDepth := viper.GetInt("maxDepth")
	parallelism := viper.GetInt("parallelism")
	delay := viper.GetInt("delay")

	crawler := &Crawler{
		sites:       sites,
		maxDepth:    maxDepth,
		parallelism: parallelism,
		delay:       delay,
	}
	crawler.init()
	return crawler
}



// Crawer initializer
func (c *Crawler) init() {

	//mutex lock
	c.mux = &sync.Mutex{}


	//total page scraped
	c.total = 0

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
	collector.OnResponse(c.onResponse())

	// Called after page scraped
	collector.OnScraped(c.onScraped())

	// On every a element which has href attribute call callback
	collector.OnHTML("a[href], area[href]", c.onHtml())

	// Called if error occured during the request
	collector.OnError(c.onError())

	c.collector = collector

}