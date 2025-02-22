package parser

import "github.com/qdarshan/markdown-parser/internal/tokenizer"

type ASTNode struct {
	Type     string
	Value    string
	Level    int
	URL      string
	Children []*ASTNode
}

func BuildAST(tokens []tokenizer.Token) *ASTNode {
	root := &ASTNode{
		Type:     "document",
		Children: []*ASTNode{},
	}

	for _, token := range tokens {
		node := convertTokenToASTNode(token)
		root.Children = append(root.Children, node)
	}
	return root
}

func convertTokenToASTNode(token tokenizer.Token) *ASTNode {
	node := &ASTNode{
		Type:  token.Type,
		Value: token.Value,
		Level: token.Level,
		URL:   token.URL,
	}

	for _, child := range token.Children {
		node.Children = append(node.Children, convertTokenToASTNode(child))
	}

	return node
}
