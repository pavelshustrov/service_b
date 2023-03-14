package main

import (
	"context"
	"log"

	"github.com/pavelshustrov/service_b/client/v1"
)

func main() {
	ctx := context.Background()

	client := v1.New()
	id, err := client.Create(ctx)
	if err != nil {
		log.Fatalln(err.Error())
	}

	log.Println(id.UUID)
}
