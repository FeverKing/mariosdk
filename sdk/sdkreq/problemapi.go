package sdkreq

import (
	"encoding/json"
	"github.com/FeverKing/mariosdk/sdk/sdklog"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
)

type SearchPublicProblemReq struct {
	Name             string   `json:"name"`
	Category         string   `json:"category"`
	Tags             []string `json:"tags"`
	ProblemType      int      `json:"problemType"`
	Difficulty       int      `json:"difficulty"`
	PublicType       int      `json:"publicType"`
	PublicIdWithType string   `json:"publicIdWithType"`
	IsSolved         int      `json:"isSolved"`
	ProblemBankId    int      `json:"problemBankId"`
	FuzzyQuery       string   `json:"fuzzyQuery"`
	IgnoreIds        []string `json:"ignoreIds"`
	Page             struct {
		Page int `json:"page"`
		Size int `json:"size"`
	} `json:"page"`
}

type GetMyProblemAnalysisReq struct{}

type GetPremiumProblemAnalysisReq struct {
	RefType string `json:"refType,omitempty"`
	RefId   string `json:"refId,omitempty"`
}

func (ac *ApiClient) CallSearchPublicProblemApi(request interface{}) (*sdkmodel.SearchPublicProblemModel, error) {
	res, err := ac.CallApi("/problem/searchPublicProblem", "POST", request)
	if err != nil {
		return nil, err
	}
	var searchPublicProblemResp sdkmodel.SearchPublicProblemModel
	err = json.Unmarshal(ConvertInterfaceToJson(res), &searchPublicProblemResp)
	if err != nil {
		return nil, err
	}
	sdklog.Infof("got search public problem resp: %v", searchPublicProblemResp)
	return &searchPublicProblemResp, nil
}

func (ac *ApiClient) CallGetMyProblemAnalysisApi(request interface{}) (*sdkmodel.GetMyProblemAnalysisModel, error) {
	res, err := ac.CallApi("/problem/getMyProblemAnalysis", "POST", request)
	if err != nil {
		return nil, err
	}
	var resp sdkmodel.GetMyProblemAnalysisModel
	if err = json.Unmarshal(ConvertInterfaceToJson(res), &resp); err != nil {
		return nil, err
	}
	sdklog.Infof("got my problem analysis resp: %v", resp)
	return &resp, nil
}

func (ac *ApiClient) CallGetPremiumProblemAnalysisApi(request interface{}) (*sdkmodel.GetPremiumProblemAnalysisModel, error) {
	res, err := ac.CallApi("/problem/getPremiumProblemAnalysis", "POST", request)
	if err != nil {
		return nil, err
	}
	var resp sdkmodel.GetPremiumProblemAnalysisModel
	if err = json.Unmarshal(ConvertInterfaceToJson(res), &resp); err != nil {
		return nil, err
	}
	sdklog.Infof("got premium problem analysis resp: %v", resp)
	return &resp, nil
}
