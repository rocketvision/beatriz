package debug

import (
	"log"

	"github.com/rocketvision/beatriz/core"
	"golang.org/x/net/html"
)

func Reset(ctx *core.Context) {
	log.Println("Reset")
}

func EnterElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	if Verbose {
		log.Println("Enter:", tag)
		for _, attr := range attrs {
			log.Println(" Attr:", attr.Key, "=>", attr.Val)
		}
	}
}

func LeaveElement(ctx *core.Context, tag string) {
	if Verbose {
		log.Println("Leave:", tag)
	}
}

func EnterLeaveElement(ctx *core.Context, tag string, attrs []html.Attribute) {
	log.Println("EnterLeave:", tag)
	for _, attr := range attrs {
		log.Println(" Attr:", attr.Key, "=>", attr.Val)
	}
}

func EnterComment(ctx *core.Context, text string) {
	log.Println("Comment:", text)
}

func EnterDoctype(ctx *core.Context, text string) {
	log.Println("Doctype:", text)
}

func EnterText(ctx *core.Context, text string) {
	if Verbose {
		log.Println("Text:", text)
	}
}

func init() {
	core.RegisterChecker(&core.Checker{
		Reset:             Reset,
		EnterElement:      EnterElement,
		LeaveElement:      LeaveElement,
		EnterLeaveElement: EnterLeaveElement,
		EnterComment:      EnterComment,
		EnterDoctype:      EnterDoctype,
		EnterText:         EnterText,
	})
}
