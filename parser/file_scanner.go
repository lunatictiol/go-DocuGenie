package parser

import (
	"io/fs"
	"path/filepath"
	"slices"
	"strings"
)

func listFiles(dir string, filetypes []string) []string {
	skipDirs := map[string]bool{
		".git":         true,
		"node_modules": true,
		"vendor":       true,
		"build":        true,
	}
	var files []string
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			dirName := d.Name()
			if skipDirs[dirName] || strings.HasPrefix(dirName, ".") {
				return filepath.SkipDir
			}
			return nil
		}

		// Check file extension
		if contains(filetypes, filepath.Ext(d.Name())) {
			files = append(files, path)
		}

		return nil
	})

	return files

}

func contains(list []string, item string) bool {
	return slices.Contains(list, item)
}
