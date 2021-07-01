package word

import (
	"fmt"
	"strings"
	"unicode"
)

// ToUpper 全部转换为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 全部转换为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线转大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	s = strings.Title(s)
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线转小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		output = append(output, unicode.ToLower(r))
	}
	return string(output)
}

// RemoveAllSpaces 移除字符串中所有空格
func RemoveAllSpaces(s string) string {
	fmt.Println(s)
	s = strings.Replace(s," ","",-1)
	return s
}
