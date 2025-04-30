package requests

import (
	"fmt"
	"net/http"
)

type GenerateGroupInvtoken struct { 
	TokenId string
	Expiration int64
}

func NewGenerateGroupInvtoken(baseUrl, tokenId string, expiration int64) GenerateGroupInvtoken { 
	return GenerateGroupInvtoken{
		TokenId: tokenId,
		Expiration: expiration,
	} 
}

func (g GenerateGroupInvtoken) Request(baseUrl string) (*http.Request, error) { 
	route := "/integrations/auth-svc/invtoken"

	req, err := http.NewRequest(http.MethodPost, baseUrl + route, nil) 
	if err != nil {
		return nil, err
	}

	req.Header.Add("token-id", g.TokenId)
	req.Header.Add("day-duration", fmt.Sprint(g.Expiration))

	return req, nil
}
