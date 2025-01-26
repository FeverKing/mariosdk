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
		UserIds:       []string{""},
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

func TestGetCompetitionSetting(t *testing.T) {
	client := sdkclient.NewClient()
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	req := &sdkreq.GetCompetitionSettingReq{
		CompetitionId: "",
	}
	res, err := client.GetCompetitionSetting(req)
	if err != nil {
		t.Errorf("GetCompetitionSetting() failed: %v", err)
	}
	t.Logf("GetCompetitionSetting() result: %v", res)
}

func TestGetCompetitionAllIdentities(t *testing.T) {
	client := sdkclient.NewClient()
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	req := &sdkreq.GetCompetitionAllIdentitiesReq{
		CompetitionId: "",
	}
	res, err := client.GetCompetitionAllIdentities(req)
	if err != nil {
		t.Errorf("GetCompetitionAllIdentities() failed: %v", err)
	}
	t.Logf("GetCompetitionAllIdentities() result: %v", res)
}

func TestGetCompetitionAllTeams(t *testing.T) {
	client := sdkclient.NewClient()
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	req := &sdkreq.GetCompetitionAllTeamsReq{
		CompetitionId: "",
	}
	res, err := client.GetCompetitionAllTeams(req)
	if err != nil {
		t.Errorf("GetCompetitionAllTeams() failed: %v", err)
	}
	t.Logf("GetCompetitionAllTeams() result: %v", res)
}

func TestGetCompetitionAllUsers(t *testing.T) {
	client := sdkclient.NewClient()
	client.Config.SetAccessKey("")
	client.Config.SetSecretKey("")
	client.Config.AddEndpoint("")
	err := client.Auth()
	if err != nil {
		t.Errorf("Auth() failed: %v", err)
	}
	req := &sdkreq.GetCompetitionAllUsersReq{
		CompetitionId: "",
	}
	res, err := client.GetCompetitionAllUsers(req)
	if err != nil {
		t.Errorf("GetCompetitionAllUsers() failed: %v", err)
	}
	t.Logf("GetCompetitionAllUsers() result: %v", res)
}
