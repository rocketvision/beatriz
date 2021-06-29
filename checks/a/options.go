package a

var RejectedHREF = map[string]bool{
	"":                    true,
	"#":                   true,
	"javascript:void(0)":  true,
	"javascript:void(0);": true,
}
