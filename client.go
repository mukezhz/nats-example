package main

import (
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
)

type Order struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func main() {
	order := Order{ProductID: "ABC123", Quantity: 2}

	orderData, err := json.Marshal(order)
	if err != nil {
		panic(err)
	}
	fmt.Println(nats.DefaultURL)
	conn, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	err = conn.Publish("orders", orderData)
	if err != nil {
		panic(err)
	}

	fmt.Println("Order submitted successfully!")
}
