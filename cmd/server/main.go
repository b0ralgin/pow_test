package main

import (
	"fmt"
	"github.com/b0ralgin/pow_test/domain"
	"github.com/b0ralgin/pow_test/gates/bow"
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
	// Запускаем TCP сервер на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error in launching server", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is started")
	worker := pow.HashCach{
		Size:       8,
		Difficulty: 8,
		Algo:       domain.SHA256,
	}
	book := bow.NewSimpleBook()
	for {
		// Принимаем входящее соединение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to create new connection", err)
			continue
		}

		fmt.Println("New client connected", conn.RemoteAddr())

		// Обрабатываем клиента в отдельной горутине
		go handleConnection(conn, worker, book, logger)
	}
}
