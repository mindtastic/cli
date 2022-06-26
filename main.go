package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/mindtastic/cli/openapi"
)

type SubmitRegistrationResponse struct {
	Identity struct {
		Traits struct {
			AccountKey string
		}
	}
}

func main() {
	configuration := openapi.NewConfiguration()
	apiClient := openapi.NewAPIClient(configuration)
	initRegistrationResp, r, err := apiClient.AuthenticationApi.InitUserRegistration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.InitUserRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	flowId := initRegistrationResp.Id

	_, r, err = apiClient.AuthenticationApi.SubmitUserRegistration(context.Background()).Flow(flowId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuthenticationApi.SubmitUserRegistration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}

	decoder := json.NewDecoder(r.Body)
	var srr SubmitRegistrationResponse
	err = decoder.Decode(&srr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error decoding SubmitRegistrationResponse: %v\n", err)
	}

	fmt.Println(srr.Identity.Traits.AccountKey)
}
