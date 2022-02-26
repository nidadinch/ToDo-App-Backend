package main

import (
	"backend/server"
	"net/http"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/utils"
)

func createPact() (pact *dsl.Pact) {
	port, _ := utils.GetFreePort()
	server := server.NewServer()
	server.StartServer(port)

	pact = &dsl.Pact{
		Consumer: "FrontEnd",
		Provider: "Backend",
		Host:     "localhost",
	}
	return pact
}

func Test_GetAllTodoItems(t *testing.T) {
	pact := createPact()

	pact.AddInteraction().
		Given("I get all todo items successfully").
		UponReceiving("A get request for /items").
		WithRequest(dsl.Request{
			Method: http.MethodGet,
			Path:   dsl.String("/items"),
		}).
		WillRespondWith(dsl.Response{
			Status:  http.StatusOK,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body: dsl.StructMatcher{
				"id":   dsl.Integer,
				"text": dsl.Like("blabla"),
			},
		})

	// err := pact.Verify(func() error){
	// 	return
	// }
}
