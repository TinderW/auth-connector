package authcon

import (
	"encoding/json"
	"net/http"

	"github.com/TinderW/auth-connector/authcon/requests"
	"github.com/TinderW/auth-connector/authcon/requests/regources"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

const (
	UserRole = 0	
	AdminRole = 7
) 

type AuthConnector interface {
	Do(req requests.AuthConnectorRequest, value interface{}) error

	CreateTokens(accountId string, roleId int) (*regources.Tokens, error)
	VerifyToken(token string) (*regources.Payload, error)
}

type authConnector struct { 
	url string
}

func (ac *authConnector) Do(req requests.AuthConnectorRequest, value interface{}) error {
	request, err := req.Request(ac.url)
	if err != nil {
		return errors.Wrap(err, "failed to build request")
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}

	if err := json.NewDecoder(response.Body).Decode(&value); err != nil {
		return errors.Wrap(err, "failed to decode response")
	}

	return nil
}

func (ac *authConnector) VerifyToken(token string) (*regources.Payload, error) { 
	req := requests.NewGetPayload(ac.url, token)

	var payload regources.PayloadResponse
	if err := ac.Do(req, &payload); err != nil {
		return nil, errors.Wrap(err, "failed to get payload")
	}

	return &payload.Data, nil
}

func (ac *authConnector) CreateTokens(accountId string, roleId int) (*regources.CreateTokens, error) {
	req := requests.NewCreateTokens(accountId, roleId)

	var tokens regources.CreateTokensResponse
	if err := ac.Do(req, &tokens); err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}

	return &tokens.Data, nil
}