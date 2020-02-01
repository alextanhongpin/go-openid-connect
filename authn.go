package openid

import "errors"

var ScopeOpenID = "openid"

var ResponseTypeCode = "code"

// Separate required and invalid.
var (
	ErrScopeRequired        = errors.New("scope is required")
	ErrScopeInvalid         = errors.New("scope is invalid")
	ErrResponseTypeRequired = errors.New("response type is required")
	ErrClientIDRequired     = errors.New("client_id is required")
	ErrRedirectURIRequired  = errors.New("redirect_uri is required")
)

type Authn struct {
	Scope        string
	ResponseType string
	ClientID     string
	RedirectURI  string
	State        string
}

func (a *Authn) ValidateRequired() error {
	if a.Scope == "" {
		return ErrScopeRequired
	}
	if a.ResponseType == "" {
		return ErrResponseTypeRequired
	}
	if a.ClientID == "" {
		return ErrClientIDRequired
	}
	if a.RedirectURI == "" {
		return ErrRedirectURIRequired
	}
	return nil
}

func (a *Authn) ValidateScope() error {
	if a.Scope != ScopeOpenID {
		return ErrScopeInvalid
	}
	return nil
}

func NewAuthnRequest() *Authn {
	return &Authn{}
}
