package mariosdk

import (
	"github.com/FeverKing/mariosdk/sdk/sdkclient"
	"github.com/FeverKing/mariosdk/sdk/sdkreq"
	"testing"
	"time"
)

func TestDefaultClient_Auth(t *testing.T) {

	client := sdkclient.NewClient()
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}

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
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")
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

func TestGetAuthToken(t *testing.T) {

	client := sdkclient.NewClient()
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	res, err := client.GetAuthToken()
	if err != nil {
		t.Errorf("failed: %v", err)
	}
	t.Logf("result: %v", res)
}

func TestReAuth(t *testing.T) {
	client := sdkclient.NewClient()
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")

	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}

	for {
		res, err := client.GetAuthToken()
		if err != nil {
			t.Errorf("failed: %v", err)
		}
		t.Logf("result: %v", res)

		time.Sleep(30 * time.Second)
	}
}
