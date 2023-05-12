package tigaclient

import (
	"fmt"
	"net/http"
)

type Caller struct {
}

type CallerInterface interface {
	DoCall(req *http.Request) (*http.Response, error)
}

type TigaError struct {
	StatusCode int
	Message    string
}

func (e *TigaError) Error() string {
	return fmt.Sprintf("%d: %s", e.StatusCode, e.Message)
}

// Client -
type Client struct {
	tigaURL      string
	token        *jwtToken
	Caller       CallerInterface
	LogResponses bool
}

type jwtToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Role struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Template  string `json:"template,omitempty"`
	ValidFrom string `json:"validFrom,omitempty"`
	//otherBool          bool             `json:"otherBool,omitempty"`
	ValidTo            string           `json:"validTo,omitempty"`
	PreventSelfService bool             `json:"preventSelfService"`
	Description        string           `json:"description,omitempty"`
	SystemInstance     string           `json:"systemInstance,omitempty"`
	ProvisioningType   string           `json:"provisioningType,omitempty"`
	ChildRoles         []string         `json:"childRoles,omitempty"`
	Owners             []string         `json:"owners,omitempty"`
	ApprovalSettings   ApprovalSettings `json:"approvalSettings"`
	UserRequirements   UserRequirements `json:"userRequirements"`
}

type ApprovalSettings struct {
	SkipSystemOwnerApproval    bool     `json:"skipSystemOwnerApproval,omitempty"`
	SkipManagerApproval        bool     `json:"skipManagerApproval,omitempty"`
	SkipRoleOwnerApproval      bool     `json:"skipRoleOwnerApproval,omitempty"`
	NamedApprovers             []string `json:"namedApprovers,omitempty"`
	SecurityClearanceApprovers []string `json:"securityClearanceApprovers,omitempty"`
}

type UserRequirements struct {
	DigitalCommittment bool     `json:"digitalCommittment"`
	TermsAndConditions string   `json:"termsAndConditions,omitempty"`
	Countries          []string `json:"countries,omitempty"`
	BusinessContexts   []string `json:"businessContexts,omitempty"`
}
