package parser

type ProjectSummary struct {
	ProjectName string        `json:"projectName"`
	Description string        `json:"description"`
	MainLang    string        `json:"mainLang"`
	Files       []FileSummary `json:"files"`
}

type FileSummary struct {
	Path           string `json:"path"`
	FileType       string `json:"fileType"`
	ContentSummary string `json:"contentSummary"`
}
