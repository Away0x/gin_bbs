package validate

import (
	"regexp"
	"strconv"
)

// RequiredValidator : value 必须存在
func RequiredValidator(value string) ValidatorFunc {
	return func() string {
		if value == "" {
			return "$name 必须存在"
		}

		return ""
	}
}

// MixLengthValidator -
func MixLengthValidator(value string, minStrLen int) ValidatorFunc {
	return func() string {
		l := len(value)

		if l < minStrLen {
			return "$name 必须大于 " + strconv.Itoa(minStrLen)
		}

		return ""
	}
}

// MaxLengthValidator -
func MaxLengthValidator(value string, maxStrLen int) ValidatorFunc {
	return func() string {
		l := len(value)

		if l > maxStrLen {
			return "$name 必须小于 " + strconv.Itoa(maxStrLen)
		}

		return ""
	}
}

// BetweenValidator - 限制字段的 string length
func BetweenValidator(value string, minStrLen, maxStrLen int) ValidatorFunc {
	return func() string {
		l := len(value)

		if l > maxStrLen || l < minStrLen {
			return "$name 必须介于" + strconv.Itoa(minStrLen) + "-" + strconv.Itoa(maxStrLen) + "个字符之间"
		}

		return ""
	}
}

// RegexpValidator 正则验证
func RegexpValidator(value string, regexpStr string) ValidatorFunc {
	return func() string {
		ok, err := regexp.MatchString(regexpStr, value)
		if !ok || err != nil {
			return "$name 格式错误"
		}

		return ""
	}
}

// EqualValidator -
func EqualValidator(v1, v2 string, other ...string) ValidatorFunc {
	return func() string {
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
	return func() string {
		pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` // 匹配电子邮箱
		reg := regexp.MustCompile(pattern)
		status := reg.MatchString(value)

		if !status {
			return "$name 邮箱格式错误"
		}

		return ""
	}
}
