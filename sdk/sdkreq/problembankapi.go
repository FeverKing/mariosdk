package sdkreq

import (
	"encoding/json"
	"github.com/FeverKing/mariosdk/sdk/sdklog"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
)

// ============== 题库授权相关请求 ==============

// GetAuthorizedProblemBanksReq 获取已授权的题库列表请求
type GetAuthorizedProblemBanksReq struct {
	CompetitionId string `json:"competitionId"` // 比赛ID
	SecretKey     string `json:"secretKey"`     // 比赛SecretKey
	Page          int64  `json:"page"`          // 页码
	PageSize      int64  `json:"pageSize"`      // 每页数量
}

// GetProblemBankForCompetitionReq 获取题库详情请求
type GetProblemBankForCompetitionReq struct {
	CompetitionId string   `json:"competitionId"`       // 比赛ID
	SecretKey     string   `json:"secretKey"`           // 比赛SecretKey
	ProblemBankId string   `json:"problemBankId"`       // 题库ID
	ProblemIds    []string `json:"problemIds,optional"` // 可选，指定获取的题目ID列表
}

// ============== API调用方法 ==============

// CallGetAuthorizedProblemBanksApi 获取已授权的题库列表
func (ac *ApiClient) CallGetAuthorizedProblemBanksApi(request interface{}) (*sdkmodel.GetAuthorizedProblemBanksModel, error) {
	res, err := ac.CallApi("/competition/problemBank/getAuthorizedProblemBanks", "POST", request)
	if err != nil {
		return nil, err
	}
	var resp sdkmodel.GetAuthorizedProblemBanksModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &resp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got authorized problem banks resp: %v", resp)
	return &resp, nil
}

// CallGetProblemBankForCompetitionApi 获取题库详情(包含题目)
func (ac *ApiClient) CallGetProblemBankForCompetitionApi(request interface{}) (*sdkmodel.GetProblemBankForCompetitionModel, error) {
	res, err := ac.CallApi("/competition/problemBank/getProblemBankForCompetition", "POST", request)
	if err != nil {
		return nil, err
	}
	var resp sdkmodel.GetProblemBankForCompetitionModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &resp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got problem bank for competition resp: bank=%s, problems=%d", resp.Name, len(resp.Problems))
	return &resp, nil
}
