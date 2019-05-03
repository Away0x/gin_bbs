package validate

import (
	"regexp"
	"strconv"
)

// RequiredValidator : value 必须存在
func RequiredValidator(value string) ValidatorFunc {
	return func() (msg string) {
		if value == "" {
			return "$name 必须存在"
		}

		return ""
	}
}

// MixLengthValidator -
func MixLengthValidator(value string, minStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)

		if l < minStrLen {
			return "$name 必须大于 " + strconv.Itoa(minStrLen)
		}

		return ""
	}
}

// MaxLengthValidator -
func MaxLengthValidator(value string, maxStrLen int) ValidatorFunc {
	return func() (msg string) {
		l := len(value)

		if l > maxStrLen {
			return "$name 必须小于 " + strconv.Itoa(maxStrLen)
		}

		return ""
	}
}

// EqualValidator -
func EqualValidator(v1, v2 string, other ...string) ValidatorFunc {
	return func() (msg string) {
		if v1 != v2 {
			if len(other) != 0 && other[0] != "" {
				return "$name 不等于 " + other[0]
			} else {
				return "$name 验证失败"
			}
		}

		return ""
	}
}

// EmailValidator 验证邮箱格式
func EmailValidator(value string) ValidatorFunc {
	return func() (msg string) {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		status := reg.MatchString(value)

		if !status {
			return "$name 邮箱格式错误"
		}

		return ""
	}
}
