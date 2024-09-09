package sdkreq

import (
	"encoding/json"
	"mariosdk/sdk/sdklog"
	"mariosdk/sdk/sdkmodel"
)

type BatchUserInfoRequest struct {
	Ids []string `json:"ids"`
}

func (ac *ApiClient) CallUserInfoApi(request interface{}) (*sdkmodel.BatchUserInfoModel, error) {
	res, err := ac.CallApi("/user/benchGetUserBase", "POST", request)
	if err != nil {
		return nil, err
	}
	var batchUserInfo sdkmodel.BatchUserInfoModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &batchUserInfo)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got user info: %v", batchUserInfo)
	return &batchUserInfo, nil
}
