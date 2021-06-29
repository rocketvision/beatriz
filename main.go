package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rocketvision/beatriz/core"

	_ "github.com/rocketvision/beatriz/checks/a"
	_ "github.com/rocketvision/beatriz/checks/attrs"
	_ "github.com/rocketvision/beatriz/checks/basic"
	_ "github.com/rocketvision/beatriz/checks/img"
	_ "github.com/rocketvision/beatriz/checks/inline"
	_ "github.com/rocketvision/beatriz/checks/stack"
)

func main() {
	log.SetFlags(0)

	args := ParseFlags()
	if len(args) < 1 {
		log.Printf("Uso: %v <arquivo ou diretório>", os.Args[0])
		os.Exit(1)
	}
	// TODO: Single file.
	// TODO: Merge results?
	processTree(args[0])
}

var formats = []string{
	".html",
	".htm",
}

var except = []string{
	".ttf",
	".woff",
}

func acceptFormat(path string) bool {
	for _, format := range formats {
		if strings.HasSuffix(path, format) {
			return true
		}
	}
	return false
}

func rejectFormat(path string) bool {
	for _, invalid := range except {
		if strings.Contains(path, invalid) {
			return true
		}
	}
	return false
}

func processTree(root string) {
	processor := core.NewProcessor()
	total := 0

	err := filepath.WalkDir(root,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.IsDir() {
				return nil
			}
			if !acceptFormat(path) || rejectFormat(path) {
				return nil
			}

			rel, _ := filepath.Rel(root, path)
			log.Println("Processando:", rel)
			issues, err := processor.Process(path)
			if err != nil {
				return err
			}

			count := 0
			for _, issue := range issues {
				if FilterIssue(&issue) {
					if FullFormatting {
						log.Printf("  Linha %4v | %v [%v]", issue.Pos.Line, issue.Text, issue.Code)
					} else {
						log.Printf("  (%v) %v", issue.Pos.Line, issue.Text)
					}
					count++
				}
			}
			log.Println("Total:", count)
			log.Println()

			total += count
			return nil
		})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Total global:", total)
}
