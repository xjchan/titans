package tools

import (
	"errors"
	// "fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

const CONFIG_FILES = "./configs/config_files.yaml"

//Builder 构造配置(config)
type Builder interface {
	Config([]byte) (interface{}, error)
}

var configFiles = make(map[string]string, 1)
var configBytes = make(map[string][]byte, 1)
var configs = make(map[string]interface{}, 1)

//初始化设置
func init() {
	var files interface{}
	var err error
	var buf []byte

	f, err := os.Open(CONFIG_FILES)
	CheckError(err)

	buf, err = ioutil.ReadAll(f)
	CheckError(err)

	err = yaml.Unmarshal(buf, &files)
	CheckError(err)

	configMap, ok := files.(map[interface{}]interface{})

	if ok {
		for k, v := range configMap {
			key, okk := k.(string)
			value, okv := v.(string)

			if okk && okv {
				configFiles[key] = value
			}
		}
	}

	for k, v := range configFiles {
		reader, _ := os.Open(v)
		buf, _ := ioutil.ReadAll(reader)
		configBytes[k] = buf
	}
}

//GetConfigFunc 获取
func GetConfigFunc(key string, builder func(buf []byte) (interface{}, error)) (interface{}, error) {
	if configs[key] != nil { //已经初始化，直接返回
		return configs[key], nil
	} else if configBytes[key] == nil { //不存在对应的文件，返回错误
		return nil, errors.New("There is no config of key:" + key + ", please check it\n")
	} else { //写入
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
