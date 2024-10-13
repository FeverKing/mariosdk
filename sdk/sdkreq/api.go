package sdkreq

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
	"io"
)

type ApiClient struct {
	requester Requester
	baseUrl   string
	Token     string
}

func NewApiClient(requester Requester, baseUrl string) *ApiClient {
	return &ApiClient{
		requester: requester,
		baseUrl:   baseUrl,
	}
}

func (ac *ApiClient) CallApi(path, method string, payload interface{}) (interface{}, error) {
	var body io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		body = bytes.NewBuffer(jsonData)
		// sdklog.Infof("request body: %s", jsonData)
	}

	url := ac.baseUrl + path
	req, err := NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	if ac.Token != "" {
		req.Header.Set("Authorization", ac.Token)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	ac.requester = NewHttpRequester()
	res, err := ac.requester.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var responseBody []byte
	responseBody, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var baseResponse *sdkmodel.BaseResponse
	err = json.Unmarshal(responseBody, &baseResponse)
	if err != nil {
		return nil, err
	}
	if baseResponse.Code != 200 {
		return nil, errors.New(baseResponse.Msg)
	}
	return baseResponse.Data, nil
}
