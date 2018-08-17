package query

import (
	"github.com/blevesearch/bleve"
	"log"
	"github.com/atthakorn/web-scraper/internal/config"

	_ "github.com/atthakorn/web-scraper/internal/blevex/lang/th"
	"encoding/json"
	"github.com/blevesearch/bleve/search"
	"time"
)


type Hit struct {
	Title string
	URL string
	Description string
	Score float64
}


type Result struct {
	TotalHit uint64
	//time in ms
	Time float64
	Hits []Hit
}



func Search(keyword string) *Result {

	index, err := openIndex()
	if err != nil {
		return nil
	}
	defer index.Close()

	//use fuzzy query with highlight
	query := bleve.NewFuzzyQuery(keyword)
	request := bleve.NewSearchRequest(query)
	request.Highlight = bleve.NewHighlight()
	request.Fields = []string{"Title", "URL", "Body"}

	searchResults, err := index.Search(request)
	if err != nil {
		log.Printf("Search error: %v", err)
		return nil
	}


	result :=  newResult(searchResults)

	jsonByte, err := json.Marshal(result)


	log.Printf("result: %v", string(jsonByte))

	return result
}




func openIndex() (bleve.Index,  error)  {


	index, err := bleve.Open(config.IndexPath)
	if err != nil  {
		log.Printf("Cannot open index at %s with error %v", config.IndexPath, err)
		return nil, err
	}

	log.Printf("Opening index...")
	return index, nil
}



func newResult(searchResult *bleve.SearchResult) *Result {

	result := &Result{
		TotalHit: searchResult.Total,
		Time: float64(searchResult.Took) / float64(time.Millisecond),
	}
	for _, searchResultHit := range searchResult.Hits {
		result.Hits = append(result.Hits, *newHit(searchResultHit))
	}

	return result
}

func newHit(searchResultHit *search.DocumentMatch) *Hit {

	hit := &Hit {
		URL: searchResultHit.ID,
		Score: searchResultHit.Score,
	}

	fields := searchResultHit.Fields
	hit.Title = fields["Title"].(string)

	//if fragment version existed, replace with fragment value
	fragments := searchResultHit.Fragments
	if fragments["Tittle"] != nil {
		hit.Title = fragments["Tittle"][0]
	}
	if fragments["Body"] != nil {
		hit.Description = fragments["Body"][0]
	}
	if fragments["URL"] != nil {
		hit.URL = fragments["URL"][0]
	}

	return hit
}