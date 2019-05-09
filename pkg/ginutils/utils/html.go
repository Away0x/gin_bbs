package utils

import (
	"regexp"
)

// XSSClean 去除 html string 里面的 style script 标签
func XSSClean(content string) string {
	// 去除 STYLE
	re, _ := regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	content = re.ReplaceAllString(content, "")
	//去除 SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	content = re.ReplaceAllString(content, "")

	return content
}
