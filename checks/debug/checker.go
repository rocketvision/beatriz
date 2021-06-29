package debug

import (
	"log"

	"github.com/rocketvision/beatriz/core"
	"golang.org/x/net/html"
)

type Checker struct{}

func (Checker) Reset(ctx *core.Context) {
	log.Println("Reset")
}

func (Checker) EnterElement(ctx *core.Context, tag string, attrs []html.Attribute, closed bool) {
	log.Println("Enter:", tag)
	for _, attr := range attrs {
		log.Println(" Attr:", attr.Key, "=>", attr.Val)
	}
}

func (Checker) LeaveElement(ctx *core.Context, tag string) {
	log.Println("Leave:", tag)
}

func init() {
	core.RegisterChecker(Checker{})
}
