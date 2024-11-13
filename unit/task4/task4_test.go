package task4

import (
	"bytes"
	"github.com/dimoktorr/unit_tests/unit/task4/mock"
	"io"
	"net/http"
	"net/url"
	"testing"
)

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

func TestSuites(t *testing.T) {
	suite.Run(t, new(ServiceSuite))
}

type ServiceSuite struct {
	suite.Suite
	service    *Service
	httpClient *mock.MockHTTPClient
}

func (s *ServiceSuite) SetupTest() {
	ctrl := gomock.NewController(s.T())

	s.httpClient = mock.NewMockHTTPClient(ctrl)
	s.service = NewService("host_test", s.httpClient)
}

func (s *ServiceSuite) TearDownTest() {
	s.T().Log("TearDownTest")
}

func (s *ServiceSuite) TestGet() {
	//Arrange
	want := Response{
		StatusCode: 200,
		Body:       "ok",
	}

	s.httpClient.EXPECT().Do(gomock.Any()).Return(
		http.Response{
			StatusCode: http.StatusOK,
			Header: map[string][]string{
				"Content-type": {"application/json"},
			},
			Body: io.NopCloser(bytes.NewReader([]byte(`{"StatusCode": 200, "Body": "ok"}`))),
		}, nil,
	)

	//Act
	got, err := s.service.Get("url_test")

	//Assert
	s.Nil(err)
	s.Equal(&want, got)
}

func (s *ServiceSuite) TestGet_notGoMockAny() {
	//Arrange
	want := Response{
		StatusCode: 200,
		Body:       "ok",
	}

	req := http.Request{
		Method: http.MethodGet,
		URL: &url.URL{
			Host: "host_test",
			Path: "url_test",
		},
		Header: map[string][]string{
			"Content-type": {"application/json"},
		},
		Body: nil,
	}

	s.httpClient.EXPECT().Do(req).Return(
		http.Response{
			StatusCode: http.StatusOK,
			Header: map[string][]string{
				"Content-type": {"application/json"},
			},
			Body: io.NopCloser(bytes.NewReader([]byte(`{"StatusCode": 200, "Body": "ok"}`))),
		}, nil,
	)

	//Act
	got, err := s.service.Get("url_test")

	//Assert
	s.Nil(err)
	s.Equal(&want, got)
}

func (s *ServiceSuite) TestGet_ErrCode() {
	//Arrange
	want := Response{
		StatusCode: 200,
		Body:       "ok",
	}

	s.httpClient.EXPECT().Do(gomock.Any()).Return(
		http.Response{
			StatusCode: http.StatusOK,
			Header: map[string][]string{
				"Content-type": {"application/json"},
			},
			Body: io.NopCloser(bytes.NewReader([]byte(`{"StatusCode": 400, "Body": "bad_request"}`))),
		}, nil,
	)

	//Act
	got, err := s.service.Get("url_test")

	//Assert
	s.Nil(err)
	s.Equal(&want, got)
}

func (s *ServiceSuite) TestGet_ErrExpect() {
	//Arrange
	want := Response{
		StatusCode: 200,
		Body:       "ok",
	}

	//Act
	got, err := s.service.Get("url_test")

	//Assert
	s.Nil(err)
	s.Equal(&want, got)
}
