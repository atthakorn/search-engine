package crawler

import (
	"testing"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {

	sites := viper.GetStringSlice("sites")
	maxDepth := viper.GetInt("maxDepth")
	parallelism := viper.GetInt("parallelism")
	delay := viper.GetInt("delay")


	assert.True(t, len(sites) > 0, "Should be greater than zero")
	assert.True(t, maxDepth > 0, "should be greater than zero")
	assert.True(t, parallelism > 0, "Should be greater than zero")
	assert.True(t, delay > 0, "Should be greater than zero")

}
