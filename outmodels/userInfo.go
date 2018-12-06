package outmodels

type UserInfoToken struct {
	UserName string
	Phone    string
}

// 菜单列表数据
type MenuInfo struct {
	Id       int64
	MenuName string
	SysName  string
	MenuText string
}
