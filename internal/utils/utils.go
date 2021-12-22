package utils

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/browser"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
)

var ConfDefaultPath = []string{
	"fofax.yaml",
	filepath.Join(getHomedir(), "fofax.yaml"),
	"/etc/fofa.yaml",
}

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

// ConvertByte2String 解决 windows cmd 下不能正确显示中文的问题
func ConvertByte2String(byte []byte, charset string) string {
	var str string
	switch charset {
	case "GB18030":
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case "UTF8":
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func GetHidePasswd(key string) string {
	var HiddenPasswdChar = "*******************"
	if len(key) == 32 {
		return key[0:3] + HiddenPasswdChar + key[28:]
	}
	return ""
}

func GetDefaultConf() string {
	for _, cfile := range ConfDefaultPath {
		if FileExist(cfile) {
			return cfile
		}
	}
	if runtime.GOOS == "windows" {
		return ConfDefaultPath[0]
	} else {
		return ConfDefaultPath[1]
	}
}

func BinBaseDir() string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	return filepath.Dir(ex)
}

func getHomedir() (home string) {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, ".config", "fofax")
}

func OpenFofa(query string) {
	b := []byte(query)

	qbase64 := base64.StdEncoding.EncodeToString(b)
	browser.OpenURL(fmt.Sprintf("https://fofa.so/result?qbase64=%s", qbase64))
}
