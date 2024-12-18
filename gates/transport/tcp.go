package transport

import (
	"github.com/b0ralgin/pow_test/domain"
	"net"
	"go.uber.org/zap"
)


type TCP struct {
	pow domain.ProofOfWorker
	bow domain.Wisdomer
	logger *zap.Logger
}



