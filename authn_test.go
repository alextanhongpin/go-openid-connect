package openid_test

import (
	"testing"

	"github.com/alextanhongpin/go-openid-connect"
	"github.com/stretchr/testify/assert"
)

func TestAuthnRequestRequired(t *testing.T) {
	assert := assert.New(t)

	req := openid.NewAuthnRequest()
	assert.NotNil(req.ValidateRequired())
}

func TestAuthnRequestValidation(t *testing.T) {
	assert := assert.New(t)

	t.Run("scope is invalid", func(t *testing.T) {
		req := openid.Authn{
			Scope: "xyz",
		}
		assert.Equal(openid.ErrScopeInvalid, req.ValidateScope())
	})

	t.Run("scope is valid", func(t *testing.T) {
		req := openid.Authn{
			Scope: "openid",
		}
		assert.Nil(req.ValidateScope())
	})
}
