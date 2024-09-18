package mariosdk

import (
	"mariosdk/sdk/sdkclient"
	"mariosdk/sdk/sdkreq"
	"testing"
)

func TestDefaultClient_Auth(t *testing.T) {

	client := sdkclient.NewClient()
	client.Config.SetAccessKey("Z9r1DyumUwpUYaNg")
	client.Config.SetSecretKey("qoTYWJM88pBb8H-2qDTI2ayqswFkKYQj")
	client.Config.AddEndpoint("https://mario-syclover.geesec.com/api")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	//	res, err := client.GetBatchUserInfo([]string{"1811603579241238528"})
	//	if err != nil {
	//		t.Errorf("GetBatchUserInfo() failed: %v", err)
	//	}
	//	t.Logf("GetBatchUserInfo() result: %v", res)

	req := &sdkreq.SearchPublicProblemReq{
		Page: struct {
			Page int `json:"page"`
			Size int `json:"size"`
		}(struct {
			Page int
			Size int
		}{Page: 2, Size: 17}),
	}
	res, err := client.SearchPublicProblem(req)
	if err != nil {
		t.Errorf("SearchPublicProblem() failed: %v", err)
	}
	t.Logf("SearchPublicProblem() result: %v", res)
}

func TestDefaultClient_GetUserInfoForCompetition(t *testing.T) {
	client := sdkclient.NewClient()
	client.Config.SetAccessKey("Z9r1DyumUwpUYaNg")
	client.Config.SetSecretKey("qoTYWJM88pBb8H-2qDTI2ayqswFkKYQj")
	client.Config.AddEndpoint("https://mario.test.geesec.com/api")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	req := &sdkreq.GetUserInfoForCompetitionReq{
		CompetitionId: "",
		SecretKey:     "",
		UserId:        "",
	}
	res, err := client.GetUserInfoForCompetition(req)
	if err != nil {
		t.Errorf("GetUserInfoForCompetition() failed: %v", err)
	}
	t.Logf("GetUserInfoForCompetition() result: %v", res)
}
