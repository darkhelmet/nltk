package filter

import (
	"github.com/darkhelmet/nltk"
	"strings"
)

func Lowercase(in nltk.TokenChan) nltk.TokenChan {
	return start(in, func(tok nltk.Token, out nltk.TokenChan) {
		out <- nltk.Token(strings.ToLower(tok.String()))
	})
}
