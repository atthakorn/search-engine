package crawler

import (
	"github.com/spf13/viper"
	"log"
	"fmt"
)

func init() {


	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}


	fmt.Printf("Fetch Sites: %v\n", viper.GetStringSlice("sites"))


}
