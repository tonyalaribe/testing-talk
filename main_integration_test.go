// +build integration

package main

import (
	"context"
	"flag"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/imroc/req"
	"github.com/stretchr/testify/assert"
)

var updateGolden = flag.Bool("update_golden", false, "set to update_golden flag to true if you want to hit the live server")

func TestSendNotification(t *testing.T) {
  var calledRemoteServer bool
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    calledRemoteServer = true
		if *updateGolden {
			res, err := req.Post(notificationURL+"/mail", r.Body)
			assert.NoError(t, err)
			err = os.WriteFile("testdata/notification.golden", res.Bytes(), 0666)
			assert.NoError(t, err)
		}
		data, err := os.ReadFile("testdata/notification.golden")
		assert.NoError(t, err)
		w.Write(data)
	}))
	defer ts.Close()

	err := sendNotification(context.Background(), 22, requestCtx{
		Name:  "Test User",
		Age:   33,
		Email: "user@ex.com",
	}, ts.URL)
	assert.NoError(t, err)
  assert.True(t, calledRemoteServer)
}
