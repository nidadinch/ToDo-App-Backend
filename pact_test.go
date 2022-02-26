package main

import (
	"backend/server"
	"fmt"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestProvider(t *testing.T) {

	server := server.NewServer()
	go server.StartServer(3000)

	pact := &dsl.Pact{
		Host:                     "127.0.0.1",
		Provider:                 "Backend",
		Consumer:                 "FrontEnd",
		DisableToolValidityCheck: true,
	}

	request := types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://localhost:%d", 3000),
		PactURLs:        []string{"/Users/nidadinc/Desktop/Assignment/to-do-app-frontend/pact/pacts/frontend-backend.json"},
	}
	verifyResponse, err := pact.VerifyProvider(t, request)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(len(verifyResponse), "pact tests run")
}
