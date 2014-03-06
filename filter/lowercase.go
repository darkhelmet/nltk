package filter

import (
	"strings"

	"github.com/darkhelmet/nltk"
)

func Lowercase(in nltk.TokenChan) nltk.TokenChan {
	return start(in, func(tok nltk.Token, out nltk.TokenChan) {
		out <- nltk.Token(strings.ToLower(tok.String()))
	})
}
