package filter

import (
	"github.com/darkhelmet/nltk"
)

var ChanSize = 10

func start(in nltk.TokenChan, f func(nltk.Token, nltk.TokenChan)) nltk.TokenChan {
	out := make(nltk.TokenChan, ChanSize)
	go func() {
		for tok := range in {
			f(tok, out)
		}
		close(out)
	}()
	return out
}
