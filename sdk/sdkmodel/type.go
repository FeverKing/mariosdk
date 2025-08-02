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
		UserId     string     `json:"userId"`
		Username   string     `json:"username"`
		UserAvatar string     `json:"userAvatar"`
		Motto      string     `json:"motto"`
		TeamId     string     `json:"teamId"`
		Identities []Identity `json:"identities"`
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

type GetTeamInfoForCompetitionModel struct {
	Teams []struct {
		TeamId     string            `json:"teamId"`
		TeamName   string            `json:"teamName"`
		TeamMotto  string            `json:"teamMotto"`
		TeamToken  string            `json:"teamToken"`
		TeamAvatar string            `json:"teamAvatar"`
		Captain    CompetitionUser   `json:"captain"`
		Members    []CompetitionUser `json:"members"`
		Identities []Identity        `json:"identities"`
	} `json:"teams"`
}
type Identity struct {
	IdentityId     string `json:"identityId,optional"`
	IdentityValue  string `json:"identityValue"`
	IdentityName   string `json:"identityName,optional"`
	IdentityBaseId string `json:"identityBaseId"`
}
type CompetitionUser struct {
	UserId     string     `json:"userId"`
	TeamId     string     `json:"teamId"`
	IsHaveTeam bool       `json:"isHaveTeam"`
	Identities []Identity `json:"identities"`
}
