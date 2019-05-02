package pongo2

import (
	"gin_bbs/pkg/ginutils"

	"github.com/flosch/pongo2"
)

type tagStaticTag struct {
	path string
}

func (node *tagStaticTag) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	path := ginutils.StaticPath(node.path)
	writer.WriteString(path)
	return nil
}

func StaticTag(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	pathToken := arguments.MatchType(pongo2.TokenString)
	if pathToken == nil {
		return nil, arguments.Error("static tag error: path 必须为 string.", nil)
	}

	nowNode := &tagStaticTag{
		path: pathToken.Val,
	}

	if arguments.Remaining() > 0 {
		return nil, arguments.Error("Malformed static-tag arguments.", nil)
	}

	return nowNode, nil
}
