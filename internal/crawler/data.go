package crawler

import (
	"encoding/json"
	"strings"
	"github.com/PuerkitoBio/goquery"
	"log"
)

type Data struct {

	Title string
	URL   string
	Texts []string
}




func Marshall(data []Data) string {

	b, err := json.Marshal(data)

	if err != nil {
		log.Printf("Unable to convert %v to json", data)
		return ""
	}

	return string(b)
}


func Unmarshall(s string, datas *[]Data) error {
	err := json.Unmarshal([]byte(s), datas)
	if err != nil {
		return err
	}
	return nil
}



func ParseText(dom *goquery.Selection) string {

	//remove script, style
	dom.Find("script, style").Remove()
	return dom.Text()
}

func ParseTexts(dom *goquery.Selection) []string {

	var texts []string

	for _, text := range  strings.Split(ParseText(dom), "\n") {
		t := strings.TrimSpace(text)
		if t != "" {
			texts = append(texts, t)
		}
	}
	return texts

}

