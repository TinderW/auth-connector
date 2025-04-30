package requests

import (
	"net/http"
	"strconv"
)

type GenerateGroupInvtoken struct { 
	TokenId string
	DaysDuration int
}

func NewGenerateGroupInvtoken(baseUrl, tokenId string, daysDuration int) GenerateGroupInvtoken { 
	return GenerateGroupInvtoken{
		TokenId: tokenId,
		DaysDuration: daysDuration,
	} 
}

func (g GenerateGroupInvtoken) Request(baseUrl string) (*http.Request, error) { 
	route := "/integrations/auth-svc/invtoken"

	req, err := http.NewRequest(http.MethodPost, baseUrl + route, nil) 
	if err != nil {
		return nil, err
	}

	req.Header.Add("token-id", g.TokenId)
	req.Header.Add("day-duration", strconv.Itoa(g.DaysDuration))

	return req, nil
}
