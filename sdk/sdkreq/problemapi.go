package sdkreq

import (
	"encoding/json"
	"mariosdk/sdk/sdklog"
	"mariosdk/sdk/sdkmodel"
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
