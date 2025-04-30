package requests

import (
	"net/http"
	"strconv"
)

type GenerateGroupInvtoken struct { 
	DaysDuration int
}

func NewGenerateGroupInvtoken(baseUrl string, daysDuration int) GenerateGroupInvtoken { 
	return GenerateGroupInvtoken{
		DaysDuration: daysDuration,
	} 
}

func (g GenerateGroupInvtoken) Request(baseUrl string) (*http.Request, error) { 
	route := "/integrations/auth-svc/invtoken"

	req, err := http.NewRequest(http.MethodPost, baseUrl + route, nil) 
	if err != nil {
		return nil, err
	}

	req.Header.Add("day-duration", strconv.Itoa(g.DaysDuration))

	return req, nil
}
