package util

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

func MakeApiCall(req *http.Request) (map[string]interface{}, error) {
	client := &http.Client{Timeout: 30 * time.Second}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	} else if res.StatusCode == http.StatusBadRequest {
		return nil, errors.New(res.Status)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	return ParseApiResponse(res)
}

func ParseApiResponse(result *http.Response) (map[string]interface{}, error) {
	body, readErr := io.ReadAll(result.Body)
	if readErr != nil {
		return nil, readErr
	}

	var docType map[string]interface{}
	err := json.Unmarshal(body, &docType)
	if err != nil {
		log.Fatalf("MapTo: %v", err)
	}
	return docType, err
}
