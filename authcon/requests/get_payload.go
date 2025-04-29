package requests

import "net/http"

type GetPayload struct { 
	Token string
	route string
}

func NewGetPayload(baseUrl string, token string) GetPayload {
	return GetPayload{
		Token: token,
		route: "/integrations/auth-svc/payload",
	}
}

func (g GetPayload) Request(baseUrl string) (*http.Request, error) {
	request, err := http.NewRequest(http.MethodGet, baseUrl + g.route, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Add("token", g.Token)
	return request, nil
}