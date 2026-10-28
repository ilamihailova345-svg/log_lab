package main

import (
	"fmt"
	"time"
)

// logs <-chan string — канал для чтения

// читает сообщения из канала и выводит их с подписью [LOG]
func logger(logs <-chan string) {
	// for range читает из канала, он не закрывается
	// поэтому цикл будет ждать новых сообщений бесконечно
	for message := range logs {
		fmt.Println("[LOG]:", message)
	}
}

func main() {
	// позволяет отправителю положить 5 сообщений, не дожидаясь читателя
	logCh := make(chan string, 5)

	// логгер в отдельной горутине
	go logger(logCh)

	// отправляем 8 сообщений
	// Первые 5 попадут в буфер мгновенно. Остальные 3 будут ждать,
	// пока логгер не освободит место
	for i := 0; i < 8; i++ {
		logCh <- fmt.Sprintf("Message %d", i)
	}

	// 2 секунды, чтобы логгер успел всё вывести
	time.Sleep(2 * time.Second)

	// закрываем канал  
	// иначе логгер будет висеть вечно 
	close(logCh)
}