package requester

import "net/http"

type Requester interface {
	SendRequest(*http.Request) (*http.Response, error)
}
