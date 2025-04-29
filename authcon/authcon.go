package authcon

import (
	"auth-connector/authcon/requests"
	"encoding/json"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

type AuthConnector interface {
	Do(req requests.AuthConnectorRequest) (interface{}, error)
}

type authConnector struct { 
	url string
}

func (ac *authConnector) Do(req requests.AuthConnectorRequest) (interface{}, error) {
	request, err := req.Request(ac.url)
	if err != nil {
		return nil, errors.Wrap(err, "failed to build request")
	}

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send request")
	}

	var resourceResponse interface{}
	if err := json.NewDecoder(response.Body).Decode(&resourceResponse); err != nil {
		return nil, errors.Wrap(err, "failed to decode response")
	}

	return response, nil
}