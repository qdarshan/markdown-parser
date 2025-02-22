package main

import (
	"encoding/json"
	"fmt"

	tokenizer "github.com/qdarshan/markdown-parser/internal/tokenizer"
)

func main() {
	content := ` # h1
				## h2
				### h3
				#### h4
				##### h5
				###### h6

				## [text](url)
				#### [text](https://example.com)
`
	tokens := tokenizer.Tokenize(content)
	jsonBytes, err := json.MarshalIndent(tokens, "", "  ") // Use MarshalIndent for pretty printing
	if err != nil {
		fmt.Println("Error marshaling tokens to JSON:", err)
		return
	}
	fmt.Println(string(jsonBytes))
}
