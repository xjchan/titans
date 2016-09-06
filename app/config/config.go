package config

import (
	// // "encoding/json"
	"fmt"
	// // json "github.com/pquerna/ffjson/ffjson"
	// "gopkg.in/yaml.v2"
	"errors"
	"io/ioutil"
	"os"
	// "test/test/test"
	// "test/test/test1"
)

//Builder 构造配置(config)
type Builder interface {
	Config([]byte) (interface{}, error)
}

var configFiles = make(map[string]string, 1)
var configBytes = make(map[string][]byte, 1)
var configs = make(map[string]interface{}, 1)

//初始化设置
func init() {
	configFiles["a"] = "./config/config.yaml"
	for k, v := range configFiles {
		reader, _ := os.Open(v)
		buf, _ := ioutil.ReadAll(reader)
		configBytes[k] = buf
	}
}

//GetConfigFunc 获取
func GetConfigFunc(key string, builder func(buf []byte) (interface{}, error)) (interface{}, error) {
	if configs[key] != nil { //已经初始化，直接返回
		fmt.Println(1)
		return configs[key], nil
	} else if configBytes[key] == nil { //不存在对应的文件，返回错误
		fmt.Println(2)
		return nil, errors.New("There is no config of key:" + key + ", please check it\n")
	} else { //写入
		fmt.Println(3)
		r, err := builder(configBytes[key])
		if err == nil {
			configs[key] = r
		}
		return r, err
	}
}

//GetConfig 获取
func GetConfig(key string, builder Builder) (interface{}, error) {
	if configs[key] != nil { //已经初始化，直接返回
		return configs[key], nil
	} else if configBytes[key] == nil { //不存在对应的文件，返回错误
		return nil, errors.New("There is no config of key:" + key + ", please check it\n")
	} else { //写入
		r, err := builder.Config(configBytes[key])
		if err == nil {
			configs[key] = r
		}
		return r, err
	}
}
