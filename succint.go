package main

type ProofRequestData struct {
	Input string `json:"input"`
}

type ProofRequest struct {
	Type      string           `json:"type"`
	ReleaseID string           `json:"releaseId"`
	Data      ProofRequestData `json:"data"`
}

type Organization struct {
	ID        string `json:"id"`
	OrgName   string `json:"org_name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Edges     string `json:"edges"`
}

type Project struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	GitAccountName string `json:"git_account_name"`
	OrganizationID string `json:"organization_id"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	Edges          struct {
		Organization Organization `json:"organization"`
	} `json:"edges"`
}

type Release struct {
	ID        string `json:"id"`
	Number    int    `json:"number"`
	ProjectID string `json:"project_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Edges     struct {
		Project Project `json:"project"`
	} `json:"edges"`
}

type Request struct {
	ID               string `json:"id"`
	ChainID          int    `json:"chain_id"`
	GatewayAddress   string `json:"gateway_address"`
	Sender           string `json:"sender"`
	Origin           string `json:"origin"`
	FunctionID       string `json:"function_id"`
	Input            string `json:"input"`
	CallbackAddress  string `json:"callback_address"`
	CallbackData     string `json:"callback_data"`
	CallbackGasLimit int    `json:"callback_gas_limit"`
	RequestType      string `json:"request_type"`
	Status           string `json:"status"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

type Edges struct {
	Release  Release   `json:"release"`
	Requests []Request `json:"requests"`
}

type SuccintProof struct {
	ID           string       `json:"id"`
	Status       string       `json:"status"`
	CreatedAt    string       `json:"created_at"`
	UpdatedAt    string       `json:"updated_at"`
	ProofRequest ProofRequest `json:"proof_request"`
	ProofRelease string       `json:"proof_release"`
	Edges        Edges        `json:"edges"`
}
