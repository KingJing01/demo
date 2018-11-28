package outmodels

//系统操作的返回结果
type OperResult struct {
	Result  int // 1 表示成功  0表示失败
	Message string
	Data    interface{}
}

//登陆成功的返回结果
type LoginResult struct {
	OperResult
	Token string
}
