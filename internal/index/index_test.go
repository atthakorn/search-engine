package index

import (
	"github.com/atthakorn/web-scraper/internal/config"
	"github.com/stretchr/testify/mock"
	"github.com/blevesearch/bleve"
	"testing"
	"github.com/blevesearch/bleve/mapping"
	"github.com/stretchr/testify/assert"
	"os"
	"log"
)


func init() {
	config.IndexPath = "./data/index"
	config.DataPath = "./data"
}



type indexMock struct {
	mock.Mock
}


func (m *indexMock) indexing(index bleve.Index) (count int, err error) {

	args := m.Called(index)
	return args.Int(0), args.Error(1)
}

func (m *indexMock) openOrCreate() bleve.Index {
	args := m.Called()
	return args.Get(0).(bleve.Index)
}


func TestCreateIndexMapping (t *testing.T) {

	indexMapping := buildIndexMapping().(*mapping.IndexMappingImpl)


	assert.Equal(t,"th", indexMapping.DefaultAnalyzer )

	assert.IsType(t, &mapping.IndexMappingImpl{}, indexMapping)
}



func TestCreateIndex(t *testing.T) {


	index := openOrCreate()

	defer func() {
		index.Close()
		err := os.RemoveAll(config.DataPath)
		if err != nil {
			log.Printf("cannot remove test data path: %v", err)
		}
	}()

	_, ok := index.(bleve.Index)

	assert.True(t, ok)

	assert.Equal(t, "th", index.Mapping().(*mapping.IndexMappingImpl).DefaultAnalyzer)

}