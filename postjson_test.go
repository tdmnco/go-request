package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type postJSONRequest struct {
	FirstName string `json:"firstName,omitempty"`
	ID        int    `json:"id,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

type postJSONResponse struct {
	Status string `json:"status,omitempty"`
}

func TestPostJSON(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

			return
		}

		w.Header().Set("Status", "200")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		res := postJSONResponse{"Success"}

		payload, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Write(payload)
	}))

	defer s.Close()

	req := postJSONRequest{"Kasper", 1, "Tidemann"}

	m, err := json.Marshal(req)

	if err != nil {
		t.Errorf("Marshalling failed")
	}

	b, _, err := PostJSON(s.URL, nil, m)

	if err != nil {
		t.Errorf("Request failed")
	}

	if len(b) == 0 {
		t.Errorf("Response body is empty")
	}

	var res postJSONResponse

	err = json.Unmarshal(b, &res)

	if err != nil {
		t.Errorf("Error when unmarshalling payload")
	}

	if res.Status != "Success" {
		t.Errorf("Response status mismatch")
	}
}
