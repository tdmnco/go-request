package request

import (
	"io/ioutil"
	"net/http"
)

// GetJSON performs a GET request and returns a byte array for unmarshalling JSON
func GetJSON(url string) ([]byte, error) {
	c := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")

	var b []byte

	if err != nil {
		return b, err
	}

	res, err := c.Do(req)

	if err != nil {
		return b, err
	}

	defer res.Body.Close()

	b, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return b, err
	}

	return b, err
}
