### 1. "获取用户信息"

1. 路由定义

- Url: /user/info
- Method: POST
- Request: `UserInfoReq`
- Response: `UserInfoResp`

2. 请求定于

```golang
type UserInfoReq struct {
	UserId int64 `json:"userId"`
}
```


3. 返回定义 

```golang
type UserInfoResp struct {
	UserId int64 `json:"userId"`
	NickName string `json:"nickname"`
}
```

