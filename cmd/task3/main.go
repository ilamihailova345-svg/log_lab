package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()

	fmt.Println(<-ch) // 1
	fmt.Println(<-ch) // 2
	fmt.Println(<-ch) // 3

	// Ошибка была в четвёртом чтении. Когда канал закрыт, <-ch не блокируется, 
	// а возвращает 0 value - это можно спутать с реальным значением.
	//  v, ok := <-ch — ok == false означает, что канал закрыт и данных больше нет

	// Правильное чтение из канала: проверяем, закрыт ли он
	// v — значение, ok — флаг -  true, если канал ещё открыт
	v, ok := <-ch
	if !ok {
		fmt.Println("Канал закрыт")
		return
	}
	fmt.Println(v)
}