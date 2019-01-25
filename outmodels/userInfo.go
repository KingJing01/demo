package outmodels

type UserInfoToken struct {
	UserName string
	Phone    string
}

//TMSRespData TMS系统对接返回消息接受对象
type TMSRespData struct {
	Success bool
	Msg     string
	Data    interface{}
}
