package parser

func Parse(dir string, fileType []string) ProjectSummary {
	files := listFiles(dir, fileType)
	summary := SummariseProject(files)
	return summary
}
