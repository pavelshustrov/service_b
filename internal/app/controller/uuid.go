package controller

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"

	"github.com/pavelshustrov/service_b/client/v1/response"
)

type Service interface {
	GenerateID() string
}

type Controller struct {
	Service Service
}

func (contr Controller) GetUUID(writer http.ResponseWriter, request *http.Request) {
	response := response.ID{
		UUID: uuid.New().String(),
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		http.Error(writer, "json parse error", http.StatusInternalServerError)
	}

	_, _ = writer.Write(bytes)
}
