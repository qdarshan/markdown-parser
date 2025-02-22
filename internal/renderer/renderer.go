package renderer

import (
	"fmt"
	"html"
	"strings"

	parser "github.com/qdarshan/markdown-parser/internal/parser"
	tokenizer "github.com/qdarshan/markdown-parser/internal/tokenizer"
)

// RenderHTML converts the AST to an HTML string.
func RenderHTML(ast *parser.ASTNode) string {
	if ast == nil {
		return ""
	}
	var sb strings.Builder
	for _, child := range ast.Children {
		sb.WriteString(renderNode(child))
	}
	return sb.String()
}

// renderNode converts an individual AST node into HTML.
func renderNode(node *parser.ASTNode) string {
	if node == nil {
		return ""
	}

	switch node.Type {
	case tokenizer.TOKEN_HEADING:
		return renderHeading(node)
	case tokenizer.TOKEN_PARAGRAPH:
		return renderParagraph(node)
	case tokenizer.TOKEN_LINK:
		return renderLink(node)
	case tokenizer.TOKEN_BOLD:
		return renderBold(node)
	case tokenizer.TOKEN_ITALIC:
		return renderItalic(node)
	case tokenizer.TOKEN_INLINE_CODE:
		return renderInlineCode(node)
	case tokenizer.TOKEN_TEXT:
		return renderText(node)
	default:
		return ""
	}
}

// renderHeading generates HTML for headings (H1-H6).
func renderHeading(node *parser.ASTNode) string {
	childrenHTML := renderChildren(node.Children)
	return fmt.Sprintf("<h%d>%s</h%d>", node.Level, childrenHTML, node.Level)
}

// renderParagraph generates HTML for paragraphs.
func renderParagraph(node *parser.ASTNode) string {
	childrenHTML := renderChildren(node.Children)
	return fmt.Sprintf("<p>%s</p>", childrenHTML)
}

// renderBold generates HTML for bold text.
func renderBold(node *parser.ASTNode) string {
	childrenHTML := renderChildren(node.Children)
	return fmt.Sprintf("<strong>%s</strong>", childrenHTML)
}

// renderItalic generates HTML for italic text.
func renderItalic(node *parser.ASTNode) string {
	childrenHTML := renderChildren(node.Children)
	return fmt.Sprintf("<em>%s</em>", childrenHTML)
}

// renderInlineCode generates HTML for inline code.
func renderInlineCode(node *parser.ASTNode) string {
	return fmt.Sprintf("<code>%s</code>", html.EscapeString(node.Value))
}

// renderLink generates HTML for links.
func renderLink(node *parser.ASTNode) string {
	childrenHTML := renderChildren(node.Children)
	return fmt.Sprintf("<a href=\"%s\">%s</a>", html.EscapeString(node.URL), childrenHTML)
}

// renderText escapes text content.
func renderText(node *parser.ASTNode) string {
	return html.EscapeString(node.Value)
}

// renderChildren renders all child nodes of a given node.
func renderChildren(nodes []*parser.ASTNode) string {
	if len(nodes) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, node := range nodes {
		sb.WriteString(renderNode(node))
	}
	return sb.String()
}
