package validate

import (
	"strings"

	"gin_bbs/pkg/ginutils/flash"

	"encoding/json"

	"github.com/gin-gonic/gin"
)

const (
	ValidateMessageArrKeyName = "validatorMessageArr"
	ValidateMessageMapKeyName = "validatorMessageMap"
	ValidateSeparator         = "$$|$$"
)

// SaveValidateMessage 存储参数验证的错误信息
func SaveValidateMessage(c *gin.Context, errArr []string, errMap map[string][]string) {
	SaveValidateMessageArr(c, errArr)
	SaveValidateMessageMap(c, errMap)
}

// ReadValidateMessage 读取参数验证的错误信息
func ReadValidateMessage(c *gin.Context) ([]string, map[string][]string) {
	errArr := ReadValidateMessageArr(c)
	errMap := ReadValidateMessageMap(c)
	return errArr, errMap
}

// SaveValidateMessageArr 存储参数验证的错误信息
func SaveValidateMessageArr(c *gin.Context, msgArr []string) {
	f := flash.NewFlashByName(ValidateMessageArrKeyName)
	f.Data = map[string]string{
		"errors": strings.Join(msgArr, ValidateSeparator),
	}
	f.SaveByName(c, ValidateMessageArrKeyName)
}

// ReadValidateMessageArr 读取参数验证的错误信息
func ReadValidateMessageArr(c *gin.Context) []string {
	errorStr := flash.ReadByName(c, ValidateMessageArrKeyName).Data["errors"]

	if errorStr == "" {
		return []string{}
	}
	// 不做上面的判断，Split 切分空字符串会得 [""]
	return strings.Split(errorStr, ValidateSeparator)
}

// SaveValidateMessageMap 存储参数验证的错误信息
func SaveValidateMessageMap(c *gin.Context, msgMap map[string][]string) {
	f := flash.NewFlashByName(ValidateMessageMapKeyName)
	json_str, _ := json.Marshal(msgMap)
	f.Data = map[string]string{
		"errors": string(json_str),
	}
	f.SaveByName(c, ValidateMessageMapKeyName)
}

// ReadValidateMessageMap 读取参数验证的错误信息
func ReadValidateMessageMap(c *gin.Context) map[string][]string {
	data := make(map[string][]string)
	errorStr := flash.ReadByName(c, ValidateMessageMapKeyName).Data["errors"]

	if errorStr == "" {
		return data
	}

	json.Unmarshal([]byte(errorStr), &data)
	return data
}

// AddMessageAndSaveToFlash 添加错误信息并且保存到 flash 中
func AddMessageAndSaveToFlash(c *gin.Context, keyName, msg string, errArr []string, errMap MessagesMap) {
	errArr, errMap = AddMessage(keyName, msg, errArr, errMap)
	SaveValidateMessage(c, errArr, errMap)
}
