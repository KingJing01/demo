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

// 菜单列表数据
type MenuInfo struct {
	Id       int64
	MenuName string
	SysName  string
	MenuText string
}

type SelectInfo struct {
	code  string
	label string
}

// 权限系统 系统配置返回结构
type SysInfo struct {
	Id      int64
	IsValid string
	SysCode string
	SysName string
}

//套餐返回结构
type SetMealInfo struct {
	Id             int64
	SetMealName    string
	SetMealCode    string
	SysName        string
	PermissionText string
}
