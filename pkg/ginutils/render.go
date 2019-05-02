package ginutils

import (
	"fmt"
	"html/template"

	"github.com/gin-gonic/gin"
)

// CsrfField csrf input html
func CsrfField(c *gin.Context, inputName string) (template.HTML, string, bool) {
	token := c.Keys[inputName]
	tokenStr, ok := token.(string)
	if !ok {
		return "", "", false
	}

	return template.HTML(fmt.Sprintf(`<input type="hidden" name="%s" value="%s">`, inputName, tokenStr)), tokenStr, true
}
