package word

import (
	"strings"
	"unicode"
)

/*
单词全部转为小写。
单词全部转为大写。
下画线单词转为大写驼峰单词。
下画线单词转为小写驼峰单词。
驼峰单词转为下画线单词。
 */

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func ToLower(s string) string  {
	return strings.ToLower(s)
}

func UnderscoreToUpperCamelCase(s string) string  {
	s = strings.Replace(s,"_"," ",-1)

	s = strings.Title(s)

	return strings.Replace(s," ","",-1)
}

func UnderscoreToLowerCamelCase(s string) string {
	s = UnderscoreToUpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i,r := range s {
		if i == 0 {
			output = append(output,unicode.ToLower(r))
			continue
		}

		if unicode.IsUpper(r) {
			output = append(output,'_')
		}
		output = append(output,unicode.ToLower(r))
	}

	return string(output)
}
