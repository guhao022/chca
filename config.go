package chca

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

func load() ([]byte, error) {

	file, err := os.Open(confile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

func Config() Website {
	data, err := load()
	if err != nil {
		panic("加载配置文件失败" + err.Error())
	}

	w := Website{}
	err = yaml.Unmarshal(data, &w)
	if err != nil {
		panic("解析配置文件失败" + err.Error())
	}

	return w
}
