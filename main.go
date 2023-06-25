package main

import (
	"github.com/alomia/fastman/cmd"
	"github.com/spf13/viper"
)

func main() {
	viper.AddConfigPath(".")
	viper.SetConfigName("fastmanconf")
	viper.SetConfigType("yaml")
	viper.ReadInConfig()

	cmd.Execute()
}
