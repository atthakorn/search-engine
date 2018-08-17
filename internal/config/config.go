package config

import (
	"github.com/spf13/viper"
	"log"
)

// data path
var (
	EntryPoints []string
	MaxDepth    int
	Parallelism int
	Delay       int
	DataPath    string
	IndexPath	string
)

// load config
func init() {

	viper.SetConfigName(".config")
	viper.SetConfigType("yaml")
	// if in project root
	viper.AddConfigPath(".")
	// if you are in cmd/crawler or internal/crawler
	viper.AddConfigPath("../..")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}


	EntryPoints = viper.GetStringSlice("entryPoint")
	MaxDepth = viper.GetInt("maxDepth")
	Parallelism = viper.GetInt("parallelism")
	Delay = viper.GetInt("delay")
	DataPath = viper.GetString("dataPath")
	IndexPath = viper.GetString("indexPath")

}