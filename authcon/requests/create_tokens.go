package requests

import (
	"auth-connector/authcon/requests/regources"
	"bytes"
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateTokens struct { 
	AccountId string
	RoleId    int
	route string
}

func NewCreateTokens(accountId string, roleId int) CreateTokens {
	return CreateTokens{
		AccountId: accountId,
		RoleId:    roleId,
		route:     "/integrations/auth-svc/tokens",
	}
}

func (r *CreateTokens) Request(baseUrl string) (*http.Request, error) {
	regource := regources.CreateTokens{
		Key: regources.Key{
			ID:   "",
			Type: regources.CREATE_TOKENS,
		},
		Attributes: regources.CreateTokensAttributes{
			AccountId: r.AccountId,
			RoleId:    r.RoleId,
		},
	}

	body, err := json.Marshal(regource)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal body")
	}

	return http.NewRequest(http.MethodPost, baseUrl, bytes.NewBuffer(body))
}