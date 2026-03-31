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

func TestGetAwdpBundleDetail(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"bundleId":"9101","bundleName":"AWDP Bundle","problemIds":["9001","9002"],"problemCount":2}}`}
	client := newAuthedClient(requester)

	resp, err := client.GetAwdpBundleDetail(&sdkreq.GetAwdpBundleDetailReq{BundleId: "9101"})
	if err != nil {
		t.Fatalf("GetAwdpBundleDetail() error = %v", err)
	}
	if resp == nil {
		t.Fatal("GetAwdpBundleDetail() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if requester.lastReq.Method != http.MethodPost {
		t.Fatalf("request method = %s", requester.lastReq.Method)
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/competition/getAwdpBundleDetail" {
		t.Fatalf("request URL = %s", got)
	}
	body, readErr := io.ReadAll(requester.lastReq.Body)
	if readErr != nil {
		t.Fatalf("ReadAll() error = %v", readErr)
	}
	if !bytes.Contains(body, []byte(`"bundleId":"9101"`)) {
		t.Fatalf("request body = %s", string(body))
	}
	if resp.BundleId != "9101" || resp.BundleName != "AWDP Bundle" {
		t.Fatalf("response = %#v", resp)
	}
	if len(resp.ProblemIds) != 2 || resp.ProblemIds[0] != "9001" || resp.ProblemCount != 2 {
		t.Fatalf("response = %#v", resp)
	}
}

func TestGetAwdpProblemRank(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"attackSpeedRank":[{"userId":"42","userName":"alice","problemId":"9001","duration":40,"rank":1}],"defenseSpeedRank":[{"userId":"78","userName":"carol","problemId":"9001","duration":10,"rank":1}]}}`}
	client := newAuthedClient(requester)

	resp, err := client.GetAwdpProblemRank(&sdkreq.GetAwdpProblemRankReq{ProblemId: "9001"})
	if err != nil {
		t.Fatalf("GetAwdpProblemRank() error = %v", err)
	}
	if resp == nil {
		t.Fatal("GetAwdpProblemRank() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if requester.lastReq.Method != http.MethodPost {
		t.Fatalf("request method = %s", requester.lastReq.Method)
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/competition/getAwdpProblemRank" {
		t.Fatalf("request URL = %s", got)
	}
	body, readErr := io.ReadAll(requester.lastReq.Body)
	if readErr != nil {
		t.Fatalf("ReadAll() error = %v", readErr)
	}
	if !bytes.Contains(body, []byte(`"problemId":"9001"`)) {
		t.Fatalf("request body = %s", string(body))
	}
	if len(resp.AttackSpeedRank) != 1 || len(resp.DefenseSpeedRank) != 1 {
		t.Fatalf("response = %#v", resp)
	}
	if resp.AttackSpeedRank[0].UserId != "42" || resp.AttackSpeedRank[0].Duration != 40 || resp.AttackSpeedRank[0].Rank != 1 {
		t.Fatalf("response = %#v", resp)
	}
	if resp.DefenseSpeedRank[0].UserName != "carol" || resp.DefenseSpeedRank[0].ProblemId != "9001" {
		t.Fatalf("response = %#v", resp)
	}
}

func TestGetUserCompetitionRecord(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"totalCount":3,"inProcessCount":1,"notStartCount":1,"endedCount":1,"competitions":[{"id":"c-1","name":"AWDP Spring","postCompetitionSnapshotPath":"/snapshots/c-1.json"}]}}`}
	client := newAuthedClient(requester)

	resp, err := client.GetUserCompetitionRecord(&sdkreq.GetUserCompetitionRecordReq{Page: 1, Size: 10})
	if err != nil {
		t.Fatalf("GetUserCompetitionRecord() error = %v", err)
	}
	if resp == nil {
		t.Fatal("GetUserCompetitionRecord() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/competition/getUserCompetitionRecord" {
		t.Fatalf("request URL = %s", got)
	}
	body, readErr := io.ReadAll(requester.lastReq.Body)
	if readErr != nil {
		t.Fatalf("ReadAll() error = %v", readErr)
	}
	for _, want := range []string{`"page":1`, `"size":10`} {
		if !bytes.Contains(body, []byte(want)) {
			t.Fatalf("request body missing %s: %s", want, string(body))
		}
	}
	if resp.TotalCount != 3 || len(resp.Competitions) != 1 {
		t.Fatalf("response = %#v", resp)
	}
}

func TestGetMyCompetitionAnalysis(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"totalCompetitions":9,"notStartCompetitions":2,"inProcessCompetitions":3,"endedCompetitions":4,"snapshotReadyCount":4,"lastSnapshotUpdatedAt":1700000001,"averageSolveSeconds":66,"bestSolveSeconds":12,"worstSolveSeconds":201,"strongestTag":"web","weakestTag":"misc","nextTrainingDirection":"pwn","recommendedFirstDirection":"web","recommendedSlowDirection":"misc","trainingSuggestion":"train pwn","strategyAdvice":"solve web first"}}`}
	client := newAuthedClient(requester)

	resp, err := client.GetMyCompetitionAnalysis(&sdkreq.GetMyCompetitionAnalysisReq{})
	if err != nil {
		t.Fatalf("GetMyCompetitionAnalysis() error = %v", err)
	}
	if resp == nil {
		t.Fatal("GetMyCompetitionAnalysis() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/competition/getMyCompetitionAnalysis" {
		t.Fatalf("request URL = %s", got)
	}
	if resp.TotalCompetitions != 9 || resp.SnapshotReadyCount != 4 || resp.LastSnapshotUpdatedAt != 1700000001 {
		t.Fatalf("response = %#v", resp)
	}
	if resp.AverageSolveSeconds != 66 || resp.BestSolveSeconds != 12 || resp.WorstSolveSeconds != 201 {
		t.Fatalf("response = %#v", resp)
	}
	if resp.StrongestTag != "web" || resp.WeakestTag != "misc" {
		t.Fatalf("response = %#v", resp)
	}
	if resp.NextTrainingDirection != "pwn" || resp.RecommendedFirstDirection != "web" {
		t.Fatalf("response = %#v", resp)
	}
	if resp.RecommendedSlowDirection != "misc" || resp.TrainingSuggestion != "train pwn" || resp.StrategyAdvice != "solve web first" {
		t.Fatalf("response = %#v", resp)
	}
}

func TestGetPremiumCompetitionAnalysis(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"accessScope":"scope","totalCompetitions":5,"endedCompetitions":4,"snapshotReadyCount":4,"strongestDirection":{"tag":"web","solvedCount":7},"weakestDirection":{"tag":"pwn","solvedCount":1},"nextTrainingDirection":"pwn","recommendedFirstDirection":"web","recommendedSlowDirection":"pwn","trainingSuggestion":"train pwn","strategyAdvice":"solve web first"}}`}
	client := newAuthedClient(requester)

	resp, err := client.GetPremiumCompetitionAnalysis(&sdkreq.GetPremiumCompetitionAnalysisReq{
		RefType: "awdp_bundle",
		RefId:   "9101",
	})
	if err != nil {
		t.Fatalf("GetPremiumCompetitionAnalysis() error = %v", err)
	}
	if resp == nil {
		t.Fatal("GetPremiumCompetitionAnalysis() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/competition/getPremiumCompetitionAnalysis" {
		t.Fatalf("request URL = %s", got)
	}
	body, readErr := io.ReadAll(requester.lastReq.Body)
	if readErr != nil {
		t.Fatalf("ReadAll() error = %v", readErr)
	}
	for _, want := range []string{`"refType":"awdp_bundle"`, `"refId":"9101"`} {
		if !bytes.Contains(body, []byte(want)) {
			t.Fatalf("request body missing %s: %s", want, string(body))
		}
	}
	if resp.AccessScope != "scope" || resp.StrongestDirection.Tag != "web" || resp.WeakestDirection.SolvedCount != 1 {
		t.Fatalf("response = %#v", resp)
	}
}

func TestGetMyProblemAnalysis(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"directionProgress":[{"tag":"web","count":3},{"tag":"pwn","count":1}],"rank":5,"totalSolved":4,"averageSolveSeconds":66,"solvedWithDurationCount":3,"fastestProblem":{"problemId":"301","problemName":"web-quick","solveSeconds":12},"slowestProblem":{"problemId":"302","problemName":"pwn-slow","solveSeconds":201}}}`}
	client := newAuthedClient(requester)

	resp, err := client.GetMyProblemAnalysis(&sdkreq.GetMyProblemAnalysisReq{})
	if err != nil {
		t.Fatalf("GetMyProblemAnalysis() error = %v", err)
	}
	if resp == nil {
		t.Fatal("GetMyProblemAnalysis() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/problem/getMyProblemAnalysis" {
		t.Fatalf("request URL = %s", got)
	}
	if resp.Rank != 5 || resp.TotalSolved != 4 || resp.AverageSolveSeconds != 66 || resp.SolvedWithDurationCount != 3 {
		t.Fatalf("response = %#v", resp)
	}
	if len(resp.DirectionProgress) != 2 || resp.DirectionProgress[0].Tag != "web" || resp.DirectionProgress[0].Count != 3 {
		t.Fatalf("response = %#v", resp)
	}
	if resp.FastestProblem.ProblemId != "301" || resp.SlowestProblem.SolveSeconds != 201 {
		t.Fatalf("response = %#v", resp)
	}
}

func TestGetPremiumProblemAnalysis(t *testing.T) {
	requester := &fakeRequester{respBody: `{"code":200,"msg":"ok","data":{"accessScope":"scope","strongestDirection":{"tag":"web","solvedCount":7},"weakestDirection":{"tag":"pwn","solvedCount":1},"averageSolveSeconds":88,"fastestProblem":{"problemId":"301","problemName":"web-quick","solveSeconds":12},"slowestProblem":{"problemId":"302","problemName":"pwn-slow","solveSeconds":201},"nextTrainingDirection":"pwn","recommendedFirstDirection":"web","recommendedSlowDirection":"pwn","trainingSuggestion":"train pwn","strategyAdvice":"solve web first"}}`}
	client := newAuthedClient(requester)

	resp, err := client.GetPremiumProblemAnalysis(&sdkreq.GetPremiumProblemAnalysisReq{
		RefType: "awdp_problem",
		RefId:   "9101",
	})
	if err != nil {
		t.Fatalf("GetPremiumProblemAnalysis() error = %v", err)
	}
	if resp == nil {
		t.Fatal("GetPremiumProblemAnalysis() returned nil response")
	}
	if requester.lastReq == nil {
		t.Fatal("requester did not receive request")
	}
	if got := requester.lastReq.URL.String(); got != "http://sdk.test/problem/getPremiumProblemAnalysis" {
		t.Fatalf("request URL = %s", got)
	}
	body, readErr := io.ReadAll(requester.lastReq.Body)
	if readErr != nil {
		t.Fatalf("ReadAll() error = %v", readErr)
	}
	for _, want := range []string{`"refType":"awdp_problem"`, `"refId":"9101"`} {
		if !bytes.Contains(body, []byte(want)) {
			t.Fatalf("request body missing %s: %s", want, string(body))
		}
	}
	if resp.AccessScope != "scope" || resp.StrongestDirection.Tag != "web" || resp.WeakestDirection.SolvedCount != 1 {
		t.Fatalf("response = %#v", resp)
	}
	if resp.AverageSolveSeconds != 88 || resp.FastestProblem.ProblemId != "301" || resp.SlowestProblem.SolveSeconds != 201 {
		t.Fatalf("response = %#v", resp)
	}
	if resp.NextTrainingDirection != "pwn" || resp.RecommendedFirstDirection != "web" || resp.RecommendedSlowDirection != "pwn" {
		t.Fatalf("response = %#v", resp)
	}
	if resp.TrainingSuggestion != "train pwn" || resp.StrategyAdvice != "solve web first" {
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
