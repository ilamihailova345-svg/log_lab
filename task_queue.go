package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	TimeCreated string 
}

// читает задачи из канала и выполняет их по одной
func worker(id int, tasks <-chan Task) {
	for task := range tasks {
		fmt.Printf("Воркер %d начал задачу %d (создана в %s)\n", id, task.ID, task.TimeCreated)

		time.Sleep(500 * time.Millisecond)

		fmt.Printf("Воркер %d закончил задачу %d\n", id, task.ID)
	}
}

func main() {
	// буфер на 10 задач позволяет main быстро вывести первыми 10 задачами,
	// не дожидаясь, пока воркер их заберёт
	taskCh := make(chan Task, 10)


	// дополнительное задание
	// конкурируют за задачи из одного канала
	go worker(1, taskCh)
	go worker(2, taskCh)
	go worker(3, taskCh)

	// отправляем 15 задач
	// Первые 10 попадут в буфер мгновенно. Остальные 5 будут ждать,
	// пока воркеры не освободят место
	for i := 1; i <= 15; i++ {
		task := Task{
			ID:          i,
			TimeCreated: time.Now().Format(time.RFC3339),
		}
		taskCh <- task
		fmt.Printf("Отправлена задача %d\n", i)
	}
	
	// сигнал воркерам, что цикл range можно завершать
	close(taskCh)

	// ждём, пока воркеры разберут остальное
	time.Sleep(4 * time.Second)

	fmt.Println("Все задачи обработаны.")
}