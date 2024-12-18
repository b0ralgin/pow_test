package pow

import (
	"crypto/rand"
	"github.com/b0ralgin/pow_test/domain"
	"github.com/pkg/errors"
)

type HashCach struct {
	Size       uint8
	Difficulty uint8
	Algo       domain.Algoritm
	MaxAttempts int
}

func (h HashCach) Create() []byte {
	b := make([]byte, 0, h.Size)
	b = append(b, byte(h.Algo), h.Size, h.Difficulty)
	_, err := rand.Read(b[3:h.Size])
	if err != nil {
		panic(err) //
	}
	return b[:h.Size]
}

func (h HashCach) Challenge(prefix []byte) ([]byte, error) {
	nonceBytes := make([]byte, 8)
	plen := len(prefix) - 1
	algo := domain.Algoritm(prefix[0])
	puzzle := append(append(prefix, nonceBytes...)) // alloc memory

	for range h.MaxAttempts {
		if _, err := rand.Read(puzzle[plen:]); err != nil {
			return nil, errors.Wrap(err, "failed to create nonce ")
		}
		hash := algo.Hash(puzzle)
		if !h.checkZeroes(hash) {
			continue
		}
		return puzzle[plen:], nil
	}
	return nil, domain.ErrMaxAttemptsReached
}
func (h HashCach) checkZeroes(d []byte) bool {
	leadingBytes := h.Difficulty / 8
	if leadingBytes > 0 {
		for _, b := range d[:leadingBytes] {
			if b != 0 {
				return false
			}
		}
	}

	remainingZeros := h.Difficulty % 8
	if remainingZeros == 0 {
		return true
	}
	mask := byte(255 >> remainingZeros)
	return (d[leadingBytes] & mask) == 0
}

func (h HashCach) Verify(prefix, nonce []byte) bool {
	puzzle := append(prefix, nonce...)
	hash := h.Algo.Hash(puzzle)
	return h.checkZeroes(hash)
}
