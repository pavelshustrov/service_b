package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("service B started")
	http.HandleFunc("/uuid", func(writer http.ResponseWriter, request *http.Request) {
		response := struct {
			UUID string `json:"uuid"`
		}{
			UUID: uuid.New().String(),
		}

		bytes, err := json.Marshal(response)
		if err != nil {
			http.Error(writer, "json parse error", http.StatusInternalServerError)
		}

		_, _ = writer.Write(bytes)
	})

	log.Fatalln(http.ListenAndServe(":8777", nil))
}
