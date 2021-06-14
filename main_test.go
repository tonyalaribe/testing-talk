package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthFunction(t *testing.T) {
	t.Run("should return false when auth param is absent", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/submit", nil)
		if err != nil {
			t.Error("go an error where nil was expected", err)
			return
		}
		if authFunction(req) {
			t.Error("got true where false was  expected from authFunction")
			return
		}
	})
	t.Run("should return false when auth param is incorrect", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/submit?auth=XYZ", nil)
		assert.NoError(t, err)
		assert.False(t, authFunction(req))
	})
	t.Run("should return true when auth param is correct", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/submit?auth=ABC", nil)
		assert.NoError(t, err)
		assert.True(t, authFunction(req))
	})
}

func TestBuildRequestContextAndValidate(t *testing.T) {
	testCases := map[string]struct {
		urlToValidate    string
		expectedURLError string
		expectedError    string
	}{
		"happy path with valid age and email": {
			urlToValidate: "/submit?auth=ABC&age=19&name=tony&email=abc@xyz.com",
		},
		"error: age is not a number": {
			urlToValidate: "/submit?auth=ABC&age=agexx&name=tony&email=abc@xyz.com",
			expectedError: "age is not an integer: strconv.Atoi: parsing \"agexx\": invalid syntax",
		},
		"error: age is should be above 18": {
			urlToValidate: "/submit?auth=ABC&age=12&name=tony&email=abc@xyz.com",
			expectedError: "user is below 18 yr",
		},
	}

	for k, v := range testCases {
		t.Run(k, func(t *testing.T) {
			req, err := http.NewRequest("GET", v.urlToValidate, nil)
			if len(v.expectedURLError) == 0 {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, v.expectedURLError)
			}

			_, err = buildRequestContextAndValidate(req)
			if len(v.expectedError) == 0 {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, v.expectedError)
			}
		})
	}
}
