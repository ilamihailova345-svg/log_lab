package main

import "fmt"

// ch chan<- int — канал для записи

// отправляет числа от 1 до n в канал и закрывает его
func generate(ch chan<- int, n int) {
	for i := 1; i <= n; i++ {
		ch <- i
	}
	// закрываем канал, иначе for range будет ждать вечно
	close(ch)
}

func main() {
	myCh := make(chan int)

	// генератор в горутине
	go generate(myCh, 10)

	// читаем числа и складываем в сумму
	sum := 0
	for v := range myCh {
		sum += v
	}

	fmt.Println("Итоговая сумма:", sum)
}