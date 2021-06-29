package util

import "github.com/rocketvision/beatriz/core"

func GetAttrAll(attrs []core.Attr, key string) []string {
	var result []string
	for _, attr := range attrs {
		if attr.Key == key {
			result = append(result, attr.Val)
		}
	}
	return result
}

func GetAttr(attrs []core.Attr, key string) (string, bool) {
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Val, true
		}
	}
	return "", false
}

func HasAttr(attrs []core.Attr, key string) bool {
	for _, attr := range attrs {
		if attr.Key == key && attr.Val != "" {
			return true
		}
	}
	return false
}
