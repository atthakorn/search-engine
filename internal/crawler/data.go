package crawler

import (
	"encoding/json"
	"log"
	"github.com/gocolly/colly"
	"os"
	"fmt"
)

type Data struct {

	Title string
	URL string
	Parse []string
}



func (d *Data) ToJson() string {

	b, err := json.Marshal(d)

	if err != nil {
		log.Printf("Unable to convert %v to json", d)
		return ""
	}
	return string(b)
}


func save(filename string, d *Data) {

	filename = fmt.Sprintf("./data/%s.json", filename)

	f, err := os.Create(filename)
	if err != nil {
		//os.Create("./data/crawl.json")
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(d.ToJson()); err != nil {
		panic(err)
	}
}

func writeLine( s string) {


	f, err := os.Create("./data/test.txt")
	if err != nil {
		//os.Create("./data/crawl.json")
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(s); err != nil {
		panic(err)
	}
}


func html2text(html *colly.HTMLElement) string {

	//remove script
	html.DOM.Find("script").Remove()

	//remove style
	html.DOM.Find("style").Remove()

//	text := sanitizer.Sanitize(html.DOM.Text())


	return html.DOM.Text()

}