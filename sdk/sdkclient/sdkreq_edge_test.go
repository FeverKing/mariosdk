package sdkclient

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/FeverKing/mariosdk/sdk/sdkreq"
)

type staticResponseRequester struct {
	lastReq *http.Request
	resp    *http.Response
	err     error
}

func (s *staticResponseRequester) Do(req *http.Request) (*http.Response, error) {
	s.lastReq = req
	if s.err != nil {
		return nil, s.err
	}
	return s.resp, nil
}

type errReadCloser struct {
	readErr error
}

func (e *errReadCloser) Read(_ []byte) (int, error) {
	return 0, e.readErr
}

func (e *errReadCloser) Close() error {
	return nil
}

func TestDefaultClientSetConfig(t *testing.T) {
	client := NewClient()
	cfg := Config{AccessKey: "ak", SecretKey: "sk", Endpoints: []string{"http://sdk.test"}}

	client.setConfig(cfg)

	if !reflect.DeepEqual(client.Config, cfg) {
		t.Fatalf("Config = %#v", client.Config)
	}
}

func TestConvertInterfaceToJsonNilReturnsNil(t *testing.T) {
	if got := sdkreq.ConvertInterfaceToJson(nil); got != nil {
		t.Fatalf("ConvertInterfaceToJson(nil) = %s", string(got))
	}
}

func TestConvertInterfaceToJsonUnsupportedValueReturnsNil(t *testing.T) {
	if got := sdkreq.ConvertInterfaceToJson(map[string]interface{}{"bad": func() {}}); got != nil {
		t.Fatalf("ConvertInterfaceToJson(unsupported) = %s", string(got))
	}
}

func TestApiClientCallApiRejectsNilResponse(t *testing.T) {
	apiClient := sdkreq.NewApiClient(&staticResponseRequester{}, "http://sdk.test")

	_, err := apiClient.CallApi("/nil-response", http.MethodPost, map[string]string{"id": "1"})
	if err == nil {
		t.Fatal("CallApi() error = nil")
	}
}

func TestApiClientCallApiRejectsNilResponseBody(t *testing.T) {
	apiClient := sdkreq.NewApiClient(&staticResponseRequester{
		resp: &http.Response{StatusCode: http.StatusOK},
	}, "http://sdk.test")

	_, err := apiClient.CallApi("/nil-body", http.MethodPost, map[string]string{"id": "1"})
	if err == nil {
		t.Fatal("CallApi() error = nil")
	}
}

func TestApiClientCallApiRejectsNullBaseResponse(t *testing.T) {
	apiClient := sdkreq.NewApiClient(&fakeRequester{respBody: `null`}, "http://sdk.test")

	_, err := apiClient.CallApi("/null-response", http.MethodPost, map[string]string{"id": "1"})
	if err == nil {
		t.Fatal("CallApi() error = nil")
	}
}

func TestApiClientCallApiReturnsReadBodyError(t *testing.T) {
	readErr := errors.New("read failed")
	apiClient := sdkreq.NewApiClient(&staticResponseRequester{
		resp: &http.Response{
			StatusCode: http.StatusOK,
			Body:       &errReadCloser{readErr: readErr},
			Header:     make(http.Header),
		},
	}, "http://sdk.test")

	_, err := apiClient.CallApi("/read-error", http.MethodPost, map[string]string{"id": "1"})
	if !errors.Is(err, readErr) {
		t.Fatalf("CallApi() error = %v", err)
	}
}

func TestApiClientCallApiRejectsMarshalPayloadError(t *testing.T) {
	apiClient := sdkreq.NewApiClient(&fakeRequester{}, "http://sdk.test")

	_, err := apiClient.CallApi("/marshal-error", http.MethodPost, map[string]interface{}{"bad": func() {}})
	if err == nil {
		t.Fatal("CallApi() error = nil")
	}
}

func TestApiClientCallApiRejectsInvalidRequestMethod(t *testing.T) {
	apiClient := sdkreq.NewApiClient(&fakeRequester{}, "http://sdk.test")

	_, err := apiClient.CallApi("/bad-method", " bad method ", map[string]string{"id": "1"})
	if err == nil {
		t.Fatal("CallApi() error = nil")
	}
}

func TestApiClientCallApiWithoutPayloadUsesDefaultRequester(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Fatalf("method = %s", r.Method)
		}
		if got := r.Header.Get("Authorization"); got != "token-x" {
			t.Fatalf("Authorization = %s", got)
		}
		if got := r.Header.Get("Content-Type"); got != "" {
			t.Fatalf("Content-Type = %s", got)
		}
		if _, err := io.ReadAll(r.Body); err != nil {
			t.Fatalf("ReadAll() error = %v", err)
		}
		_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"ready":true}}`))
	}))
	defer server.Close()

	apiClient := sdkreq.NewApiClient(nil, server.URL)
	apiClient.Token = "token-x"

	res, err := apiClient.CallApi("/default-requester", http.MethodGet, nil)
	if err != nil {
		t.Fatalf("CallApi() error = %v", err)
	}
	if res == nil {
		t.Fatal("CallApi() response = nil")
	}
}

func TestApiClientCallApiRejectsInvalidJSONResponse(t *testing.T) {
	apiClient := sdkreq.NewApiClient(&fakeRequester{respBody: `{"code":200,"msg":"ok","data":`}, "http://sdk.test")

	_, err := apiClient.CallApi("/bad-json", http.MethodPost, map[string]string{"id": "1"})
	if err == nil {
		t.Fatal("CallApi() error = nil")
	}
}

func TestTypedApiWrappersRejectNullData(t *testing.T) {
	cases := []struct {
		name string
		call func(*sdkreq.ApiClient) error
	}{
		{
			name: "CallAuthApi",
			call: func(ac *sdkreq.ApiClient) error {
				return ac.CallAuthApi(&sdkreq.AuthApiRequest{AccessKey: "ak", SecretKey: "sk"})
			},
		},
		{
			name: "CallUserInfoApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallUserInfoApi(&sdkreq.BatchUserInfoRequest{Ids: []string{"1"}})
				return err
			},
		},
		{
			name: "CallGetUserInfoForCompetitionApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallGetUserInfoForCompetitionApi(&sdkreq.GetUserInfoForCompetitionReq{})
				return err
			},
		},
		{
			name: "CallStartChallengeContainerApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallStartChallengeContainerApi(&sdkreq.StartChallengeContainerReq{})
				return err
			},
		},
		{
			name: "CallStopChallengeContainerApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallStopChallengeContainerApi(&sdkreq.StopChallengeContainerReq{})
				return err
			},
		},
		{
			name: "CallCheckTmpLoginVerifyTokenApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallCheckTmpLoginVerifyTokenApi(&sdkreq.CheckTmpLoginVerifyTokenReq{})
				return err
			},
		},
		{
			name: "CallCheckCompetitionAWDPApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallCheckCompetitionAWDPApi(&sdkreq.CheckCompetitionAWDPReq{})
				return err
			},
		},
		{
			name: "CallSubmitAwdpPatchApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallSubmitAwdpPatchApi(&sdkreq.SubmitAwdpPatchReq{})
				return err
			},
		},
		{
			name: "CallAwdpPatchApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallAwdpPatchApi(&sdkreq.AwdpPatchApplyReq{})
				return err
			},
		},
		{
			name: "CallUploadCompetitionScoreApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallUploadCompetitionScoreApi(&sdkreq.UploadCompetitionScoreRequest{})
				return err
			},
		},
		{
			name: "CallGetTeamInfoForCompetitionApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallGetTeamInfoForCompetitionApi(&sdkreq.GetTeamInfoForCompetitionRequest{})
				return err
			},
		},
		{
			name: "CallPullCompetitionSnapshotForAdminApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallPullCompetitionSnapshotForAdminApi(&sdkreq.PullCompetitionSnapshotForAdminRequest{})
				return err
			},
		},
		{
			name: "CallGetAuthorizedProblemBanksApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallGetAuthorizedProblemBanksApi(&sdkreq.GetAuthorizedProblemBanksReq{})
				return err
			},
		},
		{
			name: "CallGetProblemBankForCompetitionApi",
			call: func(ac *sdkreq.ApiClient) error {
				_, err := ac.CallGetProblemBankForCompetitionApi(&sdkreq.GetProblemBankForCompetitionReq{})
				return err
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ac := sdkreq.NewApiClient(&fakeRequester{respBody: `{"code":200,"msg":"ok","data":null}`}, "http://sdk.test")

			err := tc.call(ac)
			if err == nil {
				t.Fatalf("%s() error = nil", tc.name)
			}
			if !strings.Contains(err.Error(), "unexpected end of JSON input") {
				t.Fatalf("%s() error = %v", tc.name, err)
			}
		})
	}
}
