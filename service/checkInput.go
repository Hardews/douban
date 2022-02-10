package service

import "strings"

func CheckLength(word string) bool {
	if len(word) < 6 || len(word) > 20 {
		return false
	}
	return true
}

func CheckTxtLengthS(word string) bool {
	if len(word) > 100 {
		return false
	}
	return true
}

func CheckTxtLengthL(word string) bool {
	if len(word) > 150 {
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

var administrators = make([]string, 0) //管理员账号切片

func CheckAdministratorUsername(username string) bool {
	administrators = append(administrators, "1225101127")

	for i, _ := range administrators {
		flag := strings.Compare(username, administrators[i])
		if flag != 0 {
			return false
		}
	}
	return true
}
