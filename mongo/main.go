package main

import (
	"context"
	"fmt"
	"log"
	"mongo/database"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("no .env file found")
	}

	uri := os.Getenv("MONGO_CONN")
	if uri == "" {
		log.Fatal("you must set MONGO_CONN in .env file")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// database.ReadSampleMflix(client, "The Room")

	// database.AddBook(client, "Atonment", "Ian McEwan")

	database.FindAllEpisodes(client)

	type Currency int

	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	q := [3]int{1, 2}
	v := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	w := [...]int{10: -1}

	fmt.Printf("%+v %T\n", q, q)
	fmt.Printf("%+v %T\n", v, v)
	fmt.Printf("%+v %T\n", w, w)
}
