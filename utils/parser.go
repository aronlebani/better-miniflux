package utils

import (
    "strings"
    "golang.org/x/net/html"
)

func isImg(tn []byte) bool {
    return string(tn) == "img"
}

func isIframe(tn []byte) bool {
    return string(tn) == "iframe"
}

func isYoutubeEmbed(tn []byte, z *html.Tokenizer) bool {
    if !isIframe(tn) {
        return false
    }

    for {
        key, val, moreAttr := z.TagAttr()
        switch string(key) {
        case "src":
            return strings.Contains(string(val), "youtube")
        default:
            if moreAttr {
                continue
            } else {
                break
            }
        }
    }
}

func GetMediaElement(htmlString string) string {
	r := strings.NewReader(htmlString)
	z := html.NewTokenizer(r)

	for {
		tt := z.Next()
        tn, _ := z.TagName()
		raw := z.Raw()
		switch tt {
		case html.ErrorToken:
			return ""
		case html.SelfClosingTagToken, html.StartTagToken:
			if isImg(tn) {
				return string(raw)
			}
			if isYoutubeEmbed(tn, z) {
				return string(raw) + "</iframe>"
			}
		}
	}
}
