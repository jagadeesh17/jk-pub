package api

import (
	"app/clients"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateNewContract
func CreateNewContract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var contracts []clients.NewContractRequestBody

	err := json.NewDecoder(r.Body).Decode(&contracts)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	accountIds := make([]string, len(contracts))
	for _, contract := range contracts {
		if contract.AccountId == "" {
			http.Error(w, "AccountId is required", http.StatusBadRequest)
			return
		}
		accountIds = append(accountIds, contract.AccountId)
	}
	fmt.Printf("Creating contract for AccountIds: %s\n", accountIds)
	response, err := clients.NewContract(contracts)
	if err != nil {
		http.Error(w, "Failed to create contracts", http.StatusInternalServerError)
		return
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set content type header and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
	return
}

// CreateNewContract
func ConfigureSubscription(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var subscriptions []clients.ConfiguredSubscriptionRequestBody

	err := json.NewDecoder(r.Body).Decode(&subscriptions)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}
	for _, subscription := range subscriptions {
		if subscription.SBQQQuantityC <= 0 {
			http.Error(w, "SBQQ__Quantity__c is required", http.StatusBadRequest)
			return
		}
	}
	response, err := clients.ConfigureSubscription(subscriptions)
	if err != nil {
		http.Error(w, "Failed to create contracts", http.StatusInternalServerError)
		return
	}
	// Convert struct to JSON
	jsonData, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set content type header and write response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonData)
	return
}
