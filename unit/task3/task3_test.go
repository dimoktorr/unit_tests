package task3

import (
	"github.com/dimoktorr/unit_tests/unit/task3/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestService_Get(t *testing.T) {
	//Arrange
	service := NewService("test", &mock.HTTPClientStub{})

	want := Response{
		StatusCode: 200,
		Body:       "ok",
	}

	//Act
	got, err := service.Get("url_test")

	//Assert
	assert.Nil(t, err)
	assert.Equal(t, &want, got)
}
