package task4

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

//тестирование с помощью генератора мока
//suite test
// //go:generate mockgen -destination=./mock/pricing_mock_generated.go -package=mock  -source pricing.go
// go get go.uber.org/mock/gomock

//Mock - это более сложный объект, который может проверять взаимодействие с ним.
//Используется для проверки того, как тестируемый код взаимодействует с внешними зависимостями.
//Mock может проверять количество вызовов методов, порядок вызовов и переданные параметры.

type Service struct {
	client HTTPClient
	host   string
}

func NewService(host string, client HTTPClient) *Service {
	return &Service{
		client: client,
		host:   host,
	}
}

func (s *Service) Get(urlIn string) (*Response, error) {
	req := http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Host: s.host,
			Path: urlIn,
		},
		Header: map[string][]string{
			"Content-type": {"application/json"},
		},
		Body: nil,
	}

	var resp Response
	respClient, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer respClient.Body.Close()

	bytes, err := io.ReadAll(respClient.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bytes, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
