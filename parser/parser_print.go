package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/php-any/origami/data"
)

// printDetailedError 打印详细的错误信息
func (p *Parser) printDetailedError(err string, from data.From) {
	_, _ = fmt.Fprintln(os.Stderr, "\n"+strings.Repeat("=", 80))
	_, _ = fmt.Fprintln(os.Stderr, "🚨 解析错误")
	_, _ = fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))

	if from == nil {
		_, _ = fmt.Fprintf(os.Stderr, "📍文件位置信息为空\n")
		// 显示错误消息
		_, _ = fmt.Fprintf(os.Stderr, "❌ 错误: %s\n", err)
		return
	}

	// 错误位置信息
	start, end := from.GetPosition()
	sl, sp := from.GetStartPosition()
	_, _ = fmt.Fprintf(os.Stderr, "📍 位置: 第 %d 行, 第 %d 列 (位置: %d-%d)\n", sl+1, sp+1, start, end)
	_, _ = fmt.Fprintf(os.Stderr, "📄 文件: %s:%d:%d\n", from.GetSource(), sl+1, sp+1)

	// 当前 token 信息
	currentToken := p.current()
	_, _ = fmt.Fprintf(os.Stderr, "🔍 当前 Token: %s (类型: %d)\n", currentToken.Literal, currentToken.Type)

	// 显示错误消息
	_, _ = fmt.Fprintf(os.Stderr, "❌ 错误: %s\n", err)

	// 显示上下文（前后几个 token）
	_, _ = fmt.Fprintln(os.Stderr, "\n📝 上下文:")
	p.printContext()

	_, _ = fmt.Fprintln(os.Stderr, strings.Repeat("=", 80))
}

// printContext 打印当前解析位置的上下文
func (p *Parser) printContext() {
	// 保存当前位置
	originalPos := p.position

	// 显示前3个token
	_, _ = fmt.Fprint(os.Stderr, "   前文: ")
	for i := 3; i > 0; i-- {
		if p.position-i >= 0 {
			token := p.tokens[p.position-i]
			_, _ = fmt.Fprintf(os.Stderr, "%s ", token.Literal)
		}
	}

	// 显示当前token（高亮）
	_, _ = fmt.Fprintf(os.Stderr, "\n   👉 当前: [%s] ", p.current().Literal)

	// 显示后3个token
	_, _ = fmt.Fprint(os.Stderr, "\n   后文: ")
	for i := 1; i <= 3; i++ {
		if p.position+i < len(p.tokens) {
			token := p.tokens[p.position+i]
			_, _ = fmt.Fprintf(os.Stderr, "%s ", token.Literal)
		}
	}
	_, _ = fmt.Fprintln(os.Stderr)

	// 恢复位置
	p.position = originalPos
}
