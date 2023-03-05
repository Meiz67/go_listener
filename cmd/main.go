package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	fmt.Println("Listener started")
	db, _ := sql.Open("postgres", "port=5432 host=postgres user=postgres password=root dbname=postgres sslmode=disable")
	err := db.Ping()
	checkError(err)

	conn, err := amqp.Dial("amqp://crm:Tm9oOGVlc29qYWhYMkNhago=@mq.l-wine.ru:5672")
	checkError(err)

	defer conn.Close()

	ch, err := conn.Channel()
	checkError(err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"test_go_consumer", // name
		false,              // durable
		false,              // delete when unused
		false,              // exclusive
		false,              // no-wait
		nil,                // arguments
	)

	checkError(err)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	checkError(err)

	var forever chan struct{}

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages...")
	<-forever

}
