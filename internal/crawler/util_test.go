package crawler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"os"
)



func TestGetDataPath(t *testing.T) {

	file := GetDataPath("www.domain.com")

	assert.Equal(t, "./data/www.domain.com.json", file)
}



func TestReadWriteString(t *testing.T) {

	//set data path related to test
	dataPath = "../../data"
	content :=  "test"

	file := GetDataPath("www.domain.com")

	WriteString(file, content)

	s, _ := LoadString(file)

	assert.Equal(t, content, s)

	//remove file
	os.Remove(file)

}


func TestReadFileNotFound(t *testing.T) {

	_, err := LoadString("www.domain.com")

	assert.NotNil(t,err)


}