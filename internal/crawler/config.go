package crawler

import (
	"github.com/spf13/viper"
	"log"
)

// data path
var dataPath string


// load config
func init() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	// if in project root
	viper.AddConfigPath(".")
	// if you are in cmd/crawler or internal/crawler
	viper.AddConfigPath("../..")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	dataPath = viper.GetString("crawlerDataPath")

}