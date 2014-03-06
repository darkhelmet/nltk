package filter

import (
	"strings"

	"github.com/darkhelmet/nltk"
)

func Superstrip(in nltk.TokenChan) nltk.TokenChan {
	return start(in, func(tok nltk.Token, out nltk.TokenChan) {
		cleaned := strings.Map(func(r rune) rune {
			switch {
			case 48 <= r && r <= 57: // numbers
				fallthrough
			case 97 <= r && r <= 122: // lowercase
				return r
			case 65 <= r && r <= 90: // uppercase
				return r + 32 // Make lowercase
			}
			return -1
		}, tok.String())
		if cleaned != "" {
			out <- nltk.Token(cleaned)
		}
	})
}
