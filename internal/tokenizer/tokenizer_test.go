package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenizer(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "# Heading 1",
			expected: []Token{
				{
					Type:  TOKEN_HEADING,
					Value: "Heading 1",
					Level: 1,
					Children: []Token{
						{
							Type:  TOKEN_TEXT,
							Value: "Heading 1",
						},
					},
				},
			},
		},
		{
			input: "## [text](https://example.com)",
			expected: []Token{
				{
					Type:  TOKEN_HEADING,
					Value: "text",
					Level: 2,
					Children: []Token{
						{
							Type:  TOKEN_LINK,
							Value: "text",
							URL:   "https://example.com",
							Children: []Token{
								{
									Type:  TOKEN_TEXT,
									Value: "text",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		result := Tokenize(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Input: %q, Expected: %+v, Got: %+v", test.input, test.expected, result)
		}
	}
}
