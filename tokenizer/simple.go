package tokenizer

import (
	"strings"

	"github.com/darkhelmet/nltk"
)

func Simple(strs ...string) nltk.TokenChan {
	tc := make(nltk.TokenChan, 10)
	go func() {
		for _, str := range strs {
			for _, t := range strings.Fields(str) {
				tc <- nltk.Token(t)
			}
		}
		close(tc)
	}()
	return tc
}
