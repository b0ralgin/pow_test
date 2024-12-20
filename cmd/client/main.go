package main

import (
	"fmt"
	"github.com/b0ralgin/pow_test/domain"
	"github.com/b0ralgin/pow_test/gates/pow"
	"go.uber.org/zap"
	"math"
	"net"
	"flag"
)


func run(logger *zap.Logger, dial net.Conn) error {
	defer dial.Close()
	logger.Info("connect to server")
	arr := make([]byte, 1024)
	_, err := dial.Read(arr)
	if err != nil {
		logger.Error("failed to read prefix", zap.Error(err))
		return  err
	}
	algo := arr[0]
	size := arr[1]
	diff := arr[2]
	worker := pow.HashCach{
        Size:        size,
        Difficulty:  diff,
		Algo:        domain.Algoritm(algo),
        MaxAttempts: math.MaxInt32,
    }
	nonce, err := worker.Challenge(arr[:size])
	if err != nil {
		logger.Error("failed to solve challenge", zap.Error(err))
		return err
	}
	if _, err := dial.Write(nonce); err != nil {
		logger.Error("failed to send nonce", zap.Error(err))
		return err
	}
	_, err = dial.Read(arr)
	if err != nil {
		logger.Error("failed to read qoute", zap.Error(err))
		return  err
	}
	fmt.Println(string(arr))
	return nil
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	var serverAddr string
	flag.StringVar(&serverAddr, "s", "0.0.0.0:8080", "address of the server")
	flag.Parse()
	dial, err := net.Dial("tcp",serverAddr)
	if err != nil {
		logger.Fatal("failed to connect to server", zap.Error(err))
		return 
	}
	if err := run(logger, dial ); err != nil {
		panic(err)
	}
}