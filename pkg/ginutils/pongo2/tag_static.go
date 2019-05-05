package pongo2

import (
	ginfile "gin_bbs/pkg/ginutils/file"

	"github.com/flosch/pongo2"
)

type tagStaticTag struct {
	path string
	expr pongo2.IEvaluator
}

func (node *tagStaticTag) Execute(ctx *pongo2.ExecutionContext, writer pongo2.TemplateWriter) *pongo2.Error {
	path := ""

	if node.expr != nil {
		// 之前存储的是表达式
		val, err := node.expr.Evaluate(ctx)
		if err != nil {
			return err
		}

		path = ginfile.StaticPath(val.String())
	} else if node.path != "" {
		// 之前存储的是字符串
		path = ginfile.StaticPath(node.path)
	}

	writer.WriteString(path)
	return nil
}

// StaticTag 生成项目静态文件地址
func StaticTag(doc *pongo2.Parser, start *pongo2.Token, arguments *pongo2.Parser) (pongo2.INodeTag, *pongo2.Error) {
	staticNode := &tagStaticTag{path: "", expr: nil}

	pathToken := arguments.MatchType(pongo2.TokenString)
	if pathToken == nil {
		exprVal, err := arguments.ParseExpression() // 不是字符串而是表达式
		if err != nil {
			return nil, err
		}

		staticNode.expr = exprVal
		return staticNode, nil
	}

	staticNode.path = pathToken.Val
	if arguments.Remaining() > 0 {
		return nil, arguments.Error("Malformed static-tag arguments.", nil)
	}

	return staticNode, nil
}
