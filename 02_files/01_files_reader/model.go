package main

type Items struct {
	Items []Subject `json:"items"`
}

type Subject struct {
	MasterId        string `json:"masterId"`
	AgreementNumber string `json:"agreementNumber"`
	Inn             string `json:"inn"`
	ClientName      string `json:"clientName"`
	ActivationDate  string `json:"activationDate"`
}
