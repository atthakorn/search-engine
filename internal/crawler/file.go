package crawler

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"log"
	"github.com/atthakorn/search-engine/internal/config"
)



// LoadString loads string from file
func LoadString(file string) (string, error) {

	f, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(f), nil
}

// WriteString writes string to file
func WriteString(file string, content string) error {
	if f, err := os.Create(file); err != nil {
		return err
	} else {
		defer f.Close()

		if _,err := f.WriteString(content); err != nil {
			log.Printf("File not found: %v", f.Name())
		}

		return nil
	}
}


func JsonFileByHostname(host string) string {

	return filepath.Join(config.DataPath, host) + ".json"
}
