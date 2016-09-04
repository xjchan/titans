package mongoDB

import (
	// "github.com/xjchan/titans/server/app"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

// var configURI = app.Dir() + "/app/configs/DB/mongoDB.yaml"

type config struct {
	URL string `yaml:"url"`
}

// Connnet to make a connetion to for the DB
func Connnet() *mgo.Session {
	config := Config()
	url := config.URL
	fmt.Println(url)
	session, _ := mgo.Dial(url)
	return session
}

func Config() config {
	config := config{}
	dir, _ := os.Getwd()
	configURI := dir + "/server/app/configs/DB/mongoDB.yaml"

	file, _ := os.Open(configURI)

	// fmt.Println(configURI)

	data, _ := ioutil.ReadAll(file)

	// fmt.Println(string(data))

	yaml.Unmarshal(data, &config)
	fmt.Println(config)
	return config
}
