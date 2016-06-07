package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type baseConfig struct {
	HttpPort string
	RpcPort  string
	Env      string
}

type myConfig struct {
	baseConfig
	DbName   string
	DbUrl    string
	Services map[string]serviceConf
}

type serviceConf struct {
	Host     string
	HttpPort string
	RpcPort  string
}

func main() {

	viper.SetConfigName("config_viper")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// viper.GetString(key)

	var (
		apiConfig = myConfig{}

		userConfig = myConfig{}
	)

	viper.UnmarshalKey("api.dev", &apiConfig)

	viper.UnmarshalKey("api.dev.baseConfig", &apiConfig.baseConfig)
	fmt.Printf("%+v\n", apiConfig)

	println("-----")

	viper.UnmarshalKey("user.staging", &userConfig)
	fmt.Printf("%+v\n", userConfig)

}
