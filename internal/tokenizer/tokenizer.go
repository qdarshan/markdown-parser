package tokenizer

import (
	"regexp"
	"strings"
)

const (
	TOKEN_HEADING     = "heading"
	TOKEN_TEXT        = "text"
	TOKEN_LIST_ITEM   = "list_item"
	TOKEN_CODE_BLOCK  = "code_block"
	TOKEN_LINK        = "link"
	TOKEN_BOLD        = "bold"
	TOKEN_ITALIC      = "italic"
	TOKEN_INLINE_CODE = "inline_code"
	TOKEN_PARAGRAPH   = "paragraph"
	TOKEN_ERROR       = "error"
)

type Token struct {
	Type     string
	Value    string
	URL      string
	Level    int
	Children []Token
}

// Tokenize processes an entire Markdown document and returns a list of tokens.
func Tokenize(content string) []Token {
	lines := strings.Split(content, "\n")
	tokens := make([]Token, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue // Skip empty lines
		}
		tokens = append(tokens, tokenizeLine(line))
	}

	return tokens
}

// tokenizeLine processes block-level elements like headings, lists, paragraphs, etc.
func tokenizeLine(line string) Token {
	if strings.HasPrefix(line, "#") {
		return tokenizeHeading(line)
	}
	return tokenizeParagraph(line)
}

// tokenizeHeading processes heading lines (e.g., "# Heading")
func tokenizeHeading(text string) Token {
	level := 0
	for level < len(text) && text[level] == '#' {
		level++
	}

	level = min(level, 6) // Markdown supports up to H6
	headingText := strings.TrimSpace(text[level:])

	// Tokenize inline elements inside the heading
	children := tokenizeInline(headingText)

	plainText := ""
	for _, child := range children {
		plainText += child.Value + " "
	}
	plainText = strings.TrimSpace(plainText)
	return Token{
		Type:     TOKEN_HEADING,
		Value:    plainText,
		Level:    level,
		Children: children,
	}
}

// tokenizeParagraph processes normal text as a paragraph
func tokenizeParagraph(text string) Token {
	children := tokenizeInline(text)

	return Token{
		Type:     TOKEN_PARAGRAPH,
		Value:    text,
		Children: children,
	}
}

// Regular expressions for inline elements
var (
	linkRE       = regexp.MustCompile(`^\[([^\]]+)\]\(([^)]+)\)`)
	boldRE       = regexp.MustCompile(`^\*\*(.+?)\*\*`)
	italicRE     = regexp.MustCompile(`^\*(.+?)\*`)
	inlineCodeRE = regexp.MustCompile("^`([^`]+)`")
)

// tokenizeInline processes inline elements (bold, italic, links, inline code)
func tokenizeInline(text string) []Token {
	tokens := []Token{}
	text = strings.TrimSpace(text)

	for text != "" {
		switch {
		case linkRE.MatchString(text):
			token, match := tokenizeLink(text)
			tokens = append(tokens, token)
			text = text[len(match):]

		case boldRE.MatchString(text):
			token, match := tokenizeBold(text)
			tokens = append(tokens, token)
			text = text[len(match):]

		case italicRE.MatchString(text):
			token, match := tokenizeItalic(text)
			tokens = append(tokens, token)
			text = text[len(match):]

		case inlineCodeRE.MatchString(text):
			token, match := tokenizeInlineCode(text)
			tokens = append(tokens, token)
			text = text[len(match):]

		default:
			nextIndex := strings.IndexAny(text, "[*`")
			if nextIndex == -1 {
				tokens = append(tokens, Token{Type: TOKEN_TEXT, Value: text})
				return tokens
			}
			if nextIndex > 0 {
				tokens = append(tokens, Token{Type: TOKEN_TEXT, Value: text[:nextIndex]})
			}
			text = text[nextIndex:]
		}
	}
	return tokens
}

// tokenizeLink extracts links from text
func tokenizeLink(text string) (Token, string) {
	matches := linkRE.FindStringSubmatch(text)
	if len(matches) != 3 {
		return Token{Type: TOKEN_ERROR, Value: "Invalid link syntax"}, ""
	}
	children := tokenizeInline(matches[1]) // Process inline elements inside link text
	return Token{
		Type:     TOKEN_LINK,
		Value:    matches[1],
		URL:      matches[2],
		Children: children,
	}, matches[0]
}

// tokenizeBold extracts bold text
func tokenizeBold(text string) (Token, string) {
	matches := boldRE.FindStringSubmatch(text)
	if len(matches) != 2 {
		return Token{Type: TOKEN_ERROR, Value: "Invalid bold syntax"}, ""
	}
	children := tokenizeInline(matches[1]) // Process inline elements inside bold text
	return Token{
		Type:     TOKEN_BOLD,
		Value:    matches[1],
		Children: children,
	}, matches[0]
}

// tokenizeItalic extracts italic text
func tokenizeItalic(text string) (Token, string) {
	matches := italicRE.FindStringSubmatch(text)
	if len(matches) != 2 {
		return Token{Type: TOKEN_ERROR, Value: "Invalid italic syntax"}, ""
	}
	children := tokenizeInline(matches[1])
	return Token{
		Type:     TOKEN_ITALIC,
		Value:    matches[1],
		Children: children,
	}, matches[0]
}

// tokenizeInlineCode extracts inline code snippets
func tokenizeInlineCode(text string) (Token, string) {
	matches := inlineCodeRE.FindStringSubmatch(text)
	if len(matches) != 2 {
		return Token{Type: TOKEN_ERROR, Value: "Invalid inline code syntax"}, ""
	}
	return Token{
		Type:  TOKEN_INLINE_CODE,
		Value: matches[1],
	}, matches[0]
}

// Helper function to get the minimum of two numbers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
