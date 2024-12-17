package domain


type ProofOfWorker interface {
	Create(salt []byte) []byte
	Challenge(body []byte) ([]byte)
	Verify(puzzle,body []byte) bool
}


