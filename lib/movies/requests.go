package movies

import (
	"bytes"
	"io/ioutil"
	"net/http"
)


/// func to send http Post Request
func PostRequest(url string, data []byte, header bool, headers ...http.Header) ([]byte, http.Header, error) {
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return []byte(""), nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), nil, err
	}

	if !header {
		return body, nil, nil
	}
	return body, res.Header, nil
}


/// func to send http Get Request
func GetRequest(url string, header bool, headers ...http.Header) ([]byte, http.Header, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte(""), nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte(""), nil, err
	}

	if !header {
		return body, nil, nil
	}
	return body, res.Header, nil
}