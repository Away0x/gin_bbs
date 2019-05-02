package pongo2

import (
	"encoding/json"
	"gin_bbs/pkg/ginutils"
	"os"

	"github.com/flosch/pongo2"
)

var (
	// 存储 mix-manifest.json 解析出来的 path map
	manifests = make(map[string]string)
)

type tagMixTag struct {
	path string
}

func (node *tagMixTag) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	staticFilePath := node.path
	result := manifests[staticFilePath]

	if result == "" {
		filename := ginutils.GetGinUtilsConfig().MixFilePath
		file, err := os.Open(filename)
		if err != nil {
			writer.WriteString(ginutils.StaticPath(staticFilePath))
			return nil
		}
		defer file.Close()

		dec := json.NewDecoder(file)
		if err := dec.Decode(&manifests); err != nil {
			writer.WriteString(ginutils.StaticPath(staticFilePath))
			return nil
		}

		if string(staticFilePath[0]) == "/" {
			result = manifests[staticFilePath]
		} else {
			result = manifests["/"+staticFilePath]
		}
	}

	if result == "" {
		writer.WriteString(ginutils.StaticPath(staticFilePath))
		return nil
	}

	writer.WriteString(ginutils.StaticPath(result))
	return nil
}

// MixTag 根据 laravel-mix 生成静态文件 path
func MixTag(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	pathToken := arguments.MatchType(pongo2.TokenString)
	if pathToken == nil {
		return nil, arguments.Error("mix tag error: path 必须为 string.", nil)
	}

	nowNode := &tagMixTag{
		path: pathToken.Val,
	}

	if arguments.Remaining() > 0 {
		return nil, arguments.Error("Malformed static-tag arguments.", nil)
	}

	return nowNode, nil
}
