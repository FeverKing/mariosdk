package sdkclient

import (
	"mariosdk/sdk/sdklog"
	"mariosdk/sdk/sdkmodel"
	"mariosdk/sdk/sdkreq"
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
	return nil
}

func (c *DefaultClient) GetBatchUserInfo(ids []string) (*sdkmodel.BatchUserInfoModel, error) {
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
	res, err := c.apiClient.CallSearchPublicProblemApi(req)
	if err != nil {
		sdklog.Errorf("search public problem failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) GetUserInfoForCompetition(req *sdkreq.GetUserInfoForCompetitionReq) (*sdkmodel.GetUserInfoForCompetitionModel, error) {
	res, err := c.apiClient.CallGetUserInfoForCompetitionApi(req)
	if err != nil {
		sdklog.Errorf("get user info for competition failed: %v", err)
		return nil, err
	}
	return res, nil
}

func (c *DefaultClient) CheckCompetitionPrivilege(req *sdkreq.CheckCompetitionPrivilegeReq) (*sdkmodel.CheckCompetitionPrivilegeModel, error) {
	res, err := c.apiClient.CallCheckCompetitionPrivilegeApi(req)
	if err != nil {
		sdklog.Errorf("check competition privilege failed: %v", err)
		return nil, err
	}
	return res, nil
}
