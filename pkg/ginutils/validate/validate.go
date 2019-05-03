package validate

import (
	"gin_bbs/pkg/ginutils/utils"
)

type (
	// 验证器函数
	ValidatorFunc = func() (msg string)
	// 验证器数组 map
	ValidatorMap = map[string][]ValidatorFunc
	// 错误信息数组
	MessagesMap = map[string][]string
)

type IValidate interface {
	// IsStrict : 严格模式时，第一个验证出错时，即会停止其他验证
	IsStrict() bool
	// RegisterValidators : 注册验证器 map
	RegisterValidators() ValidatorMap
	// RegisterMessages : 注册错误信息 map
	RegisterMessages() MessagesMap
}

type Validate struct{}

func (*Validate) IsStrict() bool {
	return false
}

// RegisterValidators: 注册验证器
// 验证器数组按顺序验证，一旦验证没通过，即结束该字段的验证
func (*Validate) RegisterValidators() ValidatorMap {
	return ValidatorMap{}
}

// RegisterMessages 注册错误信息
func (*Validate) RegisterMessages() MessagesMap {
	return MessagesMap{}
}

// 执行验证
func Run(v IValidate) (bool, []string, map[string][]string) {
	return RunInParams(v.IsStrict(), v.RegisterValidators(), v.RegisterMessages())
}

// 执行验证
func RunInParams(strict bool, validatorMap ValidatorMap, messageMap MessagesMap) (ok bool, errArr []string, errMap map[string][]string) {
	errArr = make([]string, 0)
	errMap = make(map[string][]string)
	ok = true

	for key, validators := range validatorMap {
		customMsgArr := messageMap[key] // 自定义的错误信息
		customMsgArrLen := len(customMsgArr)
		errMap[key] = make([]string, 0)

		for i, fn := range validators {
			errMsg := fn() // 执行验证函数
			// 有错误
			if errMsg != "" {
				ok = false

				if i < customMsgArrLen && customMsgArr[i] != "" {
					// 采用自定义的错误信息输出
					errMsg = customMsgArr[i]
				} else {
					// 采用默认的错误信息输出
					errMsg = utils.ParseEasyTemplate(errMsg, map[string]string{
						"$name": key,
					})
				}

				errArr = append(errArr, errMsg)
				errMap[key] = append(errMap[key], errMsg)

				if strict {
					return // 严格模式: 结束所有验证
				} else {
					break // 进行下一个字段的验证
				}
			}
		}
	}

	return
}
