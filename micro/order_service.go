package main

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

type Order struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		panic(err)
	}
	defer nc.Close()

	orderConfig := micro.Config{
		Name:        "order",
		Version:     "0.0.1-alpha",
		Description: "A simple order service",
	}
	orderService, err := micro.AddService(nc, orderConfig)
	if err != nil {
		panic(err)
	}
	orderService.AddEndpoint("orders", micro.HandlerFunc(orderHandler))
	// keep the program running forever
	select {}
}

func orderHandler(r micro.Request) {
	log.Println("DATA: ", string(r.Data()), "SUBJECT: ", r.Subject())
	order := Order{
		ProductID: "ABC123",
		Quantity:  2,
	}
	r.RespondJSON(order)
}
