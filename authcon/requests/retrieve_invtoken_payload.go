package requests

import "net/http"

type RetrieveInvtokenPayload struct {
	Invtoken string
}

func NewRetrieveInvtokenPayload(invtoken string) RetrieveInvtokenPayload {
	return RetrieveInvtokenPayload{
		Invtoken: invtoken,
	}
}

func (r RetrieveInvtokenPayload) Request(baseUrl string) (*http.Request, error) {
	route := "/integrations/auth-svc/invtoken"

	req, err := http.NewRequest(http.MethodGet, baseUrl + route, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("token", r.Invtoken)
	return req, nil
}