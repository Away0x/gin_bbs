package pagination

import (
	"errors"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	// PageQueryName query 中 page 参数命名
	PageQueryName = "page"
	// PageLineQueryName query 中 pageline 参数命名
	PageLineQueryName = "pageline"
	// DefaultPageDataKeyName render data key name
	DefaultPageDataKeyName = "pagination"
)

type paginationRenderData struct {
	URL          string // 分页的 root url
	CurrentPage  int    // 当前页码
	OnFirstPage  bool   // 是否在第一页
	HasMorePages bool   // 是否有更多页
	Elements     []int  // 页码

	PreviousButtonText string // 前一页按钮文本
	PreviousPageIndex  int    // 前一页按钮的页码
	NextButtonText     string // 后一页按钮文本
	NextPageIndex      int    // 后一页按钮的页码
}

// ReturnDataFunc  返回列表所需的数据
type ReturnDataFunc = func(int, int, int, int) (interface{}, error)

// RenderPageFunc 返回分页数据已经其他需填充到模板的数据
type RenderPageFunc = func(map[string]interface{}) map[string]interface{}

// CreatePage 创建分页 (会从 query 里获取参数)
// defaultPageLine: 默认 pageline
// dataKeyName: 列表数据在 render data 中的 key name
// getTotalFunc: 获取分页数据所需的 total 数 (数据总 count)
// cb: 返回 render data key 和列表数据
func CreatePage(c *gin.Context, defaultPageLine int, dataKeyName string,
	getTotalFunc func() (int, error), returnDataFunc ReturnDataFunc) (RenderPageFunc, error) {

	// 获取数据总数
	totalCount, err := getTotalFunc()
	if err != nil {
		return nil, err
	}

	// 从 request query 中获取分页参数
	offset, limit, currentPage, totalPage := GetPageQuery(c, defaultPageLine, totalCount)
	if currentPage > totalPage {
		return nil, errors.New("没有那么多数据")
	}

	// 得到列表数据
	items, err := returnDataFunc(offset, limit, currentPage, totalPage)
	if err != nil {
		return nil, err
	}

	// 得到渲染分页的一些信息
	pageData := createPageData(c, currentPage, totalPage)

	// 返回 render 函数
	return func(otherData map[string]interface{}) map[string]interface{} {
		otherData[dataKeyName] = items
		otherData[DefaultPageDataKeyName] = pageData
		return otherData
	}, nil
}

// ------------------------- private
// 生成分页数据
func createPageData(c *gin.Context, currentPage, totalPage int) paginationRenderData {
	queryValues := url.Values{}
	for k, v := range c.Request.URL.Query() {
		if k != PageQueryName {
			queryValues.Add(k, v[0])
		}
	}
	query := queryValues.Encode()
	if query != "" {
		query = query + "&"
	}

	return paginationRenderData{
		URL:          c.Request.URL.Path + "?" + query + PageQueryName + "=",
		CurrentPage:  currentPage,
		OnFirstPage:  currentPage == 1,
		HasMorePages: currentPage != totalPage,
		Elements:     countStartAndEndPageIndex(currentPage, totalPage, 3),

		PreviousButtonText: "前一页",
		PreviousPageIndex:  currentPage - 1,
		NextButtonText:     "下一页",
		NextPageIndex:      currentPage + 1,
	}
}

// 返回一个区间数组，供生成区间页码按钮
// baseOnCurrentPageButtonOffset: 前后有多少个按钮
func countStartAndEndPageIndex(currentPage, totalPage, baseOnCurrentPageButtonOffset int) []int {
	howMuchPageButtons := baseOnCurrentPageButtonOffset*2 + 1
	startPage := 1
	endPage := 1
	result := make([]int, 0)

	if currentPage > baseOnCurrentPageButtonOffset {
		// 当前页码大于偏移量，则起始按钮为 当前页码 - 偏移量

		startPage = currentPage - baseOnCurrentPageButtonOffset
		if totalPage > (currentPage + baseOnCurrentPageButtonOffset) {
			endPage = currentPage + baseOnCurrentPageButtonOffset
		} else {
			endPage = totalPage
		}
	} else {
		// 当前页码小于偏移量

		startPage = 1
		if totalPage > howMuchPageButtons {
			endPage = howMuchPageButtons
		} else {
			endPage = totalPage
		}
	}

	if (currentPage + baseOnCurrentPageButtonOffset) > totalPage {
		startPage = startPage - (currentPage + baseOnCurrentPageButtonOffset - endPage)
	}

	if startPage <= 0 {
		startPage = 1
	}

	for i := startPage; i <= endPage; i++ {
		result = append(result, i)
	}

	return result
}
