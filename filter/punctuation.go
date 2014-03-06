package filter

import (
	"strings"

	"github.com/darkhelmet/nltk"
)

func Punctuation(in nltk.TokenChan) nltk.TokenChan {
	return start(in, func(tok nltk.Token, out nltk.TokenChan) {
		cleaned := strings.Map(func(r rune) rune {
			switch {
			case 48 <= r && r <= 57: // numbers
				fallthrough
			case 65 <= r && r <= 90: // uppercase
				fallthrough
			case 97 <= r && r <= 122: // lowercase
				return r
			}
			return -1
		}, tok.String())
		out <- nltk.Token(cleaned)
	})
}
