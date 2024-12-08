package main

import (
	"fmt"
	"time"
)

func main() {
	//region Буферизированный канал

	fmt.Println("--------Буферизированный канал---------")

	ordersBuf := make(chan string, 3) //Буферизированный канал размером 3

	// потребитель размещает заказы
	go func() {
		for i := 0; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order #%d", i)
			ordersBuf <- order // горутина блокируется, если буфер канала заполнен

			fmt.Println("ORDER Placed:", order)
		}

		// закрываем канал
		close(ordersBuf)
	}()

	// Бариста обрабатывает заказы
	for order := range ordersBuf { // Читаем из канала. Это блокирующий вызов
		fmt.Printf("Preparing: %s\n", order)
		time.Sleep(2 * time.Second) // время на подготовку заказа
		fmt.Printf("Served: %s\n", order)
	}

	//endregion

}
