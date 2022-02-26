package main

import (
	"backend/config"
	"backend/server"
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {

	server := server.NewServer()
	go server.StartServer(config.Get().ServerAddr)

	pact := &dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "Backend",
		Consumer:                 "FrontEnd",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", config.Get().ServerAddr),
		PactURLs:        []string{"/Users/nidadinc/Desktop/Assignment/to-do-app-frontend/pact/pacts/frontend-backend.json"},
	}
	verifyResponse, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponse), "pact tests run")
}

// func Test_GetAllTodoItems(t *testing.T) {

// 	pact.AddInteraction().
// 		Given("I get all todo items successfully").
// 		UponReceiving("A get request for /items").
// 		WithRequest(dsl.Request{
// 			Method: http.MethodGet,
// 			Path:   dsl.String("/items"),
// 		}).
// 		WillRespondWith(dsl.Response{
// 			Status:  http.StatusOK,
// 			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
// 			Body: dsl.StructMatcher{
// 				"id":   dsl.Integer,
// 				"text": dsl.Like("blabla"),
// 			},
// 		})

// 	// err := pact.Verify(func() error){
// 	// 	return
// 	// }
// }
