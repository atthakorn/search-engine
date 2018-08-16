package index

import (
	"github.com/atthakorn/web-scraper/internal/config"
	"github.com/stretchr/testify/mock"
	"github.com/blevesearch/bleve"
)


func init() {
	config.DataPath = "../../data/crawler"
	config.IndexPath = "../../data/index"
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



/*func TestCreateIndexMapping (t *testing.T) {

	indexMapping := buildIndexMapping().(*mapping.IndexMappingImpl)


	assert.Equal(t,"th", indexMapping.DefaultAnalyzer )
}



func TestBenchmark(t *testing.T) {

	mockObject := &indexMock{}

	var index bleve.Index

	mockObject.On("openOrCreate").Return(index)
	mockObject.On("indexing", mockObject.indexing).Return(0,nil)


	benchmark := benchmark(index, mockObject.indexing)
	benchmark()




}*/