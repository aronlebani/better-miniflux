package utils

import (
    "strings"
    "golang.org/x/net/html"
)

func GetImgTag(htmlString string) string {
	r := strings.NewReader(htmlString)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return ""
		case html.SelfClosingTagToken, html.StartTagToken:
			tn, _ := z.TagName()
			if string(tn) == "img" {
				raw := z.Raw()
				return string(raw)
			}
		}
	}
}
