package request

// LoginBody 登录参数
type LoginBody struct {
	UserName string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
}
