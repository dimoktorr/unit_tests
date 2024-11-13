package mock

import (
	"bytes"
	"io"
	"net/http"
)

type HTTPClientStub struct{}

func (h *HTTPClientStub) Do(req http.Request) (http.Response, error) {
	body := `{"StatusCode": 200, "Body": "ok"}`

	return http.Response{
		StatusCode: http.StatusOK,
		Header: map[string][]string{
			"Content-type": {"application/json"},
		},
		Body: io.NopCloser(bytes.NewReader([]byte(body))),
	}, nil
}
