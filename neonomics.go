package neonomics

import (
	"context"
	"errors"
	"time"
)

var (
	ErrUnexpectedError = errors.New("unexpected error")
)

type Path string

const (
	PathTokenRequest             Path = "/auth/realms/live/protocol/openid-connect/token"
	PathTokenRefresh             Path = "/auth/realms/live/protocol/openid-connect/token"
	PathGetSupportedBanks        Path = "/ics/v3/banks"
	PathGetSupportedBankByID     Path = "/ics/v3/banks/%s"
	PathCreateSession            Path = "/ics/v3/session"
	PathGetSessionStatus         Path = "/ics/v3/session/%s"
	PathDeleteSession            Path = "/ics/v3/session/%s"
	PathGetConsent               Path = "/ics/v3/consent/%s"
	PathGetAccounts              Path = "/ics/v3/accounts"
	PathGetAccountByID           Path = "/ics/v3/accounts/%s"
	PathGetBalancesByID          Path = "/ics/v3/accounts/%s/balances"
	PathGetTransactionsByID      Path = "/ics/v3/accounts/%s/transactions"
	PathDomesticPayment          Path = "/ics/v3/payments/domestic-transfer"
	PathDomesticScheduledPayment Path = "/ics/v3/payments/domestic-scheduled-transfer"
	PathSEPAPayment              Path = "/ics/v3/payments/sepa-credit"
	PathSEPAScheduledPayment     Path = "/ics/v3/payments/sepa-scheduled-credit"
	PathGetPaymentByID           Path = "/ics/v3/payments/%s/%s"
	PathAuthorizePayment         Path = "/ics/v3/payments/%s/%s/authorize"
	PathCompletePayment          Path = "/ics/v3/payments/%s/%s/complete"
)

type PathName string

const (
	PathNameTokenRequest             PathName = "TokenRequest"
	PathNameTokenRefresh             PathName = "TokenRefresh"
	PathNameGetSupportedBanks        PathName = "GetSupportedBanks"
	PathNameGetSupportedBankByID     PathName = "GetSupportedBankByID"
	PathNameCreateSession            PathName = "CreateSession"
	PathNameGetSessionStatus         PathName = "GetSessionStatus"
	PathNameDeleteSession            PathName = "DeleteSession"
	PathNameGetConsent               PathName = "GetConsent"
	PathNameGetAccounts              PathName = "GetAccounts"
	PathNameGetAccountByID           PathName = "GetAccountByID"
	PathNameGetBalancesByID          PathName = "GetBalancesByID"
	PathNameGetTransactionsByID      PathName = "GetTransactionsByID"
	PathNameDomesticPayment          PathName = "DomesticPayment"
	PathNameDomesticScheduledPayment PathName = "DomesticScheduledPayment"
	PathNameSEPAPayment              PathName = "SEPAPayment"
	PathNameSEPAScheduledPayment     PathName = "SEPAScheduledPayment"
	PathNameGetPaymentByID           PathName = "GetPaymentByID"
	PathNameAuthorizePayment         PathName = "AuthorizePayment"
	PathNameCompletePayment          PathName = "CompletePayment"
)

type Endpoint string

var (
	EndpointSandbox    Endpoint = "https://sandbox.neonomics.io"
	EndpointProduction Endpoint = "https://api.neonomics.io"
)

type API interface {
	TokenRequest(ctx context.Context, req *TokenRequestRequest) (*TokenRequestResponse, error)
	TokenRefresh(ctx context.Context, req *TokenRefreshRequest) (*TokenRefreshResponse, error)
	GetSupportedBanks(ctx context.Context) ([]*GetSupportedBanksResponse, error)
	GetSupportedBankByID(ctx context.Context, ID string) (*GetSupportedBanksResponse, error)
	CreateSession(ctx context.Context, req *CreateSessionRequest) (*CreateSessionResponse, error)
	GetSessionStatus(ctx context.Context, ID string) (*GetSessionStatusResponse, error)
	DeleteSession(ctx context.Context, ID string) error
	GetConsent(ctx context.Context, ID string) (*GetConsentResponse, error)
	GetAccounts(ctx context.Context) (*GetAccountsResponse, error)
	GetAccountByID(ctx context.Context, ID string) (*GetAccountByIDResponse, error)
	GetBalancesByID(ctx context.Context, ID string) (*GetBalancesByIDResponse, error)
	GetTransactionsByID(ctx context.Context, ID string) (*GetBalancesByIDResponse, error)
	DomesticPayment(ctx context.Context, req *DomesticPaymentRequest) (*DomesticPaymentResponse, error)
	DomesticScheduledPayment(ctx context.Context, req *DomesticScheduledPaymentRequest) (*DomesticScheduledPaymentResponse, error)
	SEPAPayment(ctx context.Context, req *SEPAPaymentRequest) (*SEPAPaymentResponse, error)
	SEPAScheduledPayment(ctx context.Context, req *SEPAScheduledPaymentRequest) (*SEPAScheduledPaymentResponse, error)
	GetPaymentByID(ctx context.Context, ID string) (*GetPaymentByIDResponse, error)
	AuthorizePayment(ctx context.Context, paymentProduct, paymentId string) (*AuthorizePaymentResponse, error)
	CompletePayment(ctx context.Context, paymentProduct, paymentId string) (*CompletePaymentResponse, error)
}

type ErrorResponse struct {
	Id        string `json:"id"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
	Source    string `json:"source"`
	Type      string `json:"type"`
	Timestamp int64  `json:"timestamp"`
	Links     []struct {
		Type string `json:"type"`
		Rel  string `json:"rel"`
		Href string `json:"href"`
		Meta struct {
			Id string `json:"id"`
		} `json:"meta"`
	} `json:"links"`
}

type CompletePaymentResponse struct {
	PaymentId        string    `json:"paymentId"`
	Status           string    `json:"status"`
	CreationDateTime time.Time `json:"creationDateTime"`
}

type AuthorizePaymentResponse struct {
	Links []struct {
		Type string `json:"type"`
		Rel  string `json:"rel"`
		Href string `json:"href"`
		Meta string `json:"meta"`
	} `json:"links"`
	Message   string `json:"message"`
	PaymentId string `json:"paymentId"`
}

type GetPaymentByIDResponse struct {
	InstrumentedAmount                string `json:"instrumentedAmount"`
	PaymentId                         string `json:"paymentId"`
	RemittanceInformationUnstructured string `json:"remittanceInformationUnstructured"`
	RemittanceInformationStructured   struct {
		Reference       string `json:"reference"`
		ReferenceIssuer string `json:"referenceIssuer"`
		ReferenceType   string `json:"referenceType"`
	} `json:"remittanceInformationStructured"`
	Status           string    `json:"status"`
	CreationDateTime time.Time `json:"creationDateTime"`
	CreditorAccount  struct {
		Bban                  string `json:"bban"`
		Iban                  string `json:"iban"`
		SortCodeAccountNumber string `json:"sortCodeAccountNumber"`
	} `json:"creditorAccount"`
	DebtorAccount struct {
		Bban                  string `json:"bban"`
		Iban                  string `json:"iban"`
		SortCodeAccountNumber string `json:"sortCodeAccountNumber"`
	} `json:"debtorAccount"`
	Type     string `json:"type"`
	Currency string `json:"currency"`
}

type SEPAScheduledPaymentRequest struct {
	DebtorAccount struct {
		AccountScheme string `json:"accountScheme"`
		Identifier    string `json:"identifier"`
	} `json:"debtorAccount"`
	DebtorName      string `json:"debtorName"`
	CreditorAccount struct {
		AccountScheme string `json:"accountScheme"`
		Identifier    string `json:"identifier"`
	} `json:"creditorAccount"`
	CreditorName                      string `json:"creditorName"`
	RemittanceInformationUnstructured string `json:"remittanceInformationUnstructured"`
	RemittanceInformationStructured   struct {
		Reference       string `json:"reference"`
		ReferenceIssuer string `json:"referenceIssuer"`
		ReferenceType   string `json:"referenceType"`
	} `json:"remittanceInformationStructured"`
	InstrumentedAmount     string `json:"instrumentedAmount"`
	Currency               string `json:"currency"`
	EndToEndIdentification string `json:"endToEndIdentification"`
	PaymentMetadata        struct {
		CreditorAddress struct {
			StreetName     string `json:"streetName"`
			BuildingNumber string `json:"buildingNumber"`
			PostalCode     string `json:"postalCode"`
			City           string `json:"city"`
			Country        string `json:"country"`
		} `json:"creditorAddress"`
		CreditorAgent struct {
			Identification     string `json:"identification"`
			IdentificationType string `json:"identificationType"`
		} `json:"creditorAgent"`
		PaymentContextCode             string `json:"paymentContextCode"`
		MerchantCategoryCode           string `json:"merchantCategoryCode"`
		MerchantCustomerIdentification string `json:"merchantCustomerIdentification"`
	} `json:"paymentMetadata"`
	RequestedExecutionDate string `json:"requestedExecutionDate"`
}

type SEPAScheduledPaymentResponse struct {
	PaymentId        string    `json:"paymentId"`
	Status           string    `json:"status"`
	CreationDateTime time.Time `json:"creationDateTime"`
}

type SEPAPaymentResponse struct {
	PaymentId        string    `json:"paymentId"`
	Status           string    `json:"status"`
	CreationDateTime time.Time `json:"creationDateTime"`
}

type SEPAPaymentRequest struct {
	DebtorAccount struct {
		Bban                  string `json:"bban"`
		Iban                  string `json:"iban"`
		SortCodeAccountNumber string `json:"sortCodeAccountNumber"`
	} `json:"debtorAccount"`
	DebtorName      string `json:"debtorName"`
	CreditorAccount struct {
		Iban string `json:"iban"`
	} `json:"creditorAccount"`
	CreditorName                      string `json:"creditorName"`
	RemittanceInformationUnstructured string `json:"remittanceInformationUnstructured"`
	RemittanceInformationStructured   struct {
		Reference       string `json:"reference"`
		ReferenceIssuer string `json:"referenceIssuer"`
		ReferenceType   string `json:"referenceType"`
	} `json:"remittanceInformationStructured"`
	InstrumentedAmount     string `json:"instrumentedAmount"`
	Currency               string `json:"currency"`
	EndToEndIdentification string `json:"endToEndIdentification"`
	PaymentMetadata        struct {
		CreditorAddress struct {
			StreetName     string `json:"streetName"`
			BuildingNumber string `json:"buildingNumber"`
			PostalCode     string `json:"postalCode"`
			City           string `json:"city"`
			Country        string `json:"country"`
		} `json:"creditorAddress"`
		CreditorAgent struct {
			Identification     string `json:"identification"`
			IdentificationType string `json:"identificationType"`
		} `json:"creditorAgent"`
		PaymentContextCode             string `json:"paymentContextCode"`
		MerchantCategoryCode           string `json:"merchantCategoryCode"`
		MerchantCustomerIdentification string `json:"merchantCustomerIdentification"`
	} `json:"paymentMetadata"`
}

type DomesticScheduledPaymentRequest struct {
	DebtorAccount struct {
		AccountScheme string  `json:"accountScheme"`
		Identifier    float64 `json:"identifier"`
	} `json:"debtorAccount"`
	DebtorName      string `json:"debtorName"`
	CreditorAccount struct {
		AccountScheme string `json:"accountScheme"`
		Identifier    int    `json:"identifier"`
	} `json:"creditorAccount"`
	CreditorName                      string `json:"creditorName"`
	RemittanceInformationUnstructured string `json:"remittanceInformationUnstructured"`
	RemittanceInformationStructured   struct {
		Reference       string `json:"reference"`
		ReferenceIssuer string `json:"referenceIssuer"`
		ReferenceType   string `json:"referenceType"`
	} `json:"remittanceInformationStructured"`
	InstrumentedAmount     string `json:"instrumentedAmount"`
	Currency               string `json:"currency"`
	EndToEndIdentification string `json:"endToEndIdentification"`
	PaymentMetadata        struct {
		CreditorAddress struct {
			StreetName     string `json:"streetName"`
			BuildingNumber string `json:"buildingNumber"`
			PostalCode     string `json:"postalCode"`
			City           string `json:"city"`
			Country        string `json:"country"`
		} `json:"creditorAddress"`
		CreditorAgent struct {
			Identification     string `json:"identification"`
			IdentificationType string `json:"identificationType"`
		} `json:"creditorAgent"`
		PaymentContextCode             string `json:"paymentContextCode"`
		MerchantCategoryCode           string `json:"merchantCategoryCode"`
		MerchantCustomerIdentification string `json:"merchantCustomerIdentification"`
	} `json:"paymentMetadata"`
	RequestedExecutionDate time.Time `json:"requestedExecutionDate"`
}

type DomesticScheduledPaymentResponse struct {
	PaymentId        string    `json:"paymentId"`
	Status           string    `json:"status"`
	CreationDateTime time.Time `json:"creationDateTime"`
}

type DomesticPaymentRequest struct {
	DebtorAccount struct {
		Bban                  string `json:"bban"`
		Iban                  string `json:"iban"`
		SortCodeAccountNumber string `json:"sortCodeAccountNumber"`
	} `json:"debtorAccount"`
	DebtorName      string `json:"debtorName"`
	CreditorAccount struct {
		Bban                  string `json:"bban"`
		Iban                  string `json:"iban"`
		SortCodeAccountNumber string `json:"sortCodeAccountNumber"`
	} `json:"creditorAccount"`
	CreditorName                      string `json:"creditorName"`
	RemittanceInformationUnstructured string `json:"remittanceInformationUnstructured"`
	RemittanceInformationStructured   struct {
		Reference       string `json:"reference"`
		ReferenceIssuer string `json:"referenceIssuer"`
		ReferenceType   string `json:"referenceType"`
	} `json:"remittanceInformationStructured"`
	InstrumentedAmount     string `json:"instrumentedAmount"`
	Currency               string `json:"currency"`
	EndToEndIdentification string `json:"endToEndIdentification"`
	PaymentMetadata        struct {
		CreditorAddress struct {
			StreetName     string `json:"streetName"`
			BuildingNumber string `json:"buildingNumber"`
			PostalCode     string `json:"postalCode"`
			City           string `json:"city"`
			Country        string `json:"country"`
		} `json:"creditorAddress"`
		CreditorAgent struct {
			Identification     string `json:"identification"`
			IdentificationType string `json:"identificationType"`
		} `json:"creditorAgent"`
		PaymentContextCode             string `json:"paymentContextCode"`
		MerchantCategoryCode           string `json:"merchantCategoryCode"`
		MerchantCustomerIdentification string `json:"merchantCustomerIdentification"`
	} `json:"paymentMetadata"`
}

type DomesticPaymentResponse struct {
	PaymentId        string    `json:"paymentId"`
	Status           string    `json:"status"`
	CreationDateTime time.Time `json:"creationDateTime"`
}

type GetTransactionsByIDResponse struct {
	Id                   string `json:"id"`
	TransactionReference string `json:"transactionReference"`
	TransactionAmount    struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	} `json:"transactionAmount"`
	CreditDebitIndicator string    `json:"creditDebitIndicator"`
	BookingDate          time.Time `json:"bookingDate"`
	ValueDate            time.Time `json:"valueDate"`
	CounterpartyAccount  string    `json:"counterpartyAccount"`
	CounterpartyName     string    `json:"counterpartyName"`
	CounterpartyAgent    string    `json:"counterpartyAgent"`
}

type GetBalancesByIDResponse struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Type     string `json:"type"`
}

type GetAccountByIDResponse struct {
	Id                    string `json:"id"`
	Bban                  string `json:"bban"`
	Iban                  string `json:"iban"`
	SortCodeAccountNumber string `json:"sortCodeAccountNumber"`
	AccountName           string `json:"accountName"`
	AccountType           string `json:"accountType"`
	OwnerName             string `json:"ownerName"`
	DisplayName           string `json:"displayName"`
	Balances              []struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
		Type     string `json:"type"`
	} `json:"balances"`
}

type GetAccountsResponse struct {
	Id                    string `json:"id"`
	Bban                  string `json:"bban"`
	Iban                  string `json:"iban"`
	SortCodeAccountNumber string `json:"sortCodeAccountNumber"`
	AccountName           string `json:"accountName"`
	AccountType           string `json:"accountType"`
	OwnerName             string `json:"ownerName"`
	DisplayName           string `json:"displayName"`
	Balances              []struct {
		Amount   string `json:"amount"`
		Currency string `json:"currency"`
		Type     string `json:"type"`
	} `json:"balances"`
}

type GetConsentResponse struct {
	Message string `json:"message"`
	Links   []struct {
		Type string `json:"type"`
		Rel  string `json:"rel"`
		Href string `json:"href"`
		Meta struct {
			Id string `json:"id"`
		} `json:"meta"`
	} `json:"links"`
}

type GetSessionStatusResponse struct {
	BankId     string `json:"bankId"`
	BankName   string `json:"bankName"`
	CreatedAt  string `json:"createdAt"`
	ProviderId string `json:"providerId"`
}

type CreateSessionRequest struct {
	BankId string `json:"bankId"`
}

type CreateSessionResponse struct {
	SessionId string `json:"sessionId"`
}

type GetSupportedBanksResponse struct {
	Id                             string   `json:"id"`
	CountryCode                    string   `json:"countryCode"`
	BankingGroupName               string   `json:"bankingGroupName"`
	PersonalIdentificationRequired bool     `json:"personalIdentificationRequired"`
	BankDisplayName                string   `json:"bankDisplayName"`
	SupportedServices              []string `json:"supportedServices"`
	Bic                            string   `json:"bic"`
	BankOfficialName               string   `json:"bankOfficialName"`
	Status                         string   `json:"status"`
}

type TokenRequestRequest struct {
	GrantType    string `json:"grant_type"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenRequestResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        string `json:"expires_in"`
	RefreshExpiresIn string `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	SessionState     string `json:"session_state"`
}

type TokenRefreshRequest struct {
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type TokenRefreshResponse struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        string `json:"expires_in"`
	RefreshExpiresIn string `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	TokenType        string `json:"token_type"`
	SessionState     string `json:"session_state"`
}
