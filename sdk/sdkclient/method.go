package sdkclient

import (
	"github.com/FeverKing/mariosdk/sdk/sdklog"
	"github.com/FeverKing/mariosdk/sdk/sdkmodel"
	"github.com/FeverKing/mariosdk/sdk/sdkreq"
	"time"
)

func (c *DefaultClient) Auth() error {
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
	}

	c.apiClient.TokenExpiry = time.Now().Add(12 * time.Hour)
	//c.apiClient.TokenExpiry = time.Now().Add(1 * time.Minute)
	return nil
}

func (c *DefaultClient) ensureAuth() error {
	if time.Now().After(c.apiClient.TokenExpiry) {
		sdklog.Infof("Token expired. Re-authenticating.")
		return c.Auth()
	}
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
	res, err := c.apiClient.CallGetCompetitionSettingApi(req)
	if err != nil {
		sdklog.Errorf("get competition setting failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionAllIdentities(req *sdkreq.GetCompetitionAllIdentitiesReq) ([]byte, error) {
	res, err := c.apiClient.CallGetCompetitionAllIdentitiesApi(req)
	if err != nil {
		sdklog.Errorf("get competition all identities failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionAllTeams(req *sdkreq.GetCompetitionAllTeamsReq) ([]byte, error) {
	res, err := c.apiClient.CallGetCompetitionAllTeamsApi(req)
	if err != nil {
		sdklog.Errorf("get competition all teams failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionAllUsers(req *sdkreq.GetCompetitionAllUsersReq) ([]byte, error) {
	res, err := c.apiClient.CallGetCompetitionAllUsersApi(req)
	if err != nil {
		sdklog.Errorf("get competition all users failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetCompetitionTemplate(req *sdkreq.GetCompetitionTemplateReq) ([]byte, error) {
	res, err := c.apiClient.CallGetCompetitionTemplateApi(req)
	if err != nil {
		sdklog.Errorf("get competition template failed: %v", err)
		return nil, err
	}
	return res, nil
}
