package core

import (
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

type Processor struct{}

func NewProcessor() *Processor {
	return &Processor{}
}

func (p *Processor) ProcessReader(r io.Reader) ([]Issue, error) {
	state := NewState()
	ctx := NewContext(state)

	for _, checker := range registry {
		checker.Reset(ctx)
	}

	tokenizer := html.NewTokenizer(r)
	for {
		token := tokenizer.Next()
		raw := string(tokenizer.Raw())

		switch token {

		case html.ErrorToken:
			err := tokenizer.Err()
			if err == io.EOF {
				return state.issues, nil
			}
			return nil, err

		case html.StartTagToken:
			token := tokenizer.Token()
			tag := token.Data

			var attrs []Attr
			for _, attr := range token.Attr {
				attrs = append(attrs, Attr{
					Key: attr.Key,
					Val: attr.Val,
				})
			}

			for _, checker := range registry {
				checker.EnterElement(ctx, tag, attrs, false)
			}

		case html.EndTagToken:
			token := tokenizer.Token()
			tag := token.Data

			for _, checker := range registry {
				checker.LeaveElement(ctx, tag)
			}

		}

		lines := strings.Split(raw, "\n")
		ctx.pos.Line += len(lines) - 1
		ctx.pos.Column += len(lines[len(lines)-1])
	}
}

func (p *Processor) Process(path string) ([]Issue, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return p.ProcessReader(file)
}
