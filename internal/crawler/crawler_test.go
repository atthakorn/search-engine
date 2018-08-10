package crawler

import (
	"testing"
	"github.com/spf13/viper"
)


func TestLoadConfig(t *testing.T) {

	sites := viper.GetStringSlice("sites")
	maxDepth := viper.GetInt("maxDepth")
	parallelism := viper.GetInt("parallelism")
	delay := viper.GetInt("delay")

	if len(sites) == 0 {
		t.Errorf("sites must contain at least one element")
	}
	if maxDepth == 0 {
		t.Errorf("maxDepth must be greater than zero")
	}
	if parallelism == 0 {
		t.Errorf("parallelism must be greater than zero")
	}
	if delay == 0 {
		t.Errorf("delay  must be greater than zero")
	}

}