package request

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Success"))
	}))

	defer s.Close()

	r, err := Get(s.URL)

	if err != nil {
		t.Errorf("GET request failed")
	}

	b, err := ioutil.ReadAll(r.Body)

	if err != nil {
		t.Errorf("Cannot read body from GET response")
	}

	r.Body.Close()

	if string(b) != "Success" {
		t.Errorf("GET response not successful")
	}
}
