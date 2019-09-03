package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

// getMyIP Returns your current ip by calling api.ipify.org
func getMyIP(retry int) (ip string, err error) {
	if retry == 0 {
		return "", errors.New("Request timed out")
	}

	var newRetry int

	if retry == -1 {
		newRetry = -1
	} else {
		newRetry = retry - 1
	}

	httpClient := http.Client{
		Timeout: time.Duration(3 * time.Second),
	}

	res, err := httpClient.Get("https://api.ipify.org")

	if err != nil {
		return getMyIP(newRetry)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return getMyIP(newRetry)
	}

	ip = string(body)

	return
}
