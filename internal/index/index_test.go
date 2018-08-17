package index

import (
	"github.com/atthakorn/web-scraper/internal/config"
	"github.com/blevesearch/bleve"
	"testing"
	"github.com/blevesearch/bleve/mapping"
	"github.com/stretchr/testify/assert"
	"os"
)

func closeAndDestroy(index bleve.Index, indexPath string) {
	index.Close()
	os.RemoveAll(indexPath)
}

func TestCreateIndexMapping(t *testing.T) {

	indexMapping := buildIndexMapping().(*mapping.IndexMappingImpl)

	assert.Equal(t, "th", indexMapping.DefaultAnalyzer)

	assert.IsType(t, &mapping.IndexMappingImpl{}, indexMapping)
}

func TestCreateIndex(t *testing.T) {

	config.DataPath = "../../testdata" //load data from testdata
	config.IndexPath = "./index"

	index := newIndex()

	defer closeAndDestroy(index, config.IndexPath)

	_, ok := index.(bleve.Index)

	assert.True(t, ok)

	assert.Equal(t, "th", index.Mapping().(*mapping.IndexMappingImpl).DefaultAnalyzer)
}

func TestIndexing(t *testing.T) {

	config.DataPath = "../../testdata" //load data from testdata
	config.IndexPath = "./index"

	index := newIndex()

	defer closeAndDestroy(index, config.IndexPath)


	count, _ := indexing(index)

	assert.Equal(t, 15, count)
}



func TestIndexingFail(t *testing.T) {

	config.DataPath = "./anywhere"
	config.IndexPath = "./index"

	index := newIndex()
	defer closeAndDestroy(index, config.IndexPath)

	_, err := indexing(index)

	assert.NotNil(t, err)



}
