package validate

import (
	"image"
	"mime/multipart"
	"regexp"
	"strconv"
	"strings"

	"gin_bbs/pkg/mimetype"
)

var (
	// 匹配电子邮箱
	emailReg = regexp.MustCompile(`\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`)
	// 匹配手机
	phoneReg = regexp.MustCompile(`^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199)\d{8}$`)
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

// MinLengthValidator -
func MinLengthValidator(value string, minStrLen int) ValidatorFunc {
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
		if value == "" {
			return ""
		}
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
		if value == "" {
			return ""
		}
		status := emailReg.MatchString(value)

		if !status {
			return "$name 邮箱格式错误"
		}

		return ""
	}
}

// PhoneValidator 验证手机格式
func PhoneValidator(value string) ValidatorFunc {
	return func() string {
		if value == "" {
			return ""
		}
		status := phoneReg.MatchString(value)

		if !status {
			return "$name 手机格式错误"
		}

		return ""
	}
}

// UintRangeValidator value 是否存在于指定的 range 范围内
func UintRangeValidator(value uint, rg []uint) ValidatorFunc {
	return func() string {
		for _, r := range rg {
			if r == value {
				return ""
			}
		}

		return "$name 不存在于指定范围内"
	}
}

// StringRangeValidator value 是否存在于指定的 range 范围内
func StringRangeValidator(value string, rg []string) ValidatorFunc {
	return func() string {
		for _, r := range rg {
			if r == value {
				return ""
			}
		}

		return "$name 不存在于指定范围内"
	}
}

// MimetypeValidator 文件 mimetype 校验
func MimetypeValidator(f *multipart.FileHeader, mimes []string) ValidatorFunc {
	return func() string {
		if f == nil {
			return ""
		}
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

		return "$name 格式错误，支持的格式为 " + strings.Join(mimes, ",")
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
		if f == nil {
			return ""
		}
		src, err := f.Open()
		if err != nil {
			return "$name 打开失败"
		}
		defer src.Close()

		config, _, err := image.DecodeConfig(src)
		if err != nil {
			return "$name 解码失败"
		}
		/* 或者也可这样获取图片宽高
		   img, _, err := image.Decode(src)
		   if err != nil {
		     return "$name 解码失败"
		   }
		   imgWidth := img.Bounds().Dx()
		   imgHeight := img.Bounds().Dy()
		*/

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
