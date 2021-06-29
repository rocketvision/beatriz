package a

var InvalidHREF = map[string]bool{
	"#":                   true,
	"javascript:void(0)":  true,
	"javascript:void(0);": true,
}
