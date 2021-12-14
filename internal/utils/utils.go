package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"fofax/internal/printer"
	"os"
	"reflect"
)

// StructToMap Struct 转 Map
func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // 非结构体返回错误提示
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}

	t := v.Type()
	// 遍历结构体字段
	// 指定tagName值为map中key;字段值为map中value
	for i := 0; i < v.NumField(); i++ {
		fi := t.Field(i)
		if tagValue := fi.Tag.Get(tagName); tagValue != "" {
			out[tagValue] = v.Field(i).Interface()
		}
	}
	return out, nil

}

func JsonStrToMAp(jsonStr string) (map[string]string, error) {
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func MapToJsonStr(m map[string]string) (string, error) {
	jsonByte, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(jsonByte), nil
}

func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		printer.Errorf("File  %s not exist", path)
		return false
	}
	return false
}

func FileContent2List(path string) ([]string, error) {
	urlSlice := make([]string, 100)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println(fileScanner.Text())
		urlSlice = append(urlSlice, fileScanner.Text())
	}
	if err := fileScanner.Err(); err != nil {
		return nil, err
	}

	file.Close()
	return urlSlice, nil
}

func HasStdin() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}

	mode := stat.Mode()

	isPipedFromChrDev := (mode & os.ModeCharDevice) == 0
	isPipedFromFIFO := (mode & os.ModeNamedPipe) != 0

	return isPipedFromChrDev || isPipedFromFIFO
}

func GetHidePasswd(key string) string {
	var HiddenPasswdChar = "*******************"
	if len(key) == 32 {
		return key[0:3] + HiddenPasswdChar + key[28:]
	}
	return ""
}

