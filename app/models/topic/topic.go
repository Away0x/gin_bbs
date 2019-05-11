package topic

import (
	"gin_bbs/app/helpers"
	"gin_bbs/app/models"
	"gin_bbs/database"
	"gin_bbs/pkg/ginutils/utils"
	"regexp"
	"strings"
)

// Topic 话题
type Topic struct {
	models.BaseModel
	Title           string `gorm:"column:title;type:varchar(255);not null" sql:"index"` // 帖子标题
	Body            string `gorm:"column:body;type:text;not null"`                      // 帖子内容
	UserID          uint   `gorm:"column:user_id;not null" sql:"index"`                 // 用户 ID
	CategoryID      uint   `gorm:"column:category_id;not null" sql:"index"`             // 分类 ID
	ReplyCount      int    `gorm:"column:reply_count;not null;default:0"`               // 回复数量
	ViewCount       int    `gorm:"column:view_count;not null;default:0"`                // 查看总数
	LastReplyUserID uint   `gorm:"column:last_reply_user_id;not null;default:0"`        // 最后回复的用户 ID
	Order           int    `gorm:"column:order;not null;default:0"`                     // 排序
	Excerpt         string `gorm:"column:excerpt;type:text"`                            // 文章摘要，SEO 优化时使用
	Slug            string `gorm:"column:slug;type:varchar(255)"`                       // SEO 友好的 URI
}

// TableName 表名
func (Topic) TableName() string {
	return "topics"
}

// BeforeSave - hook
func (t *Topic) BeforeSave() error {
	t.Body = utils.XSSClean(t.Body)
	t.Excerpt = makeExcerpt(t.Body, 200)

	return nil
}

// AfterSave - hook
func (t *Topic) AfterSave() error {
	if t.Slug == "" {
		// SlugTranslate 需要请求百度 api，放到协程中执行
		go func(t *Topic) {
			slug := helpers.SlugTranslate(t.Title)
			database.DB.Model(&t).UpdateColumn("slug", slug) // 这样更新可避免触发 gorm 钩子，从而导致死循环
		}(t)
	}

	return nil
}

func makeExcerpt(value string, length int) string {
	r := regexp.MustCompile(`\r\n|\r|\n+|\<[\S\s]+?\>`)
	v := string(r.ReplaceAll([]byte(value), []byte("")))
	v = strings.TrimSpace(v)
	ru := []rune(v) // utf8 字符串需先转 rune 才可 [:]

	if len(ru) < length {
		return v
	}
	return string(ru[:length])
}
