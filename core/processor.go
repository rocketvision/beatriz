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
		checker.callReset(ctx)
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
			attrs := token.Attr

			for _, checker := range registry {
				checker.callEnterElement(ctx, tag, attrs)
			}

		case html.EndTagToken:
			token := tokenizer.Token()
			tag := token.Data

			for _, checker := range registry {
				checker.callLeaveElement(ctx, tag)
			}

		case html.SelfClosingTagToken:
			token := tokenizer.Token()
			tag := token.Data
			attrs := token.Attr

			for _, checker := range registry {
				checker.callEnterLeaveElement(ctx, tag, attrs)
			}

		case html.CommentToken:
			token := tokenizer.Token()
			text := token.Data

			for _, checker := range registry {
				checker.callEnterComment(ctx, text)
			}

		case html.DoctypeToken:
			token := tokenizer.Token()
			text := token.Data

			for _, checker := range registry {
				checker.callEnterDoctype(ctx, text)
			}

		case html.TextToken:
			token := tokenizer.Token()
			text := token.Data

			for _, checker := range registry {
				checker.callEnterText(ctx, text)
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
