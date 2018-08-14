package crawler

import (
	"testing"
	"github.com/stretchr/testify/assert"
)



func TestGetDataPath(t *testing.T) {

	dataPath := GetDataPath("www.domain.com")

	assert.Equal(t, dataPath, "./data/www.domain.com.json")
}


