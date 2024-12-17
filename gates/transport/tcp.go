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

// реализуем протокол Challenge-Response
func (t *TCP) HandleConnection(conn net.Conn) {
	// Читаем данные от клиента
	defer conn.Close()
	puzzle := t.pow.Create([]byte{})
	if _, err := conn.Write(puzzle); err != nil {
		t.logger.Error("failed to send puzzle", zap.Error(err))
		return
	}
	// ждем решения
	//TODO: добавить таймаут
	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		t.logger.Error("failed to read response", zap.Error(err))
		return
	}
	if ok := t.pow.Verify(puzzle, buf); !ok {
		t.logger.Error("failed to verify connection")
		return
	}
	qoute, err := t.bow.GetQoute()
	if err != nil {
		t.logger.Error("failed to get qoute")
		return
	}
	if _,err := conn.Write([]byte(qoute)); err != nil {
		t.logger.Error("failed to send qoute")
		return
	}
}
