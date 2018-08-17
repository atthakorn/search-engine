package search

import (
	"github.com/blevesearch/bleve"
	"log"
	"github.com/atthakorn/web-scraper/internal/config"

	_ "github.com/atthakorn/web-scraper/internal/blevex/lang/th"
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
	TotalHit int
	//time in ms
	Time float64
	Hits []Hit
}





func Query(keyword string) *Result {

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

	return newResult(searchResults)
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
		Time: float64(searchResult.Took) / float64(time.Millisecond),
	}
	for _, searchResultHit := range searchResult.Hits {
		result.Hits = append(result.Hits, *newHit(searchResultHit))
		result.TotalHit = len(result.Hits)
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