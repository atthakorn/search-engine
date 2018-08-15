package crawler

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"log"
	"github.com/atthakorn/web-scraper/internal/config"
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
	if f, err := os.Create(filename); err != nil {
		return err
	} else {
		defer f.Close()

		if _,err := f.WriteString(content); err != nil {
			log.Printf("File not found: %v", f.Name())
		}

		return nil
	}
}


func GetDataPath(host string) string {

	return filepath.Join(config.DataPath, host) + ".json"
}
