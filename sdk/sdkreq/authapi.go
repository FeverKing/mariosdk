package sdkreq

import (
	"encoding/json"
	"github.com/FeverKing/mariosdk/sdk/sdklog"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
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
	ac.Token = authModel.Token
	sdklog.Infof("authenticated with Token %s", ac.Token)
	return nil
}
