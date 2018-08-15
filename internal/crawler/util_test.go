package crawler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
	"github.com/atthakorn/web-scraper/internal/config"
)


func TestGetDataPath(t *testing.T) {

	config.DataPath = "../../data"

	file := GetDataPath("www.domain.com")


	assert.Equal(t, "../../data/www.domain.com.json", file)
}



func TestReadWriteString(t *testing.T) {

	//set data path related to test
	config.DataPath = "../../data"

	content :=  "test"

	file := GetDataPath("www.domain.com")


	WriteString(file, content)

	s, _ := LoadString(file)

	assert.Equal(t, content, s)

	//clean up file
	os.Remove(file)

}


func TestReadFileNotFound(t *testing.T) {

	_, err := LoadString("www.domain.com")

	assert.NotNil(t,err)

}