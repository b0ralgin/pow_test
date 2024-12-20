package main

import (
	"fmt"
	"github.com/b0ralgin/pow_test/domain"
	"github.com/b0ralgin/pow_test/gates/bow"
	"github.com/b0ralgin/pow_test/gates/config"
	"github.com/b0ralgin/pow_test/gates/pow"
	"go.uber.org/zap"
	"net"
	"os"
)

// реализуем протокол Challenge-Response
func handleConnection(conn net.Conn, pow domain.ProofOfWorker, book domain.Wisdomer, logger *zap.Logger) {
	// Читаем данные от клиента
	defer conn.Close()
	defer func() {
		if e := recover(); e != nil {
			logger.DPanic("panic", zap.Any("description", e))
		}
	}()
	prefix := pow.Create()
	if _, err := conn.Write(prefix); err != nil {
		logger.Error("failed to send puzzle", zap.Error(err))
		return
	}
	// ждем решения
	//TODO: добавить таймаут
	buf := make([]byte, 1024)
	if _, err := conn.Read(buf); err != nil {
		logger.Error("failed to read response", zap.Error(err))
		return
	}
	fmt.Println(buf)
	if ok := pow.Verify(prefix, buf); !ok {
		logger.Error("failed to verify connection")
		return
	}
	qoute, err := book.GetQoute()
	if err != nil {
		logger.Error("failed to get qoute")
		return
	}
	if _, err := conn.Write([]byte(qoute)); err != nil {
		logger.Error("failed to send qoute")
		return
	}
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("failed to load cfg", zap.Error(err))
		os.Exit(1)
	}
	// Запускаем TCP сервер на порту 8080
	listener, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		logger.Fatal("Error in launching server", zap.Error(err))
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("Server is started")
	worker := pow.HashCach{
		Size:       cfg.Size,
		Difficulty: cfg.Difficulty,
		Algo:       cfg.Algo,
	}
	book := bow.NewSimpleBook()
	for {
		// Принимаем входящее соединение
		conn, err := listener.Accept()
		if err != nil {
			logger.Error("Failed to create new connection", zap.Error(err))
			continue
		}
		logger.Info("new client connected", zap.String("ip", conn.RemoteAddr().String()))

		// Обрабатываем клиента в отдельной горутине
		go handleConnection(conn, worker, book, logger)
	}
}
