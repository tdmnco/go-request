package request

import (
	"io/ioutil"
	"net/http"
)

// GetJSON performs a GET request and returns a byte array for unmarshalling JSON
func GetJSON(url string) ([]byte, *http.Response, error) {
	c := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, nil, err
	}

	res, err := c.Do(req)

	if err != nil {
		return nil, res, err
	}

	defer res.Body.Close()

	var b []byte

	b, err = ioutil.ReadAll(res.Body)

	if err != nil {
		return b, res, err
	}

	return b, res, nil
}
