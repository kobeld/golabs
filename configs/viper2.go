package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type appConfig struct {
	baseConfig
	DbUrl string
}

type baseConfig struct {
	HttpPort string
	RpcPort  string
	Env      string
}

func (this baseConfig) ReadInFile(path string) {
	viper.SetConfigFile(path)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.UnmarshalKey("api.dev.baseConfig", &this)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", this)
}

// type myConfig struct {
// 	baseConfig
// 	DbName   string
// 	DbUrl    string
// 	Services map[string]serviceConf
// }

// type serviceConf struct {
// 	Host     string
// 	HttpPort string
// 	RpcPort  string
// }

func main() {

	// viper.GetString(key)

	var (
		file    = "./config_viper_2.yaml"
		appConf = appConfig{}
	)

	appConf.ReadInFile(file)

	// viper.UnmarshalKey("api.dev", &apiConfig)

	// viper.UnmarshalKey("api.dev.baseConfig", &apiConfig.baseConfig)
	fmt.Printf("%+v\n", appConf)

	// println("-----")

	// viper.UnmarshalKey("user.staging", &userConfig)
	// fmt.Printf("%+v\n", userConfig)

}
