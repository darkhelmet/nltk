package stemmer

import (
	"bitbucket.org/tebeka/snowball"
)

func NewSnowball() (Stemmer, error) {
	return snowball.New("english")
}
