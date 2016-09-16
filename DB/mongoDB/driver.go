package mongoDB

import (
	"github.com/xjchan/titans/tools"
	"gopkg.in/mgo.v2"
	"gopkg.in/yaml.v2"
)

var conf = mongoConfig{}

//mongoConfig mongoConfig
type mongoConfig struct {
	URL      string `yaml:"url"`
	Database string `yaml:"database"`
}

func init() {
	c, err := tools.GetConfig("mongoDB", conf)
	if err == nil {
		value, ok := c.(mongoConfig)
		if ok {
			conf = value
		}
	}
}

func GetConfig() mongoConfig {
	return conf
}

//config.Builder interface
func (config mongoConfig) Config(buf []byte) (interface{}, error) {
	err := yaml.Unmarshal(buf, &config)
	return config, err
}

// Connnet to make a connetion to for the DB
func Connnet() (*mgo.Session, error) {
	url := conf.URL

	// fmt.Println(url)
	session, err := mgo.Dial(url)
	// fmt.Println("error:" + err.Error())
	if err == nil {
		return session, nil
	}
	return nil, err
}

func Close(s *mgo.Session) {
	s.Close()
}
