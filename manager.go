package main

import (
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/rocketvision/beatriz/core"
)

var formats = []string{
	".html",
	".htm",
}

var except = []string{
	".ttf",
	".woff",
}

type Manager struct {
	temporary []*os.File
	processor *core.Processor
	summary   []summaryItem
	total     int
}

type summaryItem struct {
	path  string
	total int
}

func NewManager() *Manager {
	return &Manager{
		processor: core.NewProcessor(),
	}
}

func (m *Manager) Process(path string) error {
	url, err := url.Parse(path)
	if err == nil && url.Scheme != "" {
		return m.processRemote(url)
	}

	if info, err := os.Stat(path); err == nil && info.IsDir() {
		return m.processTree(path)
	}

	return m.processFile(path, path)
}

func (m *Manager) processRemote(url *url.URL) error {
	log.Println("Fazendo download:", url)

	resp, err := http.Get(url.String())
	if err != nil {
		return fmt.Errorf("falha de download: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("falha de download: status %s", resp.Status)
	}

	name := filepath.Base(url.Host) + "*.html"
	temp, err := m.getTemporaryFile(name)
	if err != nil {
		return fmt.Errorf("falha ao reservar arquivo temporário: %w", err)
	}

	if _, err := io.Copy(temp, resp.Body); err != nil {
		return fmt.Errorf("falha ao copiar arquivo do servidor: %w", err)
	}

	return m.processFile(temp.Name(), url.String())
}

func (m *Manager) processTree(root string) error {
	return filepath.WalkDir(root,
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
			return m.processFile(path, rel)
		})
}

func (m *Manager) processFile(path, name string) error {
	log.Println("Processando:", name)

	issues, err := m.processor.Process(path)
	if err != nil {
		return err
	}

	total := 0
	for _, issue := range issues {
		if FilterIssue(&issue) {
			if FullFormatting {
				log.Printf("  Linha %4d | %s [%s]", issue.Pos.Line, issue.Text, issue.Code)
			} else {
				log.Printf("  (%d) %s", issue.Pos.Line, issue.Text)
			}
			total++
		}
	}
	log.Println("Total:", total)
	log.Println()

	m.summary = append(m.summary,
		summaryItem{
			path:  name,
			total: total,
		})
	m.total += total
	return nil
}

func (m *Manager) PrintSummary() {
	sort.Slice(m.summary, func(i, j int) bool {
		return m.summary[i].total > m.summary[j].total
	})

	log.Println("Sumário")
	for _, item := range m.summary {
		if FullFormatting {
			log.Printf("  Erros %4d | %s", item.total, item.path)
		} else {
			log.Printf("  (%d) %s", item.total, item.path)
		}
	}
	log.Println("Total global:", m.total)
}

func (m *Manager) DeleteTemporaries() {
	// if m.temporary != nil {
	// log.Print("Deletando arquivos temporários...")
	for _, temp := range m.temporary {
		temp.Close()
		os.Remove(temp.Name())
	}
	// }
}

func (m *Manager) getTemporaryFile(name string) (*os.File, error) {
	file, err := os.CreateTemp(".", name)
	if err != nil {
		return nil, err
	}

	m.temporary = append(m.temporary, file)
	return file, nil
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
