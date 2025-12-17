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

// ============== 题库授权相关模型 ==============

// ProblemBankBrief 题库简要信息
type ProblemBankBrief struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Des         string `json:"des"`
	Tags        string `json:"tags"`
	Img         string `json:"img"`
	ProblemNum  int    `json:"problemNum"`
	CreatorName string `json:"creatorName"`
}

// GetAuthorizedProblemBanksModel 获取已授权题库列表响应
type GetAuthorizedProblemBanksModel struct {
	ProblemBanks []ProblemBankBrief `json:"problemBanks"`
	Total        int64              `json:"total"`
}

// ExportedProblemTag 导出的题目标签
type ExportedProblemTag struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

// ExportedProblemAttachment 导出的题目附件
type ExportedProblemAttachment struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

// ExportedProblem 导出的题目(完整信息)
type ExportedProblem struct {
	Id             string                      `json:"id"`
	Name           string                      `json:"name"`
	Desc           string                      `json:"desc"`
	ProblemType    int                         `json:"problemType"` // 0-静态 1-动态
	Difficulty     int                         `json:"difficulty"`  // 0-3 对应简单到极难
	Tags           []ExportedProblemTag        `json:"tags"`
	Attachments    []ExportedProblemAttachment `json:"attachments"`
	Answer         string                      `json:"answer,omitempty"`       // 静态题目答案
	DockerImage    string                      `json:"dockerImage,omitempty"`  // 动态题目镜像
	HttpPorts      string                      `json:"httpPorts,omitempty"`    // HTTP端口
	TcpPorts       string                      `json:"tcpPorts,omitempty"`     // TCP端口
	IsStaticAnswer bool                        `json:"isStaticAnswer"`         // 是否静态Flag
	EnvPrefix      string                      `json:"envPrefix,omitempty"`    // 环境变量前缀
	AnswerPrefix   string                      `json:"answerPrefix,omitempty"` // Flag前缀
}

// GetProblemBankForCompetitionModel 获取题库详情响应
type GetProblemBankForCompetitionModel struct {
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	Des        string            `json:"des"`
	Problems   []ExportedProblem `json:"problems"`
	ProblemNum int               `json:"problemNum"`
}
