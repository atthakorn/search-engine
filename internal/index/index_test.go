package index

import (
	"github.com/atthakorn/web-scraper/internal/config"
	"github.com/blevesearch/bleve"
	"testing"
	"github.com/blevesearch/bleve/mapping"
	"github.com/stretchr/testify/assert"
	"os"
	"fmt"
)

func destroyDataFolder(index bleve.Index, dataPath string) {
	index.Close()
	config.DataPath = dataPath //remove data
	err := os.RemoveAll(config.IndexPath)
	if err != nil {
		fmt.Printf("cannot remove test data path: %v", err)
	}
}

func TestCreateIndexMapping(t *testing.T) {

	indexMapping := buildIndexMapping().(*mapping.IndexMappingImpl)

	assert.Equal(t, "th", indexMapping.DefaultAnalyzer)

	assert.IsType(t, &mapping.IndexMappingImpl{}, indexMapping)
}

func TestCreateIndex(t *testing.T) {

	config.DataPath = "../../testdata" //load data from testdata
	config.IndexPath = "./data/index"

	index := openOrCreate()

	defer destroyDataFolder(index, config.DataPath)

	_, ok := index.(bleve.Index)

	assert.True(t, ok)

	assert.Equal(t, "th", index.Mapping().(*mapping.IndexMappingImpl).DefaultAnalyzer)
}

func TestIndexing(t *testing.T) {

	config.DataPath = "../../testdata" //load data from testdata
	config.IndexPath = "./data/index"

	index := openOrCreate()

	defer destroyDataFolder(index, "./data")


	count, _ := indexing(index)

	assert.Equal(t, 15, count)
}



func TestIndexingFail(t *testing.T) {

	config.DataPath = "./anywhere"
	config.IndexPath = "./data/index"

	index := openOrCreate()
	defer destroyDataFolder(index, "./data")

	_, err := indexing(index)

	assert.NotNil(t, err)



}
