package utils

import (
	"io/ioutil"
	"net/http"
)

// RequestGet get http
func RequestGet(api string) ([]byte, error) {
	resp, err := http.Get(api)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
