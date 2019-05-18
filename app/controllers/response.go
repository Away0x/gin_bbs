package controllers

import (
	"net/http"

	"gin_bbs/pkg/errno"
	"gin_bbs/pkg/ginutils/pagination"

	"github.com/gin-gonic/gin"
)

// Response 标准响应数据
type Response struct {
	Code    int    `json:"code"`    // 自定义的异常码
	Message string `json:"message"` // 错误描述
	// 具体错误信息
	// - 只用于调试的，前端不应使用该值，因为其可能会不存在
	Errors interface{} `json:"errors,omitempty"`
	// 具体响应数据
	// - 无数据时，不返回
	// - 如果是列表数据建议使用 ListData 类型
	// - 不需分页的列表类型时，也建议 data.list 这样响应
	Data interface{} `json:"data"`
}

// ListData 带列表时的 data
type ListData struct {
	Page     int         `json:"page,omitempty"`     // 当前页数
	PageLine int         `json:"pageline,omitempty"` // 每页条数
	Total    int         `json:"total"`              // 总数
	List     interface{} `json:"list"`               // 列表数据 (无数据时，默认返回一个 [])
	Other    interface{} `json:"other,omitempty"`    // 其他数据 (可选)
}

// SendResponse -
func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message, errors := errno.Decode(err)

	r := Response{
		Code:    code,
		Message: message,
	}

	if d, ok := data.(ListData); ok {
		if d.List == nil {
			d.List = make([]interface{}, 0)
		}
		r.Data = d
	} else {
		r.Data = data
	}

	if errors != nil {
		r.Errors = errors
	}

	c.JSON(http.StatusOK, r)
}

// SendOKResponse 成功响应
func SendOKResponse(c *gin.Context, data interface{}) {
	SendResponse(c, nil, data)
}

// SendErrorResponse 带错误的响应
func SendErrorResponse(c *gin.Context, err error) {
	SendResponse(c, err, nil)
}

// SendListResponse 分页数据
func SendListResponse(c *gin.Context, defaultPageLine int, other interface{},
	getTotalFunc func() (int, error),
	returnDataFunc func(int, int, int, int) (interface{}, error)) {

	// 获取数据总数
	totalCount, err := getTotalFunc()
	if err != nil {
		SendOKResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	// 从 request query 中获取分页参数
	offset, limit, currentPage, totalPage := pagination.GetPageQuery(c, defaultPageLine, totalCount)
	if currentPage > totalPage {
		SendOKResponse(c, errno.Base(errno.ParamsError, "没有那么多数据"))
		return
	}

	// 得到列表数据
	items, err := returnDataFunc(offset, limit, currentPage, totalPage)
	if err != nil {
		SendOKResponse(c, errno.New(errno.DatabaseError, err))
		return
	}

	listData := ListData{
		Page:     currentPage,
		PageLine: pagination.GetPageLine(c, defaultPageLine),
		Total:    totalCount,
		List:     items,
	}
	if other != nil {
		listData.Other = other
	}

	SendOKResponse(c, listData)
}
