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

	defer index.Close()

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

	//bulk by indexing 100 documents at a time
	batch := index.NewBatch()
	batchSize := 100
	batchCount:= 0
	for _, entry := range entries {

		//skip entry if it is directory
		if entry.IsDir() {
			continue
		}
		file := filepath.Join(dataPath, entry.Name())
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
			return 0, err
		}

		for _, data := range datas {
			count++
			batch.Index(data.URL, &Data{
				URL:   data.URL,
				Title: data.Title,
				Body:  strings.Join(data.Texts, " · "),
			})

			batchCount++
			if batchCount >= batchSize {
				err = index.Batch(batch)
				if err != nil {
					log.Printf("Bulk indexing error: %v", err)
					return 0, err
				}
				batch = index.NewBatch()
				batchCount = 0
			}

		}
	}

	return count, nil
}
