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
func WriteString(filename string, content string) {

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(content); err != nil {
		panic(err)
	}
}


func GetDataPath(host string) string {
	return fmt.Sprintf("%s/%s.json", dataPath, host)
}


