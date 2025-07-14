package main

import (
	"cassandraClient/redis"
	"github.com/gookit/config/v2"
	"gopkg.in/yaml.v2"
	"strconv"
)
import _ "embed"

//go:embed config/config.yml
var configFile string

// Ref: https://docs.aws.amazon.com/keyspaces/latest/devguide/using_go_driver.html
func main() {

	// 8. Redis
	redis.RedisRunner()

	// 7. Interfaces play
	//interfaces.RunSeq()

	// 6. Regexp
	//regexpress.RunRegExp()

	// 5: Write file
	//file.ReadFile()
	//file.WriteFile()

	// Step Zero: read config from resources
	//fullUrl, maxNumberOfRequestsAsInt, err := readYamlConfig()
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(fullUrl)
	//fmt.Println(maxNumberOfRequestsAsInt)

	//////////////////////////////////////////////////////////////////
	// 4: Http client work
	//httpClient.RunNonBlockingV2(fullUrl, maxNumberOfRequestsAsInt)
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

func readYamlConfig() (string, int, error) {
	err := converseToYaml()
	if err != nil {
		return "", -1, err
	}
	targetHostname := config.String("target_hostname")
	pageName := config.String("page_name")
	maxNumberOfRequests := config.String("maxNumberOfRequests")
	maxRequestsAsInt, err := strconv.Atoi(maxNumberOfRequests)
	return targetHostname + "/" + pageName, maxRequestsAsInt, err
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
