package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// SummariseProject reads and summarizes the given file paths.
func SummariseProject(filePaths []string) ProjectSummary {
	projectName := inferProjectName()
	mainLang := guessMainLang(filePaths)
	var files []FileSummary

	for _, path := range filePaths {
		content := readFirst100Lines(path)
		ext := filepath.Ext(path)
		summary := extractSummary(content, ext)
		files = append(files, FileSummary{
			Path:           path,
			FileType:       ext,
			ContentSummary: summary,
		})
	}

	return ProjectSummary{
		ProjectName: projectName,
		Description: fmt.Sprintf("Auto-generated summary for %s", projectName),
		MainLang:    mainLang,
		Files:       files,
	}
}

// readFirst100Lines reads up to the first 100 lines of a file.
func readFirst100Lines(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		return []string{}
	}

	lines := strings.Split(string(data), "\n")
	if len(lines) > 100 {
		return lines[:100]
	}
	return lines
}

// inferProjectName grabs the current folder name.
func inferProjectName() string {
	wd, err := os.Getwd()
	if err != nil {
		return "unknown-project"
	}
	return filepath.Base(wd)
}

// guessMainLang returns the most common file extension.
func guessMainLang(paths []string) string {
	count := map[string]int{}
	for _, p := range paths {
		ext := filepath.Ext(p)
		count[ext]++
	}

	var maxExt string
	maxCount := 0
	for ext, c := range count {
		if c > maxCount {
			maxExt = ext
			maxCount = c
		}
	}

	if maxExt == "" {
		return "unknown"
	}
	return maxExt
}

// extractSummary dispatches to language-specific summarizers.
func extractSummary(lines []string, ext string) string {
	switch ext {
	case ".go":
		return extractGoSummary(lines)
	case ".js", ".jsx", ".ts", ".tsx":
		return extractJsTsSummary(lines)
	case ".py":
		return extractPythonSummary(lines)
	default:
		return "Unsupported file type."
	}
}

// Language-specific summarizers:

func extractGoSummary(lines []string) string {
	var summary []string
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "//") ||
			strings.HasPrefix(trim, "package") ||
			strings.HasPrefix(trim, "import") ||
			strings.HasPrefix(trim, "func") ||
			strings.HasPrefix(trim, "type") {
			summary = append(summary, trim)
		}
	}
	return joinOrFallback(summary)
}

func extractJsTsSummary(lines []string) string {
	var summary []string
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "//") ||
			strings.HasPrefix(trim, "/*") ||
			strings.HasPrefix(trim, "function") ||
			strings.HasPrefix(trim, "const") ||
			strings.HasPrefix(trim, "let") ||
			strings.HasPrefix(trim, "class") ||
			strings.Contains(trim, "=>") {
			summary = append(summary, trim)
		}
	}
	return joinOrFallback(summary)
}

func extractPythonSummary(lines []string) string {
	var summary []string
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "#") ||
			strings.HasPrefix(trim, "def") ||
			strings.HasPrefix(trim, "class") ||
			strings.HasPrefix(trim, "import") {
			summary = append(summary, trim)
		}
	}
	return joinOrFallback(summary)
}

// Fallback handler if no summary found.
func joinOrFallback(lines []string) string {
	if len(lines) == 0 {
		return "No significant comments or declarations found."
	}
	return strings.Join(lines, "\n")
}
