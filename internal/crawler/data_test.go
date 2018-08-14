package crawler

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMarshall(t *testing.T) {

	//mock data
	data := Data{
		Title: "test",
		URL:   "http://test.com",
		Texts: []string{"test"},
	}

	//test marshal
	json := Marshal([]Data{data})
	assert.Equal(t, `[{"Title":"test","URL":"http://test.com","Texts":["test"]}]`, json)


	//test unmarshal
	var restore []Data
	Unmarshal(json, &restore)
	assert.True(t, assert.ObjectsAreEqualValues( data, restore[0]))

}
