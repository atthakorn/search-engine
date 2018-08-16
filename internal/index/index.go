package index

import (
	"github.com/atthakorn/web-scraper/internal/config"
	"github.com/blevesearch/bleve"
	"log"
	"github.com/blevesearch/bleve/mapping"
	"github.com/atthakorn/web-scraper/internal/blevex/lang/th"
	"io/ioutil"
	"path/filepath"
	"github.com/atthakorn/web-scraper/internal/crawler"

	"strings"
	"time"
)



func Index() {

	index := openOrCreate()
	benchmark := benchmark(index, indexing)
	benchmark()
}





func Query(index bleve.Index) {

	//query
	query := bleve.NewFuzzyQuery("Information")
	search := bleve.NewSearchRequest(query)
	searchResults, err := index.Search(search)
	if err != nil {
		log.Printf("Search error: %v", err)
		return
	}
	log.Printf("Result: %s", searchResults)

}

func openOrCreate() bleve.Index {

	indexPath := config.IndexPath

	index, err := bleve.Open(indexPath)

	if err == bleve.ErrorIndexPathDoesNotExist || err == bleve.ErrorIndexMetaMissing {
		log.Printf("Creating new index...")

		indexMapping := buildIndexMapping()
		index, err = bleve.New(indexPath, indexMapping)

		if err != nil {
			log.Printf("Terminate indexer, cannot create index at %s", indexPath)
			return nil
		}

	} else {
		log.Printf("Open existing index ...")
	}

	return index
}

func buildIndexMapping() (mapping.IndexMapping) {

	indexMapping := bleve.NewIndexMapping()
	indexMapping.DefaultAnalyzer = th.AnalyzerName

	return indexMapping
}



func benchmark(index bleve.Index, indexing func(index bleve.Index) (int, error)) func(){

	return func() {

		startTime := time.Now()

		count, err := indexing(index)
		if err != nil {
			log.Printf("Indexing Error: %v", err)

		}
		indexDuration := time.Since(startTime)
		indexDurationSeconds := float64(indexDuration) / float64(time.Second)
		timePerDocument := float64(indexDuration) / float64(count)
		log.Printf("Indexed %d documents, in %.2fs (average %.2f ms/documents)", count, indexDurationSeconds, timePerDocument/float64(time.Millisecond))
	}
}

func indexing(index bleve.Index) (count int, err error) {

	//total number of indexed documents
	count = 0
	dataPath := config.DataPath
	entries, err := ioutil.ReadDir(dataPath)

	if err != nil {
		log.Printf("Terminate indexer, cannot load data entries at %s", dataPath)
		return 0, err
	}

	for _, entry := range entries {

		file := filepath.Join(config.DataPath, entry.Name())
		json, err := crawler.LoadString(file)

		if err != nil {
			log.Printf("Fail to load crawler data file: %s", file)
			return 0, err
		}

		//unmarshal json
		var datas []crawler.Data
		err = crawler.Unmarshal(json, &datas)

		if err != nil {
			log.Printf("Fail to unmarshal data from json: %s", file)
		}

		for _, data := range datas {
			count++
			index.Index(data.URL, &Data{
				URL:   data.URL,
				Title: data.Title,
				Body:  strings.Join(data.Texts, " · "),
			})
		}
	}

	return count, nil
}
