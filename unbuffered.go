package main

import (
	"fmt"
	"time"
)

func main() {

	//region Небуферизированный канал
	fmt.Println("--------Небуферизированный канал---------")
	orders := make(chan string)

	// потребитель размещает заказы
	go func() {
		for i := 0; i <= 5; i++ {
			order := fmt.Sprintf("Coffee order #%d", i)
			orders <- order // горутина блокируется, пока бариста не примет заказ

			fmt.Println("ORDER Placed:", order)
		}

		// закрываем канал
		close(orders)
	}()

	// Бариста обрабатывает заказы
	for order := range orders { // Читаем из канала. Это блокирующий вызов
		fmt.Printf("Preparing: %s\n", order)
		time.Sleep(2 * time.Second) // время на подготовку заказа
		fmt.Printf("Served: %s\n", order)
	}

	//endregion

}
