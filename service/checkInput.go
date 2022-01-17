package service

import "strings"

func CheckLength(word string) bool {
	if len(word) < 6 || len(word) > 20 {
		return false
	}
	return true
}

var SensitiveWords = make([]string, 0) //敏感词汇切片

func CheckSensitiveWords(word string) bool {
	SensitiveWords = append(SensitiveWords, "fuck")
	SensitiveWords = append(SensitiveWords, "傻逼")
	SensitiveWords = append(SensitiveWords, "sb")
	SensitiveWords = append(SensitiveWords, "SB")
	SensitiveWords = append(SensitiveWords, "你妈")
	SensitiveWords = append(SensitiveWords, "狗东西")

	for i, _ := range SensitiveWords {
		flag := strings.Contains(word, SensitiveWords[i])
		if flag {
			return false
		}
	}
	return true
}
