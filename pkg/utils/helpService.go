package utils

import (
	"crypto/rand"
	"encoding/base64"
	"papergen/internal/global"
	"strconv"
	"strings"
)

// GenerateRandomString 生成指定长度的随机字符串
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length], nil
}

func StringArrToIntArr(strArr []string) ([]int, error) {
	intArr := make([]int, len(strArr))

	// 遍历分割后的字符串数组，将每个元素转换为 int
	for i, s := range strArr {
		num, err := strconv.Atoi(strings.TrimSpace(s)) // 去除可能的空格并转换为 int
		if err != nil {
			global.Logger.Error("转换错误: " + err.Error() + "\n")
			return nil, err
		}
		intArr[i] = num
	}
	return intArr, nil
}
