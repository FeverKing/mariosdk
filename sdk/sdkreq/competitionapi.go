package sdkreq

import (
	"encoding/json"
	"github.com/FeverKing/mariosdk/sdk/sdklog"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
)

type GetUserInfoForCompetitionReq struct {
	CompetitionId string   `json:"competitionId"`
	SecretKey     string   `json:"secretKey"`
	UserIds       []string `json:"userIds"`
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

type CheckTmpLoginVerifyTokenReq struct {
	AuthType      int    `json:"authType"`
	Token         string `json:"token"`
	CompetitionId string `json:"competitionId,omitempty"`
}

type GetCompetitionSettingReq struct {
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
}

type GetCompetitionAllIdentitiesReq struct {
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
}

type GetCompetitionAllTeamsReq struct {
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
}

type GetCompetitionAllUsersReq struct {
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
}

type GetCompetitionTemplateReq struct {
	TemplateId string `json:"templateId"`
}

type CheckCompetitionAWDPReq struct {
	ContainerId   string `json:"containerId"`
	CheckFileUrl  string `json:"checkFileUrl"`
	CompetitionId string `json:"competitionId"`
	SecretKey     string `json:"secretKey"`
}

type AwdpPatchApplyReq struct {
	UserFilePath  string `json:"userFilePath"` //minio中存储的用戶上傳的tar文件路径
	ContainerId   string `json:"containerId"`
	CheckCommand  string `json:"checkCommand"`  // 检查命令
	CheckFilePath string `json:"checkFilePath"` // 检查文件的URL,應該是指向minio的url吧
}

type UploadCompetitionScoreRequestUserCell struct {
	UserId          string `json:"userId"`          // 用户的唯一标识符
	TotalScore      int64  `json:"totalScore"`      // 用户的总分数
	TotalSolved     int64  `json:"totalSolved"`     // 用户的总解题数
	TotalFirstBlood int64  `json:"totalFirstBlood"` // 用户的总首解数
	LastSolvedTime  uint64 `json:"lastSolvedTime"`  // 最后解题时间
	DetailStats     string `json:"detailStats"`     // 用户的详细统计信息
}

type UploadCompetitionScoreRequestTeamCell struct {
	TeamId          string                                  `json:"teamId"`          // 队伍的唯一标识符
	TotalScore      int64                                   `json:"totalScore"`      // 队伍的总分数
	TotalSolved     int64                                   `json:"totalSolved"`     // 队伍的总解题数
	TotalFirstBlood int64                                   `json:"totalFirstBlood"` // 队伍的总首解数
	LastSolvedTime  uint64                                  `json:"lastSolvedTime"`  // 最后解题时间
	IsBanned        bool                                    `json:"isBanned"`        // 队伍是否被封禁
	IsAk            bool                                    `json:"isAk"`            // 队伍是否为ak
	Userstats       []UploadCompetitionScoreRequestUserCell `json:"userStats"`       // 队伍成员的统计信息
}

type UploadCompetitionScoreRequest struct {
	CompetitionId    string                                  `json:"competitionId"`    // 比赛的唯一标识符
	SecretKey        string                                  `json:"secretKey"`        // 验证身份的密钥
	CompetitionScore []UploadCompetitionScoreRequestTeamCell `json:"competitionScore"` // 比赛的得分信息
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

func (ac *ApiClient) CallCheckTmpLoginVerifyTokenApi(request interface{}) (*sdkmodel.CheckTmpLoginVerifyTokenModel, error) {
	res, err := ac.CallApi("/user/checkTmpLoginVerifyToken", "POST", request)
	if err != nil {
		return nil, err
	}
	var checkTmpLoginVerifyTokenResp sdkmodel.CheckTmpLoginVerifyTokenModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &checkTmpLoginVerifyTokenResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got check tmp login verify Token resp: %v", checkTmpLoginVerifyTokenResp)
	return &checkTmpLoginVerifyTokenResp, nil
}

func (ac *ApiClient) CallGetCompetitionSettingApi(request interface{}) ([]byte, error) {
	res, err := ac.CallApi("/competition/getCompetitionSetting", "POST", request)
	if err != nil {
		return nil, err
	}

	sdklog.Infof("got get competition setting resp")
	return ConvertInterfaceToJson(res), nil
}

func (ac *ApiClient) CallGetCompetitionAllIdentitiesApi(request interface{}) ([]byte, error) {
	res, err := ac.CallApi("/competition/getCompetitionAllIdentities", "POST", request)
	if err != nil {
		return nil, err
	}

	sdklog.Infof("got get competition all identities resp")
	return ConvertInterfaceToJson(res), nil
}

func (ac *ApiClient) CallGetCompetitionAllTeamsApi(request interface{}) ([]byte, error) {
	res, err := ac.CallApi("/competition/getCompetitionAllTeams", "POST", request)
	if err != nil {
		return nil, err
	}

	sdklog.Infof("got get competition all teams resp")
	return ConvertInterfaceToJson(res), nil
}

func (ac *ApiClient) CallGetCompetitionAllUsersApi(request interface{}) ([]byte, error) {
	res, err := ac.CallApi("/competition/getCompetitionAllUsers", "POST", request)
	if err != nil {
		return nil, err
	}

	sdklog.Infof("got get competition all users resp")
	return ConvertInterfaceToJson(res), nil
}

func (ac *ApiClient) CallGetCompetitionTemplateApi(request interface{}) ([]byte, error) {
	res, err := ac.CallApi("/competition/getCompetitionTemplate", "POST", request)
	if err != nil {
		return nil, err
	}

	sdklog.Infof("got get competition template resp")
	return ConvertInterfaceToJson(res), nil
}

func (ac *ApiClient) CallCheckCompetitionAWDPApi(request interface{}) (*sdkmodel.CheckCompetitionAWDPModel, error) {
	res, err := ac.CallApi("/competition/checkAWDP", "POST", request)
	if err != nil {
		return nil, err
	}
	var checkCompetitionAWDPResp sdkmodel.CheckCompetitionAWDPModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &checkCompetitionAWDPResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got check competition AWDP resp: %v", checkCompetitionAWDPResp)
	return &checkCompetitionAWDPResp, nil
}

func (ac *ApiClient) CallAwdpPatchApi(request interface{}) (*sdkmodel.AwdpPatchApplyModel, error) {
	res, err := ac.CallApi("/challenge/awdpPatchApply", "POST", request)
	if err != nil {
		return nil, err
	}
	var awdpPatchApplyResp sdkmodel.AwdpPatchApplyModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &awdpPatchApplyResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got awdp patch apply resp: %v", awdpPatchApplyResp)
	return &awdpPatchApplyResp, nil
}

func (ac *ApiClient) CallUploadCompetitionScoreApi(request interface{}) (*sdkmodel.UploadCompetitionScoreModel, error) {
	res, err := ac.CallApi("/competition/uploadCompetitionScore", "POST", request)
	if err != nil {
		return nil, err
	}
	var uploadCompetitionScoreResp sdkmodel.UploadCompetitionScoreModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &uploadCompetitionScoreResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got upload competition score resp: %v", uploadCompetitionScoreResp)
	return &uploadCompetitionScoreResp, nil
}
