package helpers

import (
	"crypto/md5"
	"fmt"
	"gin_bbs/config"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/mozillazg/go-pinyin"
)

// SlugTranslate 百度翻译 slug
func SlugTranslate(text string) string {
	api := "http://api.fanyi.baidu.com/api/trans/vip/translate?"
	appID := config.AppConfig.BaiduTranslateAppID
	key := config.AppConfig.BaiduTranslateKey
	salt := strconv.Itoa(int(time.Now().UnixNano()))

	if appID == "" || key == "" {
		return Pinyin(text)
	}

	// md5
	sign := appID + text + salt + key
	has := md5.Sum([]byte(sign))
	sign = fmt.Sprintf("%x", has)

	// http
	query := fmt.Sprintf("from=zh&to=en&q=%s&appid=%s&salt=%s&sign=%s", url.QueryEscape(text), appID, salt, sign)
	resp, err := http.Get(api + query)
	if err != nil {
		return Pinyin(text)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Pinyin(text)
	}

	dst, err := jsonparser.GetString(body, "trans_result", "[0]", "dst")
	if err != nil {
		return Pinyin(text)
	}

	return strings.Replace(dst, " ", "-", -1)
}

// Pinyin 汉字转拼音
func Pinyin(text string) string {
	p := pinyin.LazyConvert(text, nil)
	return strings.Join(p, "-")
}
