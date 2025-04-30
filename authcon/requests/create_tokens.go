package requests

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/TinderW/auth-connector/authcon/requests/regources"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type CreateTokens struct { 
	AccountId string
	RoleId    int
}

func NewCreateTokens(accountId string, roleId int) CreateTokens {
	return CreateTokens{
		AccountId: accountId,
		RoleId:    roleId,
	}
}

func (r CreateTokens) Request(baseUrl string) (*http.Request, error) {
	route := "/integrations/auth-svc/auth"

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

	return http.NewRequest(http.MethodPost, baseUrl + route, bytes.NewBuffer(body))
}