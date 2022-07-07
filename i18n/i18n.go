package i18n

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

//Dictinary 字典
var Dictinary *map[interface{}]interface{}

/**
 * @Title LoadLocales
 * @Description //读取国际化文件
 * @Author Cofeesy 15:04 2022/7/1
 * @Param path string
 * @Return error
 **/
func LoadLocales(path string) error {
	//ReadFile reads the named file and returns the contents.
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		return err
	}
	Dictinary = &m
	return nil
}

/**
 * @Title Translation
 * @Description //翻译
 * @Author Cofeesy 15:08 2022/7/1
 * @Param key string
 * @Return string
 **/
func Translation(key string) string {
	dic := *Dictinary
	keys := strings.Split(key, "")
	for index, path := range keys {
		if len(keys) == (index + 1) {
			for k, v := range dic {
				if k, ok := k.(string); ok {
					if k == path {
						if value, ok := v.(string); ok {
							return value
						}
					}
				}
			}
			return path
		}
		for k, v := range dic {
			if ks, ok := k.(string); ok {
				if ks == path {
					if dic, ok = v.(map[interface{}]interface{}); !ok {
						return path
					}
				}
			} else {
				return ""
			}
		}
	}
	return ""
}
