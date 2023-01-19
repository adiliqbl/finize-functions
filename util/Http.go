package util

import (
	"fmt"
	"io"
	"net/http"
)

func GetApiCall(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	body, readErr := io.ReadAll(response.Body)
	if readErr != nil {
		return "", err
	}
	return string(body), nil
}
