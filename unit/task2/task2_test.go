package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// go get github.com/stretchr/testify
// При использовании testify Вы будете использовать 2 модуля: “require” и “assert”.
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
	if err != nil {
		t.Errorf("Error: %v", err)
	}

	assert.Equal(t, expected, string(got))
}

func TestGroupedParallel(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		//tc := tc // capture range variable
		t.Run(tc.Name, func(t *testing.T) {
			if got := foo(tc.in); got != tc.out {
				t.Errorf("got %v; want %v", got, tc.out)
			}
		})
	}
}
