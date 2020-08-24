package request

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	FirstName string `json:"firstName,omitempty"`
	ID        int    `json:"id,omitempty"`
	LastName  string `json:"lastName,omitempty"`
}

func TestGetJSON(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Status", "200")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		user := User{"Kasper", 1, "Tidemann"}

		payload, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

			return
		}

		w.Write(payload)
	}))

	defer s.Close()

	var user User

	b, _, err := GetJSON(s.URL, nil)

	if err != nil {
		t.Errorf("Request failed")
	}

	if len(b) == 0 {
		t.Errorf("Response body is empty")
	}

	err = json.Unmarshal(b, &user)

	if err != nil {
		t.Errorf("Error when unmarshalling payload")
	}

	if user.FirstName != "Kasper" || user.ID != 1 || user.LastName != "Tidemann" {
		t.Errorf("Response payload mismatch")
	}
}
