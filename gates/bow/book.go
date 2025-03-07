package bow

import "math/rand"

type SimpleBook []string

func NewSimpleBook() *SimpleBook {
	return &SimpleBook{
		"Simple qoute",
		"Another simple qoute",
		"Wise man once said",
	}
}

func (sb SimpleBook) GetQoute() (string, error) {
	l := rand.Int31n(int32(len(sb)-1))
	return sb[l], nil
}
