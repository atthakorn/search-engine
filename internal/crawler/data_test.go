package crawler

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/PuerkitoBio/goquery"
	"strings"
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
	assert.True(t, assert.ObjectsAreEqualValues(data, restore[0]))

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
  				<body class="body--12">
			    	test
					<script src="https://domain.com/asset.js"></script>
  				</body>
			</html>`

	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(html))

	text := ParseText(doc.Find("html"))

	//test ParseText
	assert.Equal(t, "\n\t\t\t\t\t    \n\t\t\t\t\n  \t\t\t\n  \t\t\t\t\n\t\t\t    \ttest\n\t\t\t\t\t\n  \t\t\t\t\n\t\t\t", text)



	//test ParseTexts
	texts := ParseTexts(doc.Find("html"))

	assert.Equal(t, "test", texts[0])

}
