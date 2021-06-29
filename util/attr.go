package util

import "golang.org/x/net/html"

func GetAttrAll(attrs []html.Attribute, key string) []string {
	var result []string
	for _, attr := range attrs {
		if attr.Key == key {
			result = append(result, attr.Val)
		}
	}
	return result
}

func GetAttr(attrs []html.Attribute, key string) (string, bool) {
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func HasAttr(attrs []html.Attribute, key string) bool {
	for _, attr := range attrs {
		if attr.Key == key && attr.Val != "" {
			return true
		}
	}
	return false
}
