

package crawler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PuerkitoBio/goquery"
	"strings"
)


func TestMarshalUnmarshal(t *testing.T) {

	//mock d1
	d1 := Data{
		Title: "test",
		URL:   "http://test.com",
		Texts: []string{"line 1", "line 2"},
	}

	//test marshal
	json := Marshal([]Data{d1})
	assert.Equal(t, `[{"title":"test","url":"http://test.com","texts":["line 1","line 2"]}]`, json)

	//test unmarshal
	var d2 []Data
	Unmarshal(json, &d2)
	assert.True(t, assert.ObjectsAreEqualValues(d1, d2[0]))

}


func TestParseText(t *testing.T) {
	html := `<html lang="en">
				<head>
					<link rel="stylesheet" href="https://domain.com/asset.css">    
				<style>
    			body{
      				min-height: 100vh;
          			}
				</style>
  			</head>
  				<body>
			    	line 1
					line 2
					<script src="https://domain.com/asset.js"></script>
  				</body>
			</html>`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	//test ParseTexts
	texts := ParseTexts(doc.Find("html"))

	assert.Equal(t, "line 1", texts[0])
	assert.Equal(t, "line 2", texts[1])

}
