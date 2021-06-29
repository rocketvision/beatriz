package main

import (
	"flag"
	"strings"

	"github.com/rocketvision/beatriz/core"
)

var (
	rawIgnoreCodes  string
	rawIncludeCodes string
)

var (
	IgnoreTypes  map[core.IssueType]bool
	IgnoreCodes  []string
	IncludeCodes []string

	FullFormatting bool
)

func FilterIssue(issue *core.Issue) bool {
	for _, filter := range IgnoreCodes {
		if matchCode(issue.Code, filter) {
			return false
		}
	}

	if IncludeCodes != nil {
		for _, filter := range IncludeCodes {
			if matchCode(issue.Code, filter) {
				return true
			}
		}
		return false
	}

	return true
}

func ParseFlags() []string {
	flag.Parse()

	if rawIgnoreCodes != "" {
		IgnoreCodes = strings.Split(rawIgnoreCodes, ",")
	}
	if rawIncludeCodes != "" {
		IncludeCodes = strings.Split(rawIncludeCodes, ",")
	}

	return flag.Args()
}

func init() {
	// TODO
	flag.StringVar(&rawIgnoreCodes, "ignore-code", "", "Códigos de erro ignorados")
	flag.StringVar(&rawIgnoreCodes, "C", "", "Códigos de erro ignorados")

	// TODO
	flag.StringVar(&rawIncludeCodes, "include-code", "", "Códigos de erro incluídos")
	flag.StringVar(&rawIncludeCodes, "c", "", "Códigos de erro incluídos")

	// TODO
	flag.BoolVar(&FullFormatting, "print-full", false, "Mostra o código de erro")
	flag.BoolVar(&FullFormatting, "f", false, "Mostra o código de erro")
}

func matchCode(actual string, expected string) bool {
	if actual == expected {
		return true
	}
	firstDigit := strings.IndexAny(actual, "1234567890")
	class := actual[:firstDigit]
	if class == expected {
		return true
	}
	return false
}
