package utils

import "encoding/json"

// JsonToString 将任意结构体转换为 JSON 字符串
func JsonToString(v interface{}) (string, error) {
	bytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
