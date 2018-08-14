package crawler

import (
	"os"
	"io/ioutil"
	"fmt"
)




// LoadString loads json string from file
func LoadString(filename string) (string, error) {

	f, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(f), err
}

// WriteString writes string to file
func WriteString(filename string, content string) error {

	var err error
	if f, err := os.Create(filename); err == nil {
		defer f.Close()
		f.WriteString(content)
	}
	return err
}


func GetDataPath(host string) string {
	return fmt.Sprintf("%s/%s.json", dataPath, host)
}


