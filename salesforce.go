package clients

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Salesforce credentials
const (
	SFUsername      = "your_salesforce_username"
	SFPassword      = "your_salesforce_password"
	SFClientID      = "your_salesforce_client_id"
	SFClientSecret  = "your_salesforce_client_secret"
	SFLoginEndpoint = "https://login.salesforce.com"
	SFInstanceUrl   = ""
)

func getSalesforceAccessToken() (string, error) {
	authURL := fmt.Sprintf("%s/services/oauth2/token", SFLoginEndpoint)
	payload := strings.NewReader(fmt.Sprintf("grant_type=password&client_id=%s&client_secret=%s&username=%s&password=%s",
		SFClientID, SFClientSecret, SFUsername, SFPassword))

	req, err := http.NewRequest("POST", authURL, payload)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response and extract access token
	var token Token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}

	return token.AccessToken, nil
}

func NewContract(contracts []NewContractRequestBody) ([]*ContractResponseBody, error) {
	requestBody := createNewContractCompositeRequest(contracts)
	// Make the HTTP request
	responseBody, err := makeCompositeRequest(requestBody)
	if err != nil {
		return nil, err
	}

	// Process the response
	fmt.Println("Response:", string(responseBody))

	// Parse the response and extract NewContract body
	var newContract ContractCompositeResponse
	err = json.Unmarshal(requestBody, &newContract)

	if err != nil {
		return nil, err
	}
	if len(newContract.CompositeResponse) == 0 {
		return nil, errors.New("no response found")
	}
	var newContracts []*ContractResponseBody
	for i := range newContract.CompositeResponse {
		newContracts = append(newContracts, &newContract.CompositeResponse[i].Body)
	}
	return newContracts, nil
}

func createNewContractCompositeRequest(contracts []NewContractRequestBody) []byte {
	var compositeRequest []NewContractSubRequest
	for i := range contracts {
		contracts[i].StartDate = time.Now()

		compositeRequest = append(compositeRequest,
			NewContractSubRequest{
				Method:      "POST",
				URL:         "/services/data/v53.0/sobjects/Contract",
				ReferenceID: "NewContract",
				Body:        contracts[i],
			},
		)
	}
	contractRequest := NewContractRequest{
		AllOrNone:        true,
		CompositeRequest: compositeRequest,
	}

	// Convert the composite request to JSON
	requestBody, _ := json.Marshal(contractRequest)

	return requestBody
}

func makeCompositeRequest(requestBody []byte) ([]byte, error) {
	path := SFInstanceUrl + "/services/data/v53.0/composite"

	// Create an HTTP client
	client := &http.Client{}

	// Create an HTTP request
	request, err := http.NewRequest("POST", path, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	//Get AccessToken
	accessToken, err := getSalesforceAccessToken()
	if err != nil {
		return nil, err
	}
	// Set request headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)

	// Make the HTTP request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("Invalid response: " + response.Status)
	}
	// Read the response body
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}

func ConfigureSubscription(subscriptions []ConfiguredSubscriptionRequestBody) ([]*ConfiguredSubscriptionResponseBody, error) {
	requestBody := createConfigureSubscriptionCompositeRequest(subscriptions)
	// Make the HTTP request
	responseBody, err := makeCompositeRequest(requestBody)
	if err != nil {
		return nil, err
	}

	// Process the response
	fmt.Println("Response:", string(responseBody))

	// Parse the response and extract ConfigureSubscription body
	var configureSubsResp ConfiguredSubscriptionResponse
	err = json.Unmarshal(requestBody, &configureSubsResp)

	if err != nil {
		return nil, err
	}
	if len(configureSubsResp.CompositeResponse) == 0 {
		return nil, errors.New("no response found")
	}
	var newSubscriptions []*ConfiguredSubscriptionResponseBody
	for i := range configureSubsResp.CompositeResponse {
		newSubscriptions = append(newSubscriptions, &configureSubsResp.CompositeResponse[i].Body)
	}
	return newSubscriptions, nil
}

func createConfigureSubscriptionCompositeRequest(subscriptions []ConfiguredSubscriptionRequestBody) []byte {
	var compositeRequest []ConfiguredSubscriptionSubRequest

	for i := range subscriptions {
		subscriptions[i].StartDate = time.Now()
		compositeRequest = append(compositeRequest,
			ConfiguredSubscriptionSubRequest{
				Method:      "POST",
				URL:         "/services/data/v53.0/sobjects/SBQQ__Subscription__c",
				ReferenceID: "Configured_Subscription",
				Body:        subscriptions[i],
			},
		)
	}
	contractRequest := ConfiguredSubscriptionRequest{
		AllOrNone:        true,
		CompositeRequest: compositeRequest,
	}

	// Convert the composite request to JSON
	requestBody, _ := json.Marshal(contractRequest)

	return requestBody
}
