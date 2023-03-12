package main

import (
	"context"
	"log"

	v1 "serviceB/pkg/client/v1"
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