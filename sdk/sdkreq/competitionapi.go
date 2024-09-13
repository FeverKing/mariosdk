package sdkreq

import (
	"encoding/json"
	"mariosdk/sdk/sdklog"
	"mariosdk/sdk/sdkmodel"
)

type GetUserInfoForCompetitionReq struct {
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
	UserId        string `json:"userId"`
}

type CheckCompetitionPrivilegeReq struct {
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
	UserId        string `json:"userId"`
}

func (ac *ApiClient) CallGetUserInfoForCompetitionApi(request interface{}) (*sdkmodel.GetUserInfoForCompetitionModel, error) {
	res, err := ac.CallApi("/competition/getUserInfoForCompetition", "POST", request)
	if err != nil {
		return nil, err
	}
	var getUserInfoForCompetitionResp sdkmodel.GetUserInfoForCompetitionModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &getUserInfoForCompetitionResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got get user info for competition resp: %v", getUserInfoForCompetitionResp)
	return &getUserInfoForCompetitionResp, nil
}

func (ac *ApiClient) CallCheckCompetitionPrivilegeApi(request interface{}) (*sdkmodel.CheckCompetitionPrivilegeModel, error) {
	res, err := ac.CallApi("/competition/checkCompetitionPrivilege", "POST", request)
	if err != nil {
		return nil, err
	}
	var checkCompetitionPrivilegeResp sdkmodel.CheckCompetitionPrivilegeModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &checkCompetitionPrivilegeResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got check competition privilege resp: %v", checkCompetitionPrivilegeResp)
	return &checkCompetitionPrivilegeResp, nil
}
