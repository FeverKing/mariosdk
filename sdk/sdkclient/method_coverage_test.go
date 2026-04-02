package sdkclient

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/FeverKing/mariosdk/sdk/sdkreq"
)

func TestGetAuthTokenBootstrapsAuthWhenAPIClientNil(t *testing.T) {
	authCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/sdk/getSdkToken":
			authCalls++
			_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"token":"bootstrapped-token"}}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	client := NewClient()
	client.Config.SetAccessKey("ak")
	client.Config.SetSecretKey("sk")
	client.Config.AddEndpoint(server.URL)

	token, err := client.GetAuthToken()
	if err != nil {
		t.Fatalf("GetAuthToken() error = %v", err)
	}
	if token != "bootstrapped-token" {
		t.Fatalf("token = %s", token)
	}
	if authCalls != 1 {
		t.Fatalf("authCalls = %d", authCalls)
	}
}

func TestAuthRejectsMissingEndpointInsteadOfPanicking(t *testing.T) {
	client := NewClient()
	client.Config.SetAccessKey("ak")
	client.Config.SetSecretKey("sk")

	if err := client.Auth(); err == nil {
		t.Fatal("Auth() error = nil")
	}
}

func TestAuthSuccessAndGetAuthTokenReauth(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/sdk/getSdkToken" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"token":"token-from-auth"}}`))
	}))
	defer server.Close()

	client := NewClient()
	client.Config.SetAccessKey("ak")
	client.Config.SetSecretKey("sk")
	client.Config.AddEndpoint(server.URL)
	client.apiClient = sdkreq.NewApiClient(nil, server.URL)
	client.apiClient.TokenExpiry = time.Now().Add(-time.Minute)

	token, err := client.GetAuthToken()
	if err != nil {
		t.Fatalf("GetAuthToken() error = %v", err)
	}
	if token != "token-from-auth" {
		t.Fatalf("token = %s", token)
	}
}

func TestAuthFailure(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":500,"msg":"auth bad","data":null}`))
	}))
	defer server.Close()

	client := NewClient()
	client.Config.SetAccessKey("ak")
	client.Config.SetSecretKey("sk")
	client.Config.AddEndpoint(server.URL)

	if err := client.Auth(); err == nil {
		t.Fatal("Auth() error = nil")
	}
}

func TestUploadCompetitionScorePropagatesEnsureAuthFailure(t *testing.T) {
	authCalls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/sdk/getSdkToken":
			authCalls++
			_, _ = w.Write([]byte(`{"code":500,"msg":"reauth failed","data":null}`))
		case "/competition/uploadCompetitionScore":
			t.Fatal("UploadCompetitionScore() should not call upload endpoint when reauth fails")
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	client := NewClient()
	client.Config.SetAccessKey("ak")
	client.Config.SetSecretKey("sk")
	client.Config.AddEndpoint(server.URL)
	client.apiClient = sdkreq.NewApiClient(nil, server.URL)
	client.apiClient.Token = "expired-token"
	client.apiClient.TokenExpiry = time.Now().Add(-time.Minute)

	_, err := client.UploadCompetitionScore(&sdkreq.UploadCompetitionScoreRequest{})
	if err == nil {
		t.Fatal("UploadCompetitionScore() error = nil")
	}
	if err.Error() != "reauth failed" {
		t.Fatalf("UploadCompetitionScore() error = %v", err)
	}
	if authCalls != 1 {
		t.Fatalf("authCalls = %d", authCalls)
	}
}

func TestGetAuthTokenReturnsExistingTokenWhenNotExpired(t *testing.T) {
	client := newAuthedClient(&fakeRequester{respBody: `{"code":200,"msg":"ok","data":{}}`})
	client.apiClient.Token = "token-still-valid"
	client.apiClient.TokenExpiry = time.Now().Add(time.Hour)

	token, err := client.GetAuthToken()
	if err != nil {
		t.Fatalf("GetAuthToken() error = %v", err)
	}
	if token != "token-still-valid" {
		t.Fatalf("token = %s", token)
	}
}

func TestGetAuthTokenInitializesMissingApiClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/sdk/getSdkToken" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"token":"token-from-nil-client"}}`))
	}))
	defer server.Close()

	client := NewClient()
	client.Config.SetAccessKey("ak")
	client.Config.SetSecretKey("sk")
	client.Config.AddEndpoint(server.URL)

	token, err := client.GetAuthToken()
	if err != nil {
		t.Fatalf("GetAuthToken() error = %v", err)
	}
	if token != "token-from-nil-client" {
		t.Fatalf("token = %s", token)
	}
	if client.apiClient == nil {
		t.Fatal("apiClient = nil")
	}
}

func TestGetCompetitionSettingInitializesMissingApiClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/competition/getCompetitionSetting" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"mode":"awdp"}}`))
	}))
	defer server.Close()

	client := NewClient()
	client.Config.AddEndpoint(server.URL)

	resp, err := client.GetCompetitionSetting(&sdkreq.GetCompetitionSettingReq{CompetitionId: "1", SecretKey: "sec"})
	if err != nil {
		t.Fatalf("GetCompetitionSetting() error = %v", err)
	}
	if string(resp) != `{"mode":"awdp"}` {
		t.Fatalf("resp = %s", string(resp))
	}
	if client.apiClient == nil {
		t.Fatal("apiClient = nil")
	}
}

func TestSdkClientAdditionalWrapperSuccess(t *testing.T) {
	cases := []struct {
		name string
		url  string
		call func(*DefaultClient) (any, error)
	}{
		{
			name: "GetBatchUserInfo",
			url:  "http://sdk.test/user/benchGetUserBase",
			call: func(c *DefaultClient) (any, error) { return c.GetBatchUserInfo([]string{"1"}) },
		},
		{
			name: "GetUserInfoForCompetition",
			url:  "http://sdk.test/competition/getUserInfoForCompetition",
			call: func(c *DefaultClient) (any, error) {
				return c.GetUserInfoForCompetition(&sdkreq.GetUserInfoForCompetitionReq{})
			},
		},
		{
			name: "StopChallengeContainer",
			url:  "http://sdk.test/competition/stopChallengeContainer",
			call: func(c *DefaultClient) (any, error) {
				return c.StopChallengeContainer(&sdkreq.StopChallengeContainerReq{})
			},
		},
		{
			name: "CheckTmpLoginVerifyToken",
			url:  "http://sdk.test/user/checkTmpLoginVerifyToken",
			call: func(c *DefaultClient) (any, error) {
				return c.CheckTmpLoginVerifyToken(&sdkreq.CheckTmpLoginVerifyTokenReq{})
			},
		},
		{
			name: "GetCompetitionSetting",
			url:  "http://sdk.test/competition/getCompetitionSetting",
			call: func(c *DefaultClient) (any, error) {
				return c.GetCompetitionSetting(&sdkreq.GetCompetitionSettingReq{})
			},
		},
		{
			name: "GetCompetitionAllIdentities",
			url:  "http://sdk.test/competition/getCompetitionAllIdentities",
			call: func(c *DefaultClient) (any, error) {
				return c.GetCompetitionAllIdentities(&sdkreq.GetCompetitionAllIdentitiesReq{})
			},
		},
		{
			name: "GetCompetitionAllTeams",
			url:  "http://sdk.test/competition/getCompetitionAllTeams",
			call: func(c *DefaultClient) (any, error) {
				return c.GetCompetitionAllTeams(&sdkreq.GetCompetitionAllTeamsReq{})
			},
		},
		{
			name: "GetCompetitionAllUsers",
			url:  "http://sdk.test/competition/getCompetitionAllUsers",
			call: func(c *DefaultClient) (any, error) {
				return c.GetCompetitionAllUsers(&sdkreq.GetCompetitionAllUsersReq{})
			},
		},
		{
			name: "GetCompetitionTemplate",
			url:  "http://sdk.test/competition/getCompetitionTemplate",
			call: func(c *DefaultClient) (any, error) {
				return c.GetCompetitionTemplate(&sdkreq.GetCompetitionTemplateReq{})
			},
		},
		{
			name: "CheckCompetitionAWDP",
			url:  "http://sdk.test/competition/checkAWDP",
			call: func(c *DefaultClient) (any, error) { return c.CheckCompetitionAWDP(&sdkreq.CheckCompetitionAWDPReq{}) },
		},
		{
			name: "AwdpPatchApi",
			url:  "http://sdk.test/challenge/awdpPatchApply",
			call: func(c *DefaultClient) (any, error) { return c.AwdpPatchApi(&sdkreq.AwdpPatchApplyReq{}) },
		},
		{
			name: "SubmitAwdpPatch",
			url:  "http://sdk.test/problem/awdp/submitPatch",
			call: func(c *DefaultClient) (any, error) { return c.SubmitAwdpPatch(&sdkreq.SubmitAwdpPatchReq{}) },
		},
		{
			name: "GetTeamInfoForCompetition",
			url:  "http://sdk.test/competition/getTeamInfoForCompetition",
			call: func(c *DefaultClient) (any, error) {
				return c.GetTeamInfoForCompetition(&sdkreq.GetTeamInfoForCompetitionRequest{})
			},
		},
		{
			name: "GetAuthorizedProblemBanks",
			url:  "http://sdk.test/competition/problemBank/getAuthorizedProblemBanks",
			call: func(c *DefaultClient) (any, error) {
				return c.GetAuthorizedProblemBanks(&sdkreq.GetAuthorizedProblemBanksReq{})
			},
		},
		{
			name: "GetProblemBankForCompetition",
			url:  "http://sdk.test/competition/problemBank/getProblemBankForCompetition",
			call: func(c *DefaultClient) (any, error) {
				return c.GetProblemBankForCompetition(&sdkreq.GetProblemBankForCompetitionReq{})
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{}}`}
			client := newAuthedClient(requester)

			resp, err := tc.call(client)
			if err != nil {
				t.Fatalf("%s() error = %v", tc.name, err)
			}
			if resp == nil {
				t.Fatalf("%s() returned nil response", tc.name)
			}
			if requester.lastReq == nil {
				t.Fatalf("%s() did not issue request", tc.name)
			}
			if got := requester.lastReq.URL.String(); got != tc.url {
				t.Fatalf("%s() request URL = %s", tc.name, got)
			}
		})
	}
}

func TestSdkClientWrapperErrorCoverage(t *testing.T) {
	reqErr := errors.New("request failed")
	cases := []struct {
		name string
		call func(*DefaultClient) error
	}{
		{name: "GetBatchUserInfo", call: func(c *DefaultClient) error { _, err := c.GetBatchUserInfo([]string{"1"}); return err }},
		{name: "GetUserInfoForCompetition", call: func(c *DefaultClient) error {
			_, err := c.GetUserInfoForCompetition(&sdkreq.GetUserInfoForCompetitionReq{})
			return err
		}},
		{name: "StartChallengeContainer", call: func(c *DefaultClient) error {
			_, err := c.StartChallengeContainer(&sdkreq.StartChallengeContainerReq{})
			return err
		}},
		{name: "SubmitAwdpPatch", call: func(c *DefaultClient) error {
			_, err := c.SubmitAwdpPatch(&sdkreq.SubmitAwdpPatchReq{})
			return err
		}},
		{name: "StopChallengeContainer", call: func(c *DefaultClient) error {
			_, err := c.StopChallengeContainer(&sdkreq.StopChallengeContainerReq{})
			return err
		}},
		{name: "CheckTmpLoginVerifyToken", call: func(c *DefaultClient) error {
			_, err := c.CheckTmpLoginVerifyToken(&sdkreq.CheckTmpLoginVerifyTokenReq{})
			return err
		}},
		{name: "GetCompetitionSetting", call: func(c *DefaultClient) error {
			_, err := c.GetCompetitionSetting(&sdkreq.GetCompetitionSettingReq{})
			return err
		}},
		{name: "GetCompetitionAllIdentities", call: func(c *DefaultClient) error {
			_, err := c.GetCompetitionAllIdentities(&sdkreq.GetCompetitionAllIdentitiesReq{})
			return err
		}},
		{name: "GetCompetitionAllTeams", call: func(c *DefaultClient) error {
			_, err := c.GetCompetitionAllTeams(&sdkreq.GetCompetitionAllTeamsReq{})
			return err
		}},
		{name: "GetCompetitionAllUsers", call: func(c *DefaultClient) error {
			_, err := c.GetCompetitionAllUsers(&sdkreq.GetCompetitionAllUsersReq{})
			return err
		}},
		{name: "GetCompetitionTemplate", call: func(c *DefaultClient) error {
			_, err := c.GetCompetitionTemplate(&sdkreq.GetCompetitionTemplateReq{})
			return err
		}},
		{name: "CheckCompetitionAWDP", call: func(c *DefaultClient) error {
			_, err := c.CheckCompetitionAWDP(&sdkreq.CheckCompetitionAWDPReq{})
			return err
		}},
		{name: "AwdpPatchApi", call: func(c *DefaultClient) error { _, err := c.AwdpPatchApi(&sdkreq.AwdpPatchApplyReq{}); return err }},
		{name: "SubmitAwdpPatch", call: func(c *DefaultClient) error {
			_, err := c.SubmitAwdpPatch(&sdkreq.SubmitAwdpPatchReq{})
			return err
		}},
		{name: "UploadCompetitionScore", call: func(c *DefaultClient) error {
			_, err := c.UploadCompetitionScore(&sdkreq.UploadCompetitionScoreRequest{})
			return err
		}},
		{name: "GetTeamInfoForCompetition", call: func(c *DefaultClient) error {
			_, err := c.GetTeamInfoForCompetition(&sdkreq.GetTeamInfoForCompetitionRequest{})
			return err
		}},
		{name: "PullCompetitionSnapshotForAdmin", call: func(c *DefaultClient) error {
			_, err := c.PullCompetitionSnapshotForAdmin(&sdkreq.PullCompetitionSnapshotForAdminRequest{})
			return err
		}},
		{name: "GetAuthorizedProblemBanks", call: func(c *DefaultClient) error {
			_, err := c.GetAuthorizedProblemBanks(&sdkreq.GetAuthorizedProblemBanksReq{})
			return err
		}},
		{name: "GetProblemBankForCompetition", call: func(c *DefaultClient) error {
			_, err := c.GetProblemBankForCompetition(&sdkreq.GetProblemBankForCompetitionReq{})
			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := newAuthedClient(&fakeRequester{err: reqErr})
			err := tc.call(client)
			if !errors.Is(err, reqErr) {
				t.Fatalf("%s() error = %v", tc.name, err)
			}
		})
	}
}

func TestSdkClientEnsureAuthFailureCoverage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`{"code":500,"msg":"reauth failed","data":null}`))
	}))
	defer server.Close()

	newExpiredClient := func() *DefaultClient {
		client := NewClient()
		client.Config.SetAccessKey("ak")
		client.Config.SetSecretKey("sk")
		client.Config.AddEndpoint(server.URL)
		client.apiClient = sdkreq.NewApiClient(nil, server.URL)
		client.apiClient.Token = "expired-token"
		client.apiClient.TokenExpiry = time.Now().Add(-time.Minute)
		return client
	}

	cases := []struct {
		name string
		call func(*DefaultClient) error
	}{
		{name: "GetBatchUserInfo", call: func(c *DefaultClient) error { _, err := c.GetBatchUserInfo([]string{"1"}); return err }},
		{name: "GetUserInfoForCompetition", call: func(c *DefaultClient) error {
			_, err := c.GetUserInfoForCompetition(&sdkreq.GetUserInfoForCompetitionReq{})
			return err
		}},
		{name: "StartChallengeContainer", call: func(c *DefaultClient) error {
			_, err := c.StartChallengeContainer(&sdkreq.StartChallengeContainerReq{})
			return err
		}},
		{name: "SubmitAwdpPatch", call: func(c *DefaultClient) error {
			_, err := c.SubmitAwdpPatch(&sdkreq.SubmitAwdpPatchReq{})
			return err
		}},
		{name: "StopChallengeContainer", call: func(c *DefaultClient) error {
			_, err := c.StopChallengeContainer(&sdkreq.StopChallengeContainerReq{})
			return err
		}},
		{name: "CheckTmpLoginVerifyToken", call: func(c *DefaultClient) error {
			_, err := c.CheckTmpLoginVerifyToken(&sdkreq.CheckTmpLoginVerifyTokenReq{})
			return err
		}},
		{name: "GetAuthToken", call: func(c *DefaultClient) error { _, err := c.GetAuthToken(); return err }},
		{name: "GetTeamInfoForCompetition", call: func(c *DefaultClient) error {
			_, err := c.GetTeamInfoForCompetition(&sdkreq.GetTeamInfoForCompetitionRequest{})
			return err
		}},
		{name: "PullCompetitionSnapshotForAdmin", call: func(c *DefaultClient) error {
			_, err := c.PullCompetitionSnapshotForAdmin(&sdkreq.PullCompetitionSnapshotForAdminRequest{})
			return err
		}},
		{name: "GetAuthorizedProblemBanks", call: func(c *DefaultClient) error {
			_, err := c.GetAuthorizedProblemBanks(&sdkreq.GetAuthorizedProblemBanksReq{})
			return err
		}},
		{name: "GetProblemBankForCompetition", call: func(c *DefaultClient) error {
			_, err := c.GetProblemBankForCompetition(&sdkreq.GetProblemBankForCompetitionReq{})
			return err
		}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.call(newExpiredClient())
			if err == nil {
				t.Fatalf("%s() error = nil", tc.name)
			}
		})
	}
}

func TestProblemApiDecodeErrors(t *testing.T) {
	cases := []struct {
		name string
		call func(*DefaultClient) error
	}{}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := newAuthedClient(&fakeRequester{respBody: `{"code":200,"msg":"ok","data":123}`})
			if err := tc.call(client); err == nil {
				t.Fatalf("%s() error = nil", tc.name)
			}
		})
	}
}
