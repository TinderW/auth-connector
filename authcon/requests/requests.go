package requests

import "net/http"

type AuthConnectorRequest interface {
	Request(baseUrl string) (*http.Request, error)
}