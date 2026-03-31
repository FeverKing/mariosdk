package sdkclient

import (
	"errors"
	"github.com/FeverKing/mariosdk/sdk/sdklog"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
	"github.com/FeverKing/mariosdk/sdk/sdkreq"
	"time"
)

func (c *DefaultClient) Auth() error {
	if len(c.Config.Endpoints) == 0 {
		return errors.New("no sdk endpoint configured")
	}
	c.apiClient = sdkreq.NewApiClient(sdkreq.NewHttpRequester(), c.Config.Endpoints[0])
	// authenticate
	sdklog.Infof("authenticating with %s", c.Config.AccessKey)
	ar := sdkreq.AuthApiRequest{
		AccessKey: c.Config.AccessKey,
		SecretKey: c.Config.SecretKey,
	}
	err := c.apiClient.CallAuthApi(&ar)
	if err != nil {
		sdklog.Errorf("auth failed: %v", err)
		return err
	}

	c.apiClient.TokenExpiry = time.Now().Add(12 * time.Hour)
	//c.apiClient.TokenExpiry = time.Now().Add(1 * time.Minute)
	return nil
}

func (c *DefaultClient) ensureAuth() error {
	if c.apiClient == nil {
		return c.Auth()
	}
	if time.Now().After(c.apiClient.TokenExpiry) {
		sdklog.Infof("Token expired. Re-authenticating.")
		return c.Auth()
	}
	return nil
}

func (c *DefaultClient) ensureAPIClient() error {
	if c.apiClient != nil {
		return nil
	}
	if len(c.Config.Endpoints) == 0 {
		return errors.New("no sdk endpoint configured")
	}
	c.apiClient = sdkreq.NewApiClient(sdkreq.NewHttpRequester(), c.Config.Endpoints[0])
	return nil
}

func (c *DefaultClient) GetBatchUserInfo(ids []string) (*sdkmodel.BatchUserInfoModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	var batchUserInfo sdkreq.BatchUserInfoRequest
	batchUserInfo.Ids = ids
	res, err := c.apiClient.CallUserInfoApi(&batchUserInfo)
	if err != nil {
		sdklog.Errorf("get user info failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) SearchPublicProblem(req *sdkreq.SearchPublicProblemReq) (*sdkmodel.SearchPublicProblemModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallSearchPublicProblemApi(req)
	if err != nil {
		sdklog.Errorf("search public problem failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetUserInfoForCompetition(req *sdkreq.GetUserInfoForCompetitionReq) (*sdkmodel.GetUserInfoForCompetitionModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetUserInfoForCompetitionApi(req)
	if err != nil {
		sdklog.Errorf("get user info for competition failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) StartChallengeContainer(req *sdkreq.StartChallengeContainerReq) (*sdkmodel.StartChallengeContainerModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallStartChallengeContainerApi(req)
	if err != nil {
		sdklog.Errorf("start challenge container failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetAwdpBundleDetail(req *sdkreq.GetAwdpBundleDetailReq) (*sdkmodel.GetAwdpBundleDetailModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetAwdpBundleDetailApi(req)
	if err != nil {
		sdklog.Errorf("get awdp bundle detail failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetAwdpProblemRank(req *sdkreq.GetAwdpProblemRankReq) (*sdkmodel.GetAwdpProblemRankModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetAwdpProblemRankApi(req)
	if err != nil {
		sdklog.Errorf("get awdp problem rank failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetUserCompetitionRecord(req *sdkreq.GetUserCompetitionRecordReq) (*sdkmodel.GetUserCompetitionRecordModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetUserCompetitionRecordApi(req)
	if err != nil {
		sdklog.Errorf("get user competition record failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetMyCompetitionAnalysis(req *sdkreq.GetMyCompetitionAnalysisReq) (*sdkmodel.GetMyCompetitionAnalysisModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetMyCompetitionAnalysisApi(req)
	if err != nil {
		sdklog.Errorf("get my competition analysis failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetPremiumCompetitionAnalysis(req *sdkreq.GetPremiumCompetitionAnalysisReq) (*sdkmodel.GetPremiumCompetitionAnalysisModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetPremiumCompetitionAnalysisApi(req)
	if err != nil {
		sdklog.Errorf("get premium competition analysis failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetMyProblemAnalysis(req *sdkreq.GetMyProblemAnalysisReq) (*sdkmodel.GetMyProblemAnalysisModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetMyProblemAnalysisApi(req)
	if err != nil {
		sdklog.Errorf("get my problem analysis failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetPremiumProblemAnalysis(req *sdkreq.GetPremiumProblemAnalysisReq) (*sdkmodel.GetPremiumProblemAnalysisModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetPremiumProblemAnalysisApi(req)
	if err != nil {
		sdklog.Errorf("get premium problem analysis failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) StopChallengeContainer(req *sdkreq.StopChallengeContainerReq) (*sdkmodel.StopChallengeContainerModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallStopChallengeContainerApi(req)
	if err != nil {
		sdklog.Errorf("stop challenge container failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) CheckTmpLoginVerifyToken(req *sdkreq.CheckTmpLoginVerifyTokenReq) (*sdkmodel.CheckTmpLoginVerifyTokenModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallCheckTmpLoginVerifyTokenApi(req)
	if err != nil {
		sdklog.Errorf("check tmp login verify token failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetAuthToken() (string, error) {
	if err := c.ensureAuth(); err != nil {
		return "", err
	}

	return c.apiClient.Token, nil
}

func (c *DefaultClient) GetCompetitionSetting(req *sdkreq.GetCompetitionSettingReq) ([]byte, error) {
	if err := c.ensureAPIClient(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetCompetitionSettingApi(req)
	if err != nil {
		sdklog.Errorf("get competition setting failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionAllIdentities(req *sdkreq.GetCompetitionAllIdentitiesReq) ([]byte, error) {
	if err := c.ensureAPIClient(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetCompetitionAllIdentitiesApi(req)
	if err != nil {
		sdklog.Errorf("get competition all identities failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionAllTeams(req *sdkreq.GetCompetitionAllTeamsReq) ([]byte, error) {
	if err := c.ensureAPIClient(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetCompetitionAllTeamsApi(req)
	if err != nil {
		sdklog.Errorf("get competition all teams failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionAllUsers(req *sdkreq.GetCompetitionAllUsersReq) ([]byte, error) {
	if err := c.ensureAPIClient(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetCompetitionAllUsersApi(req)
	if err != nil {
		sdklog.Errorf("get competition all users failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionTemplate(req *sdkreq.GetCompetitionTemplateReq) ([]byte, error) {
	if err := c.ensureAPIClient(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetCompetitionTemplateApi(req)
	if err != nil {
		sdklog.Errorf("get competition template failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) CheckCompetitionAWDP(req *sdkreq.CheckCompetitionAWDPReq) (*sdkmodel.CheckCompetitionAWDPModel, error) {
	if err := c.ensureAPIClient(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallCheckCompetitionAWDPApi(req)
	if err != nil {
		sdklog.Errorf("check competition awdp failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) AwdpPatchApi(req *sdkreq.AwdpPatchApplyReq) (*sdkmodel.AwdpPatchApplyModel, error) {
	if err := c.ensureAPIClient(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallAwdpPatchApi(req)
	if err != nil {
		sdklog.Errorf("awdp patch failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) UploadCompetitionScore(req *sdkreq.UploadCompetitionScoreRequest) (*sdkmodel.UploadCompetitionScoreModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallUploadCompetitionScoreApi(req)
	if err != nil {
		sdklog.Errorf("upload competition score failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetTeamInfoForCompetition(req *sdkreq.GetTeamInfoForCompetitionRequest) (*sdkmodel.GetTeamInfoForCompetitionModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetTeamInfoForCompetitionApi(req)
	if err != nil {
		sdklog.Errorf("get user info for competition failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) PullCompetitionSnapshotForAdmin(req *sdkreq.PullCompetitionSnapshotForAdminRequest) (*sdkmodel.BoolRespModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallPullCompetitionSnapshotForAdminApi(req)
	if err != nil {
		sdklog.Errorf("pull competition snapshot for admin failed: %v", err)
		return nil, err
	}
	return res, nil
}

// ============== 题库授权相关方法 ==============

// GetAuthorizedProblemBanks 获取已授权的题库列表
func (c *DefaultClient) GetAuthorizedProblemBanks(req *sdkreq.GetAuthorizedProblemBanksReq) (*sdkmodel.GetAuthorizedProblemBanksModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetAuthorizedProblemBanksApi(req)
	if err != nil {
		sdklog.Errorf("get authorized problem banks failed: %v", err)
		return nil, err
	}
	return res, nil
}

// GetProblemBankForCompetition 获取题库详情(包含题目)
func (c *DefaultClient) GetProblemBankForCompetition(req *sdkreq.GetProblemBankForCompetitionReq) (*sdkmodel.GetProblemBankForCompetitionModel, error) {
	if err := c.ensureAuth(); err != nil {
		return nil, err
	}

	res, err := c.apiClient.CallGetProblemBankForCompetitionApi(req)
	if err != nil {
		sdklog.Errorf("get problem bank for competition failed: %v", err)
		return nil, err
	}
	return res, nil
}
