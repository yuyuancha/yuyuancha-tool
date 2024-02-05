package tool

import (
	"regexp"
)

// SplitStrByRegex 依照 regex 切割字串
func SplitStrByRegex(str, regex string) string {
	re := regexp.MustCompile(regex)
	results := re.Split(str, -1)

	if len(results) < 2 {
		return ""
	}

	return results[1]
}
