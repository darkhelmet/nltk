package stemmer

type Stemmer interface {
	Stem(string) string
}
