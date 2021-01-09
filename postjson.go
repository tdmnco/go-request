package request

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// PostJSON performs a POST request and returns a byte array for unmarshalling JSON
func PostJSON(url string, header map[string]string, payload []byte) ([]byte, *http.Response, error) {
	c := http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	if err != nil {
		return nil, nil, err
	}

	if header == nil || header["Content-Type"] == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	for k, v := range header {
		req.Header.Set(k, v)
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
