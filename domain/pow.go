package domain

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"github.com/pkg/errors"
)

type ProofOfWorker interface {
	Create() []byte
	Challenge(body []byte) ([]byte, error)
	Verify(prefix, nonce []byte) bool
}

type Algoritm uint8

const (
	SHA1 Algoritm = iota
	SHA256
	MD5
)



func (a Algoritm) Hash(d []byte) []byte {
	r := make([]byte, 0, 32)
	switch a {
	case SHA256:
		for _, b := range sha256.Sum256(d) {
			r = append(r, b)
		}
		return r
	case MD5:
		for _, b := range md5.Sum(d) {
			r = append(r, b)
		}
		return r
	case SHA1:
		for _, b := range sha1.Sum(d) {
			r = append(r, b)
		}
		return r
	default:
		panic("unkown algoritm")
	}
}


var ErrMaxAttemptsReached = errors.New("max attempts reached")
