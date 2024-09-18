package sdkreq

import (
	"encoding/json"
	"github.com/FeverKing/mariosdk/sdk/sdklog"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
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

type StartChallengeContainerReq struct {
	CompetitionId string   `json:"competitionId"`
	SecretKey     string   `json:"secretKey"`
	ContainerId   string   `json:"containerId"`
	DockerImage   string   `json:"dockerImage"`
	HttpPort      []string `json:"httpPort"`
	TcpPort       []string `json:"tcpPort"`
	IsStatic      bool     `json:"isStatic"`
	Env           string   `json:"env"`
	Flag          string   `json:"flag"`
}

type StopChallengeContainerReq struct {
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
	ContainerId   string `json:"containerId"`
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

func (ac *ApiClient) CallStartChallengeContainerApi(request interface{}) (*sdkmodel.StartChallengeContainerModel, error) {
	res, err := ac.CallApi("/competition/startChallengeContainer", "POST", request)
	if err != nil {
		return nil, err
	}
	var startChallengeContainerResp sdkmodel.StartChallengeContainerModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &startChallengeContainerResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got start challenge container resp: %v", startChallengeContainerResp)
	return &startChallengeContainerResp, nil
}

func (ac *ApiClient) CallStopChallengeContainerApi(request interface{}) (*sdkmodel.StopChallengeContainerModel, error) {
	res, err := ac.CallApi("/competition/stopChallengeContainer", "POST", request)
	if err != nil {
		return nil, err
	}
	var stopChallengeContainerResp sdkmodel.StopChallengeContainerModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &stopChallengeContainerResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got stop challenge container resp: %v", stopChallengeContainerResp)
	return &stopChallengeContainerResp, nil
}
