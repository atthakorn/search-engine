package crawler

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync/atomic"
	"time"

	"github.com/gocolly/colly"
	"github.com/spf13/viper"
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
	sites       []string
	maxDepth    int
	parallelism int
	delay       int
	collector   *colly.Collector
	total       int64
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

// Check if  link is file
func (c *Crawler) isBlacklist(link string) bool {

	// whitelist
	for _, ext := range []string{".php", ".jsp", ".asp", ".aspx", "html", "htm"} {

		if strings.HasSuffix(strings.ToLower(link), ext) {
			return false
		}
	}

	//regex test whether url is end with any file extensions
	match, err := regexp.MatchString(`\.\w+($|\?)`, link)
	if err != nil {
		return true //if error, assume it is file
	}

	return match
}

// Start scraping
func (c *Crawler) Start() {

	start := time.Now()

	//reset count to zero
	c.total = 0
	//start crawl at entry point
	for _, site := range c.sites {
		c.collector.Visit(fmt.Sprintf("http://%s", site))
	}

	//wait until workers all done
	c.collector.Wait()

	elapsed := time.Since(start)

	log.Printf("Fetching Complete: %d pages in %s\n", c.total, elapsed)

}

func (c *Crawler) onResponse() colly.ResponseCallback {

	return func(r *colly.Response) {

		r.Ctx.Put("time", time.Now())
	}

}

func (c *Crawler) onHtml() colly.HTMLCallback {

	return func(e *colly.HTMLElement) {

		link := e.Request.AbsoluteURL(e.Attr("href"))

		//validate url
		_, err := url.ParseRequestURI(link)

		if err == nil && !c.isBlacklist(link) {

			e.Request.Visit(link)
		}
	}
}

func (c *Crawler) onScraped() colly.ScrapedCallback {

	return func(r *colly.Response) {

		atomic.AddInt64(&c.total, 1)

		elapsed := time.Since(r.Ctx.GetAny("time").(time.Time))

		log.Printf("Scraped: %s (%s)\n", r.Request.URL, elapsed)
	}
}

func (c *Crawler) onError() colly.ErrorCallback {

	return func(r *colly.Response, e error) {

		log.Printf("Error: %s (%s)\n", r.Request.URL, e)

	}
}
