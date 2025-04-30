package authcon

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

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

	RetrieveInvtokenPayload(token string) (*regources.InvtokenPayload, error)
	GenerateGroupInvtoken(dayDuration int) (*regources.GroupInviteToken, error)
}

type authConnector struct { 
	url string
	client http.Client
}

func NewAuthConnector(url string) AuthConnector {
	return &authConnector{
		url: url,
		client: http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (ac *authConnector) Do(req requests.AuthConnectorRequest, value interface{}) error {
	request, err := req.Request(ac.url)
	if err != nil {
		return errors.Wrap(err, "failed to build request")
	}

	response, err := ac.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "failed to send request")
	}
	defer response.Body.Close()

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			return errors.New("auth-svc is unavailable")
		}
		return errors.Errorf("auth-svc is unavailable, %s: %s", response.Status, string(body))
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

func (ac *authConnector) CreateTokens(accountId string, roleId int) (*regources.Tokens, error) {
	req := requests.NewCreateTokens(accountId, roleId)

	var tokens regources.TokensResponse
	if err := ac.Do(req, &tokens); err != nil {
		return nil, errors.Wrap(err, "failed to get tokens")
	}

	return &tokens.Data, nil
}

func (ac *authConnector) GenerateGroupInvtoken(dayDuration int) (*regources.GroupInviteToken, error) { 
	req := requests.NewGenerateGroupInvtoken(ac.url, dayDuration)

	var token regources.GroupInviteTokenResponse
	if err := ac.Do(req, &token); err != nil {
		return nil, errors.Wrap(err, "failed to get group invite token")
	}

	return &token.Data, nil
}

func (ac *authConnector) RetrieveInvtokenPayload(token string) (*regources.InvtokenPayload, error) {
	req := requests.NewRetrieveInvtokenPayload(token)

	var payload regources.InvtokenPayloadResponse
	if err := ac.Do(req, &payload); err != nil {
		return nil, errors.Wrap(err, "failed to get payload")
	}

	return &payload.Data, nil
}