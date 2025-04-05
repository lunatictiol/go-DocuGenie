/*
Copyright © 2025 NAME HERE bhatsabiq9@gmail.com
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/lunatictiol/go-DocuGenie/llm"
	"github.com/lunatictiol/go-DocuGenie/parser"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a professional README.md from your project files",
	Long: `DocuGenie scans your project directory, summarizes your source code files,
	and generates a polished README.md using a local LLM.
	
	Example usage:
	
	  docugenie generate
	
	This will analyze files in the current directory (excluding node_modules, .git, etc.)
	and create a README.md based on the detected languages, file structure, and code context.`,
	Run: func(cmd *cobra.Command, args []string) {
		fileTypes, _ := cmd.Flags().GetStringSlice("fileTypes")
		cwd, err := os.Getwd()
		if err != nil {
			fmt.Println("an error occured while generating the files")
		}
		fmt.Println("generate called")
		s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Fancy dots
		s.Suffix = " Generating README..."
		s.Start()
		summary := parser.Parse(cwd, fileTypes)
		llm.GenerateFile(cwd, summary)
		s.Stop()
		fmt.Println("✅ README.md generated successfully.")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringSliceP("fileTypes", "f", []string{}, "Comma-separated list of file extensions to include (e.g. .go,.js)")
}
