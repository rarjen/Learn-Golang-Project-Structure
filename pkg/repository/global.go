package repository

import (
	"bytes"
	"net/http"
)

type GlobalRepository interface{
	FetchAPIWithBody(
		url string,
		method string,
		bodyRequest []byte,
	) (*http.Response, error)
}

type globalRepository struct {
}

func NewGlobalRepository() GlobalRepository {
	return &globalRepository{}
}

func (gR *globalRepository) FetchAPIWithBody(
	url string,
	method string,
	bodyRequest []byte,
) (*http.Response, error) {
	mapInput := make(map[string]string)
	mapInput["Content-Type"] = "application/json"
	httpClient := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyRequest))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", mapInput["Content-Type"])
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
