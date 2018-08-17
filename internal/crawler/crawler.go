package crawler

import (
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"sync/atomic"
	"time"
	"sync"
	"github.com/gocolly/colly"
	"github.com/atthakorn/search-engine/internal/config"
	"os"
)

// Crawler
type Crawler struct {
	entryPoints []string
	maxDepth    int
	parallelism int
	delay       int
	collector   *colly.Collector
	total       int64
	mutexLock	sync.Mutex
}

func New() *Crawler {

	crawler := &Crawler{
		entryPoints: config.EntryPoints,
		maxDepth:    config.MaxDepth,
		parallelism: config.Parallelism,
		delay:       config.Delay,
	}
	crawler.init()
	return crawler
}

// Crawer initializer
func (c *Crawler) init() {

	var domains []string

	for _, entryPoint := range c.entryPoints {
		u, err := url.Parse(entryPoint)
		if err != nil {
			log.Printf("Fail to load entry points")
		}

		domains = append(domains, u.Hostname())

	}

	// Instantiate default collector
	collector := colly.NewCollector(
		// Visit only domains:
		colly.AllowedDomains(domains...),
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

	// Find next link
	collector.OnHTML("a[href], area[href]", c.onNext())

	// Scraping html
	collector.OnHTML("html", c.onScraping())

	// Called after page scraped
	collector.OnScraped(c.onScraped())

	// Called if error occured during the request
	collector.OnError(c.onError())

	c.collector = collector


	c.setupDataDirectory()

}

func (c *Crawler) setupDataDirectory() {

	//destroy any outdated data
	os.RemoveAll(config.DataPath)

	//create data placeholder
	os.Mkdir(config.DataPath, 0755)


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



	log.Printf("Fetch Sites: %v\n", c.collector.AllowedDomains)

	start := time.Now()


	//reset count to zero
	c.total = 0
	//start crawl at entry point
	for _, entryPoint := range c.entryPoints {
		c.collector.Visit(entryPoint)
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

func (c *Crawler) onScraping() colly.HTMLCallback {

	return func(e *colly.HTMLElement) {

		//parse text only if response is OK
		if e.Response.StatusCode == http.StatusOK {
			//construct data
			data := &Data{
				Title: e.DOM.Find("Title").Text(),
				URL:   e.Request.URL.String(),
				Texts: ParseTexts(e.DOM),
			}

			//mutual lock for mutithread

			c.mutexLock.Lock()
			defer c.mutexLock.Unlock()

			var datas []Data

			file := JsonFileByHostname(e.Request.URL.Hostname())
			jsonString, _ := LoadString(file)


			Unmarshal(jsonString, &datas)
			datas = append(datas, *data)

			WriteString(file, Marshal(datas))


		} else {
			log.Printf("Fail to parse: %s", e.Request.URL)
		}

	}
}

func (c *Crawler) onNext() colly.HTMLCallback {

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
