package sdkreq

import (
	"encoding/json"
	"mariosdk/sdk/sdklog"
	"mariosdk/sdk/sdkmodel"
)

type AuthApiRequest struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

func (ac *ApiClient) CallAuthApi(ar interface{}) error {
	res, err := ac.CallApi("/sdk/getSdkToken", "POST", ar)
	if err != nil {
		return err
	}
	var authModel sdkmodel.AuthModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &authModel)
	if err != nil {
		return err
	}
	ac.token = authModel.Token
	sdklog.Infof("authenticated with token %s", ac.token)
	return nil
}
