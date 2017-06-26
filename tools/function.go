package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type Config struct {
	Source string
	Api    []Item
}

type Item struct {
	Article  string
	Category string
}

const configFileSizeLimit = 10 << 20

//解析config.json配置文件
func ParseConfig() (config Config) {
	path := "./conf/config.json"
	config_file, err := os.Open(path)
	if err != nil {
		fmt.Println("Failed to open config file '%s': %s\n", path, err)
		return
	}

	fi, _ := config_file.Stat()
	if size := fi.Size(); size > (configFileSizeLimit) {
		fmt.Println("config file (%q) size exceeds reasonable limit (%d) - aborting", path, size)
		return
	}

	if fi.Size() == 0 {
		fmt.Println("config file (%q) is empty, skipping", path)
		return
	}

	buffer := make([]byte, fi.Size())
	_, err = config_file.Read(buffer)

	buffer, err = StripComments(buffer) //去掉注释
	if err != nil {
		fmt.Println("Failed to strip comments from json: %s\n", err)
		return
	}

	buffer = []byte(os.ExpandEnv(string(buffer))) //特殊

	err = json.Unmarshal(buffer, &config) //解析json格式数据
	if err != nil {
		fmt.Println("Failed unmarshalling json: %s\n", err)
		return
	}
	return config
}

//去除json文件中的注释
func StripComments(data []byte) ([]byte, error) {
	data = bytes.Replace(data, []byte("\r"), []byte(""), 0) // Windows
	lines := bytes.Split(data, []byte("\n"))                //split to muli lines
	filtered := make([][]byte, 0)

	for _, line := range lines {
		match, err := regexp.Match(`^\s*#`, line)
		if err != nil {
			return nil, err
		}
		if !match {
			filtered = append(filtered, line)
		}
	}

	return bytes.Join(filtered, []byte("\n")), nil
}

//HTTP Get请求
func HttpGet(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	resp, _ := http.DefaultClient.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	return string(body)
}

//检查文件是否存在
func CheckFileIsExist(filename string) bool {
	//创建dist目录
	os.MkdirAll("./static/dist/", 0777)

	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

//读取文件内容
func ReadFile(file string) ([]byte, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
