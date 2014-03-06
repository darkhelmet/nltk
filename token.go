package nltk

type Token string

func (t Token) String() string {
	return string(t)
}

type TokenChan chan Token
