package crawler

import (
	"fmt"
	"os"
	"github.com/gocolly/colly"
	"strings"
	"io/ioutil"
)

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


func LoadString(filename string) (string, error) {

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(f), err
}


func WriteString(filename string, content string) {


	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
}

func ParseHtml(html *colly.HTMLElement) string {

	//remove script, style
	html.DOM.Find("script").Remove()
	html.DOM.Find("style").Remove()

	return html.DOM.Text()
}

func ParseHtmlArray(html *colly.HTMLElement) []string {

	var phrases []string

	for _, pharse := range  strings.Split(ParseHtml(html), "\n") {
		t := strings.TrimSpace(pharse)
		 if t != "" {
			 phrases = append(phrases, t)
		 }
	}
	return phrases

}

