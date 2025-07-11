package main

import (
	"cassandraClient/httpClient"
	"fmt"
	"github.com/gookit/config/v2"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {

	fullUrl := getTargetUrl()
	fmt.Println(fullUrl)

	// 4: Http client work
	httpClient.Run(fullUrl)

	// 3: Bytes ...
	//Play with the bytes
	//messageAsBytes := []byte("Hello World")
	//fmt.Println(messageAsBytes)

	// Convert bytes back to a string
	// Convert array into slice => [:]
	//str := string(messageAsBytes[:])
	//fmt.Println(str)

	// 2.Kafka
	//kafka.Runner()

	// 1: Cassandra
	//cassandra.Runner()
}

func getTargetUrl() string {
	err := getConfig()
	if err != nil {
		panic(err)
	}
	targetHostname := config.String("target_hostname")
	pageName := config.String("page_name")
	return targetHostname + "/" + pageName
}

func getConfig() error {
	var Decoder config.Decoder = yaml.Unmarshal
	var Encoder config.Encoder = yaml.Marshal
	var Driver = config.NewDriver(config.Yaml, Decoder, Encoder).WithAliases(config.Yml)

	config.WithOptions(config.ParseEnv)
	config.AddDriver(Driver)

	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	ex, _ = os.Executable()
	currentFolder := filepath.Dir(ex)

	err = config.LoadFiles(currentFolder + "/config/config.yml")
	return err
}
