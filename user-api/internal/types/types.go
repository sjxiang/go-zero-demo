// Code generated by goctl. DO NOT EDIT.
package types

type UserInfoReq struct {
	UserId int64 `path:"userId"`
}

type UserInfoResp struct {
	UserId   int64  `json:"userId"`
	NickName string `json:"nickname"`
}

type UserCreateReq struct {
	Mobile   string `json:"mobile"`
	NickName string `json:"nickname"`
}

type UserCreateResp struct {
	Flag bool `json:"flag"`
}

type UserUpdateReq struct {
	UserId   int64  `json:"userId"`
	NickName string `json:"nickname"`
}

type UserUpdateResp struct {
	Flag bool `json:"flag"`
}

type UserTestReq struct {
}

type UserTestResp struct {
}
