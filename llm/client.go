package llm

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lunatictiol/go-DocuGenie/parser"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

func GenerateFile(wrd string, summary parser.ProjectSummary) {

	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		log.Fatal(err)
	}

	prompt := `
You're a helpful AI assistant tasked with generating a professional and developer-friendly README.md for a software project.

Use the following structured project summary to generate the README content:

---
Project Name: {{.ProjectName}}
Main Language: {{.MainLang}}
Description: {{.Description}}

Files and Key Snippets:
{{range .Files}}
- {{.Path}} ({{.FileType}})
  {{.ContentSummary | printf "%s"}}
{{end}}
---

Instructions:
1. Begin with a clear project title and description.
2. Include a "Features" or "Tech Stack" section based on the main language and file types.
3. Include a "Getting Started" section with placeholder steps.
4. Include a "Folder Structure" or "Key Files" section based on the paths.
5. Format everything using proper Markdown syntax.
6. Keep the tone professional and clear.

Return only the markdown content. Do not wrap it in code blocks or provide any explanations.

Keep formatting simple and avoid unnecessary characters — this will be written directly to a file using Go.
`

	data := "Peoject summary : " + fmt.Sprintf("%v", summary)
	query := prompt + " " + data
	ctx := context.Background()
	completion, err := llms.GenerateFromSinglePrompt(ctx, llm, query)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(wrd + "/README.md")
	if err != nil {
		log.Fatal("error creating the file")
	}
	defer f.Close()

	n, err := f.WriteString(completion)
	if err != nil {
		log.Fatal("error writting the file")
	}

	fmt.Printf("%d number of bytes written\n", n)

	f.Sync()

}
