package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go get github.com/stretchr/testify
// При использовании testify Вы будете использовать модуль: “assert”.
// табличные юнит тесты c t.Parallel()

func TestRequestHandler(t *testing.T) {
	//Arrange
	expected := "Hello john"

	req := httptest.NewRequest(http.MethodGet, "/greet?name=john", nil)
	w := httptest.NewRecorder()

	//Act
	RequestHandler(w, req)

	res := w.Result()
	defer res.Body.Close()
	got, err := io.ReadAll(res.Body)
	assert.Nil(t, err)

	assert.Equal(t, expected, string(got))
}

func TestRequestHandlerParallel(t *testing.T) {
	//Arrange
	testData := []struct {
		description string
		name        string
		expected    string
	}{
		{
			description: "name is john",
			name:        "john",
			expected:    "Hello john",
		},
		{
			description: "name is empty",
			name:        "",
			expected:    "You must supply a name",
		},
		{
			description: "name is gabriel",
			name:        "gabriel",
			expected:    "Hello gabriel",
		},
		{
			description: "name is roman",
			name:        "roman",
			expected:    "Hello roman",
		},
	}

	t.Parallel()
	for _, tt := range testData {
		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/greet?name=%s", tt.name), nil)
		w := httptest.NewRecorder()

		//tc := tc // capture range variable
		t.Run(tt.description, func(t *testing.T) {
			//Act
			RequestHandler(w, req)

			res := w.Result()
			defer res.Body.Close()
			got, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("Error: %v", err)
			}

			assert.Equal(t, tt.expected, string(got))
		})
	}
}
