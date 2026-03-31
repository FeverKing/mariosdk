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

type AwdpSpeedRankCell struct {
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
	ProblemId string `json:"problemId"`
	Duration  int64  `json:"duration"`
	Rank      int64  `json:"rank"`
}

type GetAwdpBundleDetailModel struct {
	BundleId     string   `json:"bundleId"`
	BundleName   string   `json:"bundleName"`
	ProblemIds   []string `json:"problemIds"`
	ProblemCount int64    `json:"problemCount"`
}

type GetAwdpProblemRankModel struct {
	AttackSpeedRank  []AwdpSpeedRankCell `json:"attackSpeedRank"`
	DefenseSpeedRank []AwdpSpeedRankCell `json:"defenseSpeedRank"`
}

type TagSolveCount struct {
	Tag   string `json:"tag"`
	Count uint64 `json:"count"`
}

type CompetitionRecordModel struct {
	Id                               string          `json:"id"`
	Name                             string          `json:"name"`
	ShortName                        string          `json:"shortName"`
	Image                            string          `json:"image"`
	StartTime                        uint64          `json:"startTime"`
	EndTime                          uint64          `json:"endTime"`
	Privilege                        int32           `json:"privilege"`
	CompetitionType                  int32           `json:"competitionType"`
	Status                           int32           `json:"status"`
	PostCompetitionSnapshotPath      string          `json:"postCompetitionSnapshotPath"`
	PostCompetitionSnapshotUpdatedAt uint64          `json:"postCompetitionSnapshotUpdatedAt"`
	SnapshotSolvedCount              uint64          `json:"snapshotSolvedCount"`
	SnapshotAverageSolveSeconds      uint64          `json:"snapshotAverageSolveSeconds"`
	SnapshotBestSolveSeconds         uint64          `json:"snapshotBestSolveSeconds"`
	SnapshotWorstSolveSeconds        uint64          `json:"snapshotWorstSolveSeconds"`
	SnapshotTagSummary               []TagSolveCount `json:"snapshotTagSummary"`
}

type GetUserCompetitionRecordModel struct {
	TotalCount     int64                    `json:"totalCount"`
	InProcessCount int64                    `json:"inProcessCount"`
	NotStartCount  int64                    `json:"notStartCount"`
	EndedCount     int64                    `json:"endedCount"`
	Competitions   []CompetitionRecordModel `json:"competitions"`
}

type GetMyCompetitionAnalysisModel struct {
	TotalCompetitions         int64  `json:"totalCompetitions"`
	NotStartCompetitions      int64  `json:"notStartCompetitions"`
	InProcessCompetitions     int64  `json:"inProcessCompetitions"`
	EndedCompetitions         int64  `json:"endedCompetitions"`
	SnapshotReadyCount        int64  `json:"snapshotReadyCount"`
	LastSnapshotUpdatedAt     uint64 `json:"lastSnapshotUpdatedAt"`
	AverageSolveSeconds       uint64 `json:"averageSolveSeconds"`
	BestSolveSeconds          uint64 `json:"bestSolveSeconds"`
	WorstSolveSeconds         uint64 `json:"worstSolveSeconds"`
	StrongestTag              string `json:"strongestTag"`
	WeakestTag                string `json:"weakestTag"`
	NextTrainingDirection     string `json:"nextTrainingDirection"`
	RecommendedFirstDirection string `json:"recommendedFirstDirection"`
	RecommendedSlowDirection  string `json:"recommendedSlowDirection"`
	TrainingSuggestion        string `json:"trainingSuggestion"`
	StrategyAdvice            string `json:"strategyAdvice"`
}

type PremiumCompetitionDirectionSummaryModel struct {
	Tag         string `json:"tag"`
	SolvedCount int    `json:"solvedCount"`
}

type ProblemRecordModel struct {
	Tag   string `json:"tag"`
	Count int    `json:"count"`
}

type ProblemSolveDurationCellModel struct {
	ProblemId    string `json:"problemId"`
	ProblemName  string `json:"problemName"`
	SolveSeconds int64  `json:"solveSeconds"`
}

type GetMyProblemAnalysisModel struct {
	DirectionProgress       []ProblemRecordModel          `json:"directionProgress"`
	Rank                    int                           `json:"rank"`
	TotalSolved             int                           `json:"totalSolved"`
	AverageSolveSeconds     int64                         `json:"averageSolveSeconds"`
	SolvedWithDurationCount int64                         `json:"solvedWithDurationCount"`
	FastestProblem          ProblemSolveDurationCellModel `json:"fastestProblem"`
	SlowestProblem          ProblemSolveDurationCellModel `json:"slowestProblem"`
}

type PremiumProblemDirectionSummaryModel struct {
	Tag         string `json:"tag"`
	SolvedCount int    `json:"solvedCount"`
}

type GetPremiumProblemAnalysisModel struct {
	AccessScope               string                              `json:"accessScope"`
	StrongestDirection        PremiumProblemDirectionSummaryModel `json:"strongestDirection"`
	WeakestDirection          PremiumProblemDirectionSummaryModel `json:"weakestDirection"`
	AverageSolveSeconds       int64                               `json:"averageSolveSeconds"`
	FastestProblem            ProblemSolveDurationCellModel       `json:"fastestProblem"`
	SlowestProblem            ProblemSolveDurationCellModel       `json:"slowestProblem"`
	NextTrainingDirection     string                              `json:"nextTrainingDirection"`
	RecommendedFirstDirection string                              `json:"recommendedFirstDirection"`
	RecommendedSlowDirection  string                              `json:"recommendedSlowDirection"`
	TrainingSuggestion        string                              `json:"trainingSuggestion"`
	StrategyAdvice            string                              `json:"strategyAdvice"`
}

type GetPremiumCompetitionAnalysisModel struct {
	AccessScope               string                                  `json:"accessScope"`
	TotalCompetitions         int64                                   `json:"totalCompetitions"`
	EndedCompetitions         int64                                   `json:"endedCompetitions"`
	SnapshotReadyCount        int64                                   `json:"snapshotReadyCount"`
	StrongestDirection        PremiumCompetitionDirectionSummaryModel `json:"strongestDirection"`
	WeakestDirection          PremiumCompetitionDirectionSummaryModel `json:"weakestDirection"`
	NextTrainingDirection     string                                  `json:"nextTrainingDirection"`
	RecommendedFirstDirection string                                  `json:"recommendedFirstDirection"`
	RecommendedSlowDirection  string                                  `json:"recommendedSlowDirection"`
	TrainingSuggestion        string                                  `json:"trainingSuggestion"`
	StrategyAdvice            string                                  `json:"strategyAdvice"`
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

type BoolRespModel struct {
	Result  bool   `json:"result"`
	Message string `json:"message,omitempty"`
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
	Type int64  `json:"type"`
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
