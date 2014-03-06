package filter

import (
	"strings"

	"github.com/darkhelmet/nltk"
)

func Ascii(in nltk.TokenChan) nltk.TokenChan {
	return start(in, func(tok nltk.Token, out nltk.TokenChan) {
		cleaned := strings.Map(func(r rune) rune {
			if r > 127 {
				return -1
			}
			return r
		}, tok.String())
		out <- nltk.Token(cleaned)
	})
}
