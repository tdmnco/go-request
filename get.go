package request

import (
	"net/http"
)

// Get performs a GET request
func Get(url string) (*http.Response, error) {
	c := http.Client{}
	res := &http.Response{}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return res, err
	}

	return c.Do(req)
}
