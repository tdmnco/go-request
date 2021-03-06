package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type getJSONResponse struct {
	FirstName string `json:"firstName,omitempty"`
	ID        int    `json:"id,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

func TestGetJSON(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

			return
		}

		w.Header().Set("Status", "200")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		res := getJSONResponse{"Kasper", 1, "Tidemann"}

		payload, err := json.Marshal(res)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Write(payload)
	}))

	defer s.Close()

	b, _, err := GetJSON(s.URL, nil)

	if err != nil {
		t.Errorf("Request failed")
	}

	if len(b) == 0 {
		t.Errorf("Response body is empty")
	}

	var r getJSONResponse

	err = json.Unmarshal(b, &r)

	if err != nil {
		t.Errorf("Error when unmarshalling payload")
	}

	if r.FirstName != "Kasper" || r.ID != 1 || r.LastName != "Tidemann" {
		t.Errorf("Response payload mismatch")
	}
}
