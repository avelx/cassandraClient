package main

import (
	"fmt"
	"github.com/gookit/config/v2"
	"gopkg.in/yaml.v2"
)
import _ "embed"

//go:embed config/config.yml
var configFile string

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {

	// Step Zero: read config from resources
	fullUrl, err := readYamlConfig()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(fullUrl)

	//////////////////////////////////////////////////////////////////
	// 4: Http client work
	//httpClient.RunNonBlockingSetOfCalls(fullUrl)
	//////////////////////////////////////////////////////////////////

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

func readYamlConfig() (string, error) {
	err := converseToYaml()
	if err != nil {
		return "", err
	}
	targetHostname := config.String("target_hostname")
	pageName := config.String("page_name")
	return targetHostname + "/" + pageName, err
}

func converseToYaml() error {
	var Decoder config.Decoder = yaml.Unmarshal
	var Encoder config.Encoder = yaml.Marshal
	var Driver = config.NewDriver(config.Yaml, Decoder, Encoder).WithAliases(config.Yml)

	config.WithOptions(config.ParseEnv)
	config.AddDriver(Driver)

	// convert embeded file into Yaml
	err := config.LoadSources(config.Yaml, []byte(configFile))
	return err
}
