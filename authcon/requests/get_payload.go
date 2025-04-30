package requests

import "net/http"

type GetPayload struct { 
	Token string
}

func NewGetPayload(baseUrl string, token string) GetPayload {
	return GetPayload{
		Token: token,
	}
}

func (g GetPayload) Request(baseUrl string) (*http.Request, error) {
	route := "/integrations/auth-svc/auth"

	request, err := http.NewRequest(http.MethodGet, baseUrl + route, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("token", g.Token)
	return request, nil
}