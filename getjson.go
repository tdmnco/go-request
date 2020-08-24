package request

import (
	"io/ioutil"
	"net/http"
)

// GetJSON performs a GET request and returns a byte array for unmarshalling JSON
func GetJSON(url string, header map[string]string) ([]byte, *http.Response, error) {
	c := http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	if header == nil || header["Content-Type"] == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range header {
		req.Header.Set(k, header[v])
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
