package sdkmodel

type BaseResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type AuthModel struct {
	Status       string `json:"status"`
	Name         string `json:"name"`
	Token        string `json:"token"`
	AccessExpire int    `json:"accessExpire"`
	RefreshAfter int    `json:"refreshAfter"`
}

type BatchUserInfoModel struct {
	Users []struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Avatar string `json:"avatar"`
	} `json:"users"`
}

type SearchPublicProblemModel struct {
	Problems []struct {
		Id          string      `json:"id"`
		Name        string      `json:"name"`
		CreateId    string      `json:"createId"`
		OwnerId     string      `json:"ownerId"`
		ProblemType int         `json:"problemType"`
		Tags        interface{} `json:"tags"`
		Attachments interface{} `json:"attachments"`
		CreateName  string      `json:"createName"`
		Permission  int         `json:"permission"`
		Difficulty  int         `json:"difficulty"`
		PublicId    string      `json:"publicId"`
		PublicType  int         `json:"publicType"`
		Desc        string      `json:"desc"`
		IsSolved    bool        `json:"isSolved"`
	} `json:"problems"`
	Total int `json:"total"`
}

type GetUserInfoForCompetitionModel struct {
	Users []struct {
		UserId     string `json:"userId"`
		Username   string `json:"username"`
		UserAvatar string `json:"userAvatar"`
		Motto      string `json:"motto"`
	} `json:"users"`
}

type StartChallengeContainerModel struct {
	Address     []string `json:"address"`
	ContainerId string   `json:"containerId"`
	RestTime    int      `json:"restTime"`
}

type StopChallengeContainerModel struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type CheckTmpLoginVerifyTokenModel struct {
	UserId                string `json:"userId"`
	CompetitionPermission int    `json:"competitionPermission,omitempty"`
}

type CheckCompetitionAWDPModel struct {
	IsCorrect bool `json:"isCorrect"`
}

type AwdpPatchApplyModel struct {
	PatchId int    `json:"patchId"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type UploadCompetitionScoreModel struct {
	Success bool   `json:"success"` // 请求是否成功
	Message string `json:"message"` // 返回的提示信息
}
