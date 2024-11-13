package task3

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

//тестирование с помощью самописной заглушки stub
//Stub - это простой объект, который возвращает заранее определенные ответы на вызовы методов.
//Используется для замены реальных объектов, когда не важна проверка взаимодействия с ними.
//Stub не проверяет, как и сколько раз были вызваны его методы.

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
