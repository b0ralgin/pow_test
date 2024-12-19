package main

import (
	"fmt"
	"github.com/b0ralgin/pow_test/domain"
	"github.com/b0ralgin/pow_test/gates/pow"
	"go.uber.org/zap"
	"math"
	"net"
)


func run(logger *zap.Logger) error {
	dial, err := net.Dial("tcp","0.0.0.0:8080")
	if err != nil {
		return err
	}
	defer dial.Close()
	arr := make([]byte, 1024)
	_, err = dial.Read(arr)
	if err != nil {
		return  err
	}
	algo := arr[0]
	size := arr[1]
	diff := arr[2]
	fmt.Println(domain.Algoritm(algo))
	worker := pow.HashCach{
        Size:        size,
        Difficulty:  diff,
		Algo:        domain.Algoritm(algo),
        MaxAttempts: math.MaxInt32,
    }
	fmt.Println(arr[:size])
	nonce, err := worker.Challenge(arr[:size])
	if err != nil {
		return err
	}
	fmt.Println(nonce)
	if _, err := dial.Write(nonce); err != nil {
		return err
	}
	_, err = dial.Read(arr)
	if err != nil {
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
	if err := run(logger); err != nil {
		panic(err)
	}
}