package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthFunction(t *testing.T) {
  t.Run("should return false when auth param is absent", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/submit", nil)
    if err!=nil{
      t.Error("go an error where nil was expected", err)
      return
    }
    if authFunction(req)  {
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
