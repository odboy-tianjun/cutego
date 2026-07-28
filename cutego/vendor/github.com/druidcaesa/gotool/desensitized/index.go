package desensitized

import (
	"bytes"
	"regexp"
	"strings"
)

type Desensitized struct {
}

// DefaultDesensitized 默认:匹配 手机号,邮箱,中文,身份证 等进行脱敏处理
func (d Desensitized) DefaultDesensitized(str string) (result string) {
	var buffer bytes.Buffer
	if str == "" {
		return "***"
	}
	if strings.Contains(str, "@") {
		// 邮箱
		res := strings.Split(str, "@")
		if len(res[0]) < 3 {
			resString := "***"
			buffer.WriteString(resString)
			buffer.WriteString("@")
			buffer.WriteString(res[1])
			result = buffer.String()
		} else {
			res2 := d.subStr(str, 0, 3)
			resString := res2 + "***"
			buffer.WriteString(resString)
			buffer.WriteString("@")
			buffer.WriteString(res[1])
			result = buffer.String()
		}
		return result
	} else {
		reg := `^1[0-9]\d{9}$`
		rgx := regexp.MustCompile(reg)
		mobileMatch := rgx.MatchString(str)
		if mobileMatch {
			// 手机号
			buffer.WriteString(d.subStr(str, 0, 3))
			buffer.WriteString("****")
			buffer.WriteString(d.subStr(str, 7, 11))
			result = buffer.String()
		} else {
			nameRune := []rune(str)
			lens := len(nameRune)
			if lens <= 1 {
				result = "***"
			} else if lens == 2 {
				buffer.WriteString(string(nameRune[:1]))
				buffer.WriteString("*")
				result = buffer.String()
			} else if lens == 3 {
				buffer.WriteString(string(nameRune[:1]))
				buffer.WriteString("*")
				buffer.WriteString(string(nameRune[2:3]))
				result = buffer.String()
			} else if lens == 4 {
				buffer.WriteString(string(nameRune[:1]))
				buffer.WriteString("**")
				buffer.WriteString(string(nameRune[lens-1 : lens]))
				result = buffer.String()
			} else if lens == 18 {
				buffer.WriteString(string(nameRune[:4]))
				buffer.WriteString("**********")
				buffer.WriteString(string(nameRune[lens-4 : lens]))
				result = buffer.String()
			} else {
				i := lens / 3
				buffer.WriteString(string(nameRune[:i]))
				buffer.WriteString(strings.Repeat("*", i))
				buffer.WriteString(d.subStr(str, i*2, len(str)))
				result = buffer.String()
			}
		}
		return
	}
}

// CustomizeHash 自定义机密部分
func (Desensitized) subStr(str string, start int, end int) string {
	rs := []rune(str)
	return string(rs[start:end])
}

// CustomizeHash @param 需要脱敏的字符串
//@param start 需要隐藏开始的位置
//@param end 需要隐藏结束的位置
//自定义脱敏方法
func (d Desensitized) CustomizeHash(str string, start int, end int) string {
	return d.subStr(str, 0, start) + strings.Repeat("*", end-start) + d.subStr(str, end, len(str))
}
