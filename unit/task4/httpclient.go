package task4

import "net/http"

//go:generate mockgen -destination=./mock/httpclient_mock_generated.go -package=mock  -source httpclient.go

type HTTPClient interface {
	Do(req http.Request) (http.Response, error)
}
