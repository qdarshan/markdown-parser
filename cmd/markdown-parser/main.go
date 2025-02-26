package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/qdarshan/markdown-parser/internal/parser"
	"github.com/qdarshan/markdown-parser/internal/renderer"
	"github.com/qdarshan/markdown-parser/internal/tokenizer"
)

// PageData represents the data structure for the HTML template.
type PageData struct {
	Content template.HTML
}

// parseMarkdownHandler serves the HTML template for the Markdown editor.
func parseMarkdownHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	// Render the template with empty content initially
	pageData := PageData{Content: ""}
	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

// renderMarkdownHandler handles Markdown input from the user and returns the rendered HTML.
func renderMarkdownHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Read Markdown input from the request body
	markdownInput := r.FormValue("markdown")

	// Tokenize -> Parse -> Render
	tokens := tokenizer.Tokenize(markdownInput)
	ast := parser.BuildAST(tokens)
	renderedHTML := renderer.RenderHTML(ast)

	// Return the rendered HTML as the response
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(renderedHTML))
}

func main() {
	// Serve the Markdown editor page
	http.HandleFunc("/", parseMarkdownHandler)

	// Handle Markdown rendering requests
	http.HandleFunc("/render", renderMarkdownHandler)

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
