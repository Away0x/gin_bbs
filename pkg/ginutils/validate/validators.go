package validate

import (
	"image"
	"mime/multipart"
	"regexp"
	"strconv"

	"gin_bbs/pkg/mimetype"
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

// MimetypeValidator 文件 mimetype 校验
func MimetypeValidator(f *multipart.FileHeader, mimes []string) ValidatorFunc {
	return func() string {
		src, err := f.Open()
		if err != nil {
			return "$name 打开失败"
		}
		defer src.Close()

		_, ext, err := mimetype.DetectReader(src)
		if err != nil {
			return "$name 解码失败"
		}

		for _, m := range mimes {
			if m == ext {
				return ""
			}
		}

		return "$name 格式错误"
	}
}

// DimensionsOptions -
type DimensionsOptions struct {
	MinWidth  int // px
	MinHeight int // px
	MaxWidth  int // px
	MaxHeight int // px
}

// ImageDimensionsValidator 图片分辨率校验
func ImageDimensionsValidator(f *multipart.FileHeader, options DimensionsOptions) ValidatorFunc {
	return func() string {
		src, err := f.Open()
		if err != nil {
			return "$name 打开失败"
		}
		defer src.Close()

		config, _, err := image.DecodeConfig(src)
		if err != nil {
			return "$name 解码失败"
		}

		if options.MinWidth != 0 {
			if config.Width < options.MinWidth {
				return "$name 宽度不能小于 " + strconv.Itoa(options.MinWidth) + "px"
			}
		}

		if options.MinHeight != 0 {
			if config.Height < options.MinHeight {
				return "$name 高度不能小于 " + strconv.Itoa(options.MinHeight) + "px"
			}
		}

		if options.MaxWidth != 0 {
			if config.Width > options.MaxWidth {
				return "$name 宽度不能大于 " + strconv.Itoa(options.MaxWidth) + "px"
			}
		}

		if options.MaxHeight != 0 {
			if config.Height > options.MaxHeight {
				return "$name 高度不能大于 " + strconv.Itoa(options.MaxHeight) + "px"
			}
		}

		return ""
	}
}
