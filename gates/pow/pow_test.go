package pow

import (
	"github.com/b0ralgin/pow_test/domain"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCheckZeroes(t *testing.T) {
	hash := domain.SHA256.Hash([]byte("Omar986924"))

	leading6h := HashCach{
		Size:       8,
		Difficulty: 22,
		Algo:       domain.SHA256,
	}
	require.Equal(t, true, leading6h.checkZeroes(hash))
	leading5h := HashCach{
		Size:       8,
		Difficulty: 21,
		Algo:        domain.SHA256,
	}
	require.Equal(t, false, leading5h.checkZeroes(hash))
}

func TestHashCach_Verify(t *testing.T) {
	leading6h := HashCach{
		Size:       2 ,
		Difficulty: 22,
		Algo:        domain.SHA256,
	}
	require.True(t, leading6h.Verify([]byte{79,109,97,114,57,56,54,57}, []byte{50, 52}))
}


func TestIntegration(t *testing.T) {
	hash := HashCach{
		Size:       8,
		Difficulty: 8,
		Algo:        domain.SHA256,
		MaxAttempts: 10000,
	}
	prefix := hash.Create()
	nonce, err  := hash.Challenge(prefix)
	require.NoError(t, err)
	hash.Verify(prefix, nonce)
}