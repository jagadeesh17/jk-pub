package clients

import "time"

// ConfiguredSubscriptionRequest represents the structure of the request body for the Salesforce Configured_Subscription composite API request
type ConfiguredSubscriptionRequest struct {
	AllOrNone        bool                               `json:"allOrNone"`
	CompositeRequest []ConfiguredSubscriptionSubRequest `json:"compositeRequest"`
}

// ConfiguredSubscriptionSubRequest represents a subrequest within the Configured_Subscription composite API request
type ConfiguredSubscriptionSubRequest struct {
	Method      string                            `json:"method"`
	URL         string                            `json:"url"`
	ReferenceID string                            `json:"referenceId"`
	Body        ConfiguredSubscriptionRequestBody `json:"body"`
}

// ConfiguredSubscriptionRequestBody represents the body of a Configured_Subscription body
type ConfiguredSubscriptionRequestBody struct {
	// Define fields based on the structure of the Configured_Subscription object in Salesforce
	Name                           string    `json:"Name,omitempty"`
	Description                    string    `json:"Description,omitempty"`
	StartDate                      time.Time `json:"StartDate,omitempty"`
	EndDate                        time.Time `json:"EndDate,omitempty"`
	SBQQAccountC                   string    `json:"SBQQ__Account__c,omitempty"`
	CustomerAccountC               string    `json:"Customer_Account__c,omitempty"`
	SBQQContractC                  string    `json:"SBQQ__Contract__c,omitempty"`
	DealerC                        string    `json:"Dealer__c,omitempty"`
	AutoRenewC                     bool      `json:"AutoRenew__c,omitempty"`
	InitialTermC                   int       `json:"InitialTerm__c,omitempty"`
	RenewalTermC                   int       `json:"RenewalTerm__c,omitempty"`
	TermSettingTypeC               string    `json:"TermSettingType__c,omitempty"`
	InvoiceSeparateC               bool      `json:"InvoiceSeparate__c,omitempty"`
	SBQQSubscriptionStartDateC     time.Time `json:"SBQQ__SubscriptionStartDate__c"`
	SBQQPackageProductDescriptionC string    `json:"SBQQ__PackageProductDescription__c"`
	CommentsC                      string    `json:"Comments__c,omitempty"`
	VINC                           string    `json:"VIN__c,omitempty"`
	DeviceC                        string    `json:"Device__c,omitempty"`
	OriginTypeC                    string    `json:"OriginType__c,omitempty"`
	VehicleC                       string    `json:"Vehicle__c,omitempty"`
	DataBookCodeC                  string    `json:"DataBookCode__c,omitempty"`
	NewTruckInvoiceDateC           time.Time `json:"NewTruckInvoiceDate__c"`
	NewTruckInvoiceNumberC         int       `json:"NewTruckInvoiceNumber__c"`
	ProductCategoryC               string    `json:"ProductCategory__c,omitempty"`
	SBQQProductC                   string    `json:"SBQQ__Product__c,omitempty"`
	SBQQCustomerPriceC             string    `json:"SBQQ__CustomerPrice__c"`
	SBQQListPriceC                 float64   `json:"SBQQ__ListPrice__c,omitempty"`
	SBQQNetPriceC                  float64   `json:"SBQQ__NetPrice__c,omitempty"`
	SBQQRegularPriceC              float64   `json:"SBQQ__RegularPrice__c,omitempty"`
	SBQQSpecialPriceC              float64   `json:"SBQQ__SpecialPrice__c,omitempty"`
	SBQQRenewalQuantityC           int       `json:"SBQQ__RenewalQuantity__c"`
	SBQQSubscriptionTypeC          string    `json:"SBQQ__SubscriptionType__c"`
	SBQQProductSubscriptionTypeC   string    `json:"SBQQ__ProductSubscriptionType__c"`
	SBQQSubscriptionPricingC       string    `json:"SBQQ__SubscriptionPricing__c"`
	SBQQProrateMultiplierC         float64   `json:"SBQQ__ProrateMultiplier__c"`
	SBQQNumberC                    int       `json:"SBQQ__Number__c,omitempty"`
	SBQQQuantityC                  int       `json:"SBQQ__Quantity__c,omitempty"`
	SBQQBundleC                    int       `json:"SBQQ__Bundle__c,omitempty"`
	SBQQPricingMethodC             int       `json:"SBQQ__PricingMethod__c"`
	ProviderC                      int       `json:"Provider__c,omitempty"`
	SBQQProductOptionC             int       `json:"SBQQ__ProductOption__c"`
	SBQQOptionLevelC               int       `json:"SBQQ__OptionLevel__c,omitempty"`
	SBQQOptionTypeC                int       `json:"SBQQ__OptionType__c,omitempty"`
	SBQQRequiredByIdC              int       `json:"SBQQ__RequiredById__c,omitempty"`
	// Add other fields as needed
}

// ConfiguredSubscriptionResponse represents the structure of the response body for the Salesforce Configured_Subscription composite API response
type ConfiguredSubscriptionResponse struct {
	CompositeResponse []ConfiguredSubscriptionSubResponse `json:"compositeResponse"`
}

// ConfiguredSubscriptionSubResponse represents the response for each subrequest within the Configured_Subscription composite API response
type ConfiguredSubscriptionSubResponse struct {
	HTTPStatusCode int                                `json:"httpStatusCode"`
	Body           ConfiguredSubscriptionResponseBody `json:"body"`
}

// ConfiguredSubscriptionResponseBody represents the body of the Configured_Subscription subrequest response
type ConfiguredSubscriptionResponseBody struct {
	ID      string   `json:"Id"`
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

type Token struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	Type        string `json:"token_type"`
	IssuedAt    string `json:"issued_at"`
	Signature   string `json:"signature"`
}

// NewContractRequest represents the structure of the request body for the Salesforce NewContract composite API request
type NewContractRequest struct {
	AllOrNone        bool                    `json:"allOrNone"`
	CompositeRequest []NewContractSubRequest `json:"compositeRequest"`
}

// NewContractSubRequest represents a subrequest within the NewContract composite API request
type NewContractSubRequest struct {
	Method      string                 `json:"method"`
	URL         string                 `json:"url"`
	ReferenceID string                 `json:"referenceId"`
	Body        NewContractRequestBody `json:"body"`
}

// NewContractRequestBody represents the body of a NewContract body
type NewContractRequestBody struct {
	// Define fields based on the structure of the Configured_Subscription object in Salesforce
	AccountId           string    `json:"AccountId"`
	BillingCountry      string    `json:"BillingCountry"`
	BillingCountryCode  string    `json:"BillingCountryCode"`
	StartDate           time.Time `json:"StartDate"`
	Status              string    `json:"Status"`
	ShippingCountry     string    `json:"ShippingCountry"`
	ShippingCountryCode string    `json:"ShippingCountryCode"`
	ContractTerm        int       `json:"ContractTerm"`
	SBQQRenewalTermC    int       `json:"SBQQ__RenewalTerm__c"`
	// Add other fields as needed
}

// ContractCompositeResponse represents the structure of the response body for the Salesforce NewContract composite API response
type ContractCompositeResponse struct {
	CompositeResponse []ContractResponse `json:"compositeResponse"`
}

// ContractResponse represents the response for each subrequest within the NewContract composite API response
type ContractResponse struct {
	HTTPStatusCode int                  `json:"httpStatusCode"`
	Body           ContractResponseBody `json:"body"`
}

// ContractResponseBody represents the body of the Contract response
type ContractResponseBody struct {
	ID      string   `json:"Id"`
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}
