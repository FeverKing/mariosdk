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
