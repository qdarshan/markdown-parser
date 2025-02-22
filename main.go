package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

// countPounds counts the leading '#' characters in a string.
func countPounds(line string) int {
	count := 0
	for _, char := range line {
		if char == '#' {
			count++
		} else {
			break
		}
	}
	return count
}

func parseHeading()  {
	
}

// parseMarkdown converts Markdown content into basic HTML.
func parseMarkdown(content string) string {
	lines := strings.Split(content, "\n")
	var body strings.Builder

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "#") {
			count := countPounds(line)
			text := strings.TrimSpace(line[count:])
			body.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n", count, text, count))
		} else {
			body.WriteString(fmt.Sprintf("<p>%s</p>\n", line))
		}
	}

	return body.String()
}

// PageData represents the data structure for the HTML template.
type PageData struct {
	Content template.HTML
}

// parseMarkdownHandler serves the parsed Markdown as an HTML page.
func parseMarkdownHandler(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("text.md")
	if err != nil {
		http.Error(w, "Could not open file", http.StatusInternalServerError)
		log.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	content, err := os.ReadFile("text.md")
	if err != nil {
		http.Error(w, "Error reading file", http.StatusInternalServerError)
		log.Println("Error reading file:", err)
		return
	}

	parsedMarkdown := parseMarkdown(string(content))
	pageData := PageData{Content: template.HTML(parsedMarkdown)}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		http.Error(w, "Template rendering error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}

	if err := tmpl.Execute(w, pageData); err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

func main() {
	http.HandleFunc("/", parseMarkdownHandler)
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
