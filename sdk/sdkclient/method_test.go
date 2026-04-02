package sdkclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/FeverKing/mariosdk/sdk/sdkreq"
)

type fakeRequester struct {
	lastReq  *http.Request
	respBody string
	err      error
}

func (f *fakeRequester) Do(req *http.Request) (*http.Response, error) {
	f.lastReq = req
	if f.err != nil {
		return nil, f.err
	}
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(f.respBody)),
		Header:     make(http.Header),
	}, nil
}

func newAuthedClient(requester sdkreq.Requester) *DefaultClient {
	client := NewClient()
	client.Config.SetAccessKey("ak")
	client.Config.SetSecretKey("sk")
	client.Config.AddEndpoint("http://sdk.test")
	client.apiClient = sdkreq.NewApiClient(requester, "http://sdk.test")
	client.apiClient.Token = "token-x"
	client.apiClient.TokenExpiry = time.Now().Add(time.Hour)
	return client
}

func TestApiClientCallApiKeepsInjectedRequester(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"result":true}}`}
	apiClient := sdkreq.NewApiClient(requester, "http://sdk.test")

	res, err := apiClient.CallApi("/competition/pullCompetitionSnapshotForAdmin", http.MethodPost, map[string]string{"competitionId": "42"})
	if err != nil {
		t.Fatalf("CallApi() error = %v", err)
	}
	if res == nil {
		t.Fatal("CallApi() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/competition/pullCompetitionSnapshotForAdmin" {
		t.Fatalf("request URL = %s", got)
	}
}

func TestPullCompetitionSnapshotForAdmin(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"result":true,"message":"done"}}`}
	client := newAuthedClient(requester)

	resp, err := client.PullCompetitionSnapshotForAdmin(&sdkreq.PullCompetitionSnapshotForAdminRequest{CompetitionId: "9001"})
	if err != nil {
		t.Fatalf("PullCompetitionSnapshotForAdmin() error = %v", err)
	}
	if resp == nil || !resp.Result {
		t.Fatalf("PullCompetitionSnapshotForAdmin() resp = %#v", resp)
	}
	if resp.Message != "done" {
		t.Fatalf("resp.Message = %s", resp.Message)
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if requester.lastReq.Method != http.MethodPost {
		t.Fatalf("request method = %s", requester.lastReq.Method)
	}
}

func TestStartChallengeContainerCarriesCompetitionContext(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"address":["http://x"],"containerId":"c-1","restTime":3600}}`}
	client := newAuthedClient(requester)

	_, err := client.StartChallengeContainer(&sdkreq.StartChallengeContainerReq{
		CompetitionId:        "7001",
		SecretKey:            "sec",
		CompetitionProblemId: "cp-1",
		ProblemId:            "p-7",
		BundleId:             "b-6",
		TeamId:               "t-9",
		UserId:               "u-8",
		ContainerId:          "c-1",
		DockerImage:          "img:test",
	})
	if err != nil {
		t.Fatalf("StartChallengeContainer() error = %v", err)
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	body, readErr := io.ReadAll(requester.lastReq.Body)
	if readErr != nil {
		t.Fatalf("ReadAll() error = %v", readErr)
	}
	bodyStr := string(body)
	for _, want := range []string{
		`"competitionId":"7001"`,
		`"secretKey":"sec"`,
		`"competitionProblemId":"cp-1"`,
		`"problemId":"p-7"`,
		`"bundleId":"b-6"`,
		`"teamId":"t-9"`,
		`"userId":"u-8"`,
		`"containerId":"c-1"`,
		`"dockerImage":"img:test"`,
	} {
		if !bytes.Contains(body, []byte(want)) {
			t.Fatalf("request body missing %s: %s", want, bodyStr)
		}
	}
}

func TestSubmitAwdpPatch(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"patch":{"patchId":"11","userFilePath":"upload/patch.tar","status":"failed","message":"checker rejected","submittedAt":1000,"finishedAt":1010,"durationSeconds":10}}}`}
	client := newAuthedClient(requester)

	resp, err := client.SubmitAwdpPatch(&sdkreq.SubmitAwdpPatchReq{
		ProblemId:    "9001",
		UserFilePath: "upload/patch.tar",
	})
	if err != nil {
		t.Fatalf("SubmitAwdpPatch() error = %v", err)
	}
	if resp == nil {
		t.Fatal("SubmitAwdpPatch() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/problem/awdp/submitPatch" {
		t.Fatalf("request URL = %s", got)
	}
	body, readErr := io.ReadAll(requester.lastReq.Body)
	if readErr != nil {
		t.Fatalf("ReadAll() error = %v", readErr)
	}
	bodyStr := string(body)
	for _, want := range []string{`"problemId":"9001"`, `"userFilePath":"upload/patch.tar"`} {
		if !bytes.Contains(body, []byte(want)) {
			t.Fatalf("request body missing %s: %s", want, bodyStr)
		}
	}
	if resp.Patch.PatchId != "11" || resp.Patch.Status != "failed" || resp.Patch.Message != "checker rejected" {
		t.Fatalf("response = %#v", resp)
	}
}

func TestUploadCompetitionScoreReauthsAndSerializesScorePayload(t *testing.T) {
	var paths []string
	var uploadAuth string
	var uploadPayload map[string]any

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paths = append(paths, r.URL.Path)
		switch r.URL.Path {
		case "/sdk/getSdkToken":
			_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"token":"fresh-token"}}`))
		case "/competition/uploadCompetitionScore":
			uploadAuth = r.Header.Get("Authorization")
			body, err := io.ReadAll(r.Body)
			if err != nil {
				t.Fatalf("ReadAll() error = %v", err)
			}
			if err := json.Unmarshal(body, &uploadPayload); err != nil {
				t.Fatalf("json.Unmarshal() error = %v", err)
			}
			_, _ = w.Write([]byte(`{"code":200,"msg":"ok","data":{"success":true,"message":"uploaded"}}`))
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

	_, err := client.UploadCompetitionScore(&sdkreq.UploadCompetitionScoreRequest{
		CompetitionId: "8001",
		SecretKey:     "sec",
		CompetitionScore: []sdkreq.UploadCompetitionScoreRequestTeamCell{
			{
				TeamId:          "team-1",
				TotalScore:      123,
				TotalSolved:     4,
				TotalFirstBlood: 1,
				LastSolvedTime:  1700000001,
				IsAk:            true,
				Userstats: []sdkreq.UploadCompetitionScoreRequestUserCell{
					{
						UserId:          "user-1",
						TotalScore:      77,
						TotalSolved:     3,
						TotalFirstBlood: 1,
						LastSolvedTime:  1700000002,
						DetailStats:     `{"direction":"web","solved":3}`,
					},
				},
			},
		},
		PostCompetitionSnapshotPath:      "/snapshot/8001.json",
		PostCompetitionSnapshotUpdatedAt: 1700000000,
	})
	if err != nil {
		t.Fatalf("UploadCompetitionScore() error = %v", err)
	}
	if len(paths) != 2 || paths[0] != "/sdk/getSdkToken" || paths[1] != "/competition/uploadCompetitionScore" {
		t.Fatalf("request paths = %#v", paths)
	}
	if uploadAuth != "fresh-token" {
		t.Fatalf("Authorization = %s", uploadAuth)
	}
	if got := uploadPayload["competitionId"]; got != "8001" {
		t.Fatalf("competitionId = %#v", got)
	}
	if got := uploadPayload["secretKey"]; got != "sec" {
		t.Fatalf("secretKey = %#v", got)
	}
	if got := uploadPayload["postCompetitionSnapshotPath"]; got != "/snapshot/8001.json" {
		t.Fatalf("postCompetitionSnapshotPath = %#v", got)
	}
	if got := uploadPayload["postCompetitionSnapshotUpdatedAt"]; got != float64(1700000000) {
		t.Fatalf("postCompetitionSnapshotUpdatedAt = %#v", got)
	}
	competitionScore, ok := uploadPayload["competitionScore"].([]any)
	if !ok || len(competitionScore) != 1 {
		t.Fatalf("competitionScore = %#v", uploadPayload["competitionScore"])
	}
	teamCell, ok := competitionScore[0].(map[string]any)
	if !ok {
		t.Fatalf("teamCell = %#v", competitionScore[0])
	}
	if got := teamCell["teamId"]; got != "team-1" {
		t.Fatalf("teamId = %#v", got)
	}
	if got := teamCell["userStats"]; got == nil {
		t.Fatalf("userStats = %#v", got)
	}
	userStats, ok := teamCell["userStats"].([]any)
	if !ok || len(userStats) != 1 {
		t.Fatalf("userStats = %#v", teamCell["userStats"])
	}
	userCell, ok := userStats[0].(map[string]any)
	if !ok {
		t.Fatalf("userCell = %#v", userStats[0])
	}
	if got := userCell["userId"]; got != "user-1" {
		t.Fatalf("userId = %#v", got)
	}
	if got := userCell["detailStats"]; got != `{"direction":"web","solved":3}` {
		t.Fatalf("detailStats = %#v", got)
	}
}
