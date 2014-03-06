package filter

import (
	"github.com/darkhelmet/nltk"
	"github.com/darkhelmet/nltk/stemmer"
)

func stemmerFilter(in nltk.TokenChan, s stemmer.Stemmer) nltk.TokenChan {
	return start(in, func(tok nltk.Token, out nltk.TokenChan) {
		out <- nltk.Token(s.Stem(tok.String()))
	})
}

func PorterStemmer(in nltk.TokenChan) nltk.TokenChan {
	s := stemmer.NewPorterStemmer()
	return stemmerFilter(in, s)
}

func SnowballStemmer(in nltk.TokenChan) nltk.TokenChan {
	s, err := stemmer.NewSnowball()
	if err != nil {
		// Ugh
		panic(err)
	}
	return stemmerFilter(in, s)
}
