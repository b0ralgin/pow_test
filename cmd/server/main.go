package main

import (
	"fmt"
	"os"
	"net"
	"bufio"
)


func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Читаем данные от клиента
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения данных:", err)
			return
		}
		fmt.Print("Получено сообщение от клиента:", message)

		// Отправляем ответ клиенту
		response := "Сервер получил: " + message
		_, err = conn.Write([]byte(response))
		if err != nil {
			fmt.Println("Ошибка при отправке данных:", err)
			return
		}
	}
}

func main() {
	// Запускаем TCP сервер на порту 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error in launching server", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server is started")

	for {
		// Принимаем входящее соединение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to create new connection", err)
			continue
		}

		fmt.Println("New client connected", conn.RemoteAddr())

		// Обрабатываем клиента в отдельной горутине
		go handleConnection(conn)
	}
}