package main

import (
	"fmt"
	"time"
)

type Customer struct {
	Name    string
	Address string
}

type OrderItem struct {
	ItemName  string
	Qty       int
	UnitPrice float64
}

type Order struct {
	Customer    *Customer
	OrderItems  []*OrderItem
	Status      string
	CreatedAt   time.Time
	TotalAmount float64
}

func (order *Order) CalculateTotal() {
	order.TotalAmount = 0
	for _, item := range order.OrderItems {
		order.TotalAmount += float64(item.Qty) * item.UnitPrice
	}
}

func (order *Order) ChangeStatus(newStatus string) {
	order.Status = newStatus
}

func main() {
	customer := Customer{Name: "Иван Иванов", Address: "Курган, ул. Ленина, 25"}

	item1 := OrderItem{"Футболка", 2, 500.0}
	item2 := OrderItem{"Джинсы", 1, 2000.0}
	item3 := OrderItem{"Кроссовки", 1, 3000.0}

	order := Order{
		Customer:   &customer,
		OrderItems: []*OrderItem{&item1, &item2, &item3},
		Status:     "Создан",
		CreatedAt:  time.Now(),
	}

	fmt.Println("Клиент:", order.Customer.Name)
	fmt.Println("Адрес:", order.Customer.Address)
	fmt.Println("Дата заказа:", order.CreatedAt.Format("02.01.2006"))

	fmt.Println("\nТовары в заказе:")
	for i, item := range order.OrderItems {
		fmt.Printf("%d. %s - %d шт. по %.2f руб.\n", i+1, item.ItemName, item.Qty, item.UnitPrice)
	}

	order.CalculateTotal()
	fmt.Printf("\nИтоговая сумма: %.2f руб.\n", order.TotalAmount)

	fmt.Println("\nСтатус заказа:", order.Status)
	order.ChangeStatus("Оплачен")
	fmt.Println("Новый статус заказа:", order.Status)

	order.ChangeStatus("Доставляется")
	fmt.Println("Статус заказа:", order.Status)

	order.ChangeStatus("Доставлен")
	fmt.Println("Финальный статус:", order.Status)

	fmt.Println("\nЗаказ завершен!")
}
