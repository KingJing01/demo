package inputmodels

//user login  from info
type LoginInfo struct {
	UserName string `form:"username"`
	Password string `form:"password"`
	SysCode  string `form:"sysCode"`
}

type SetMeatInput struct {
	Id          int    `from:"id"`
	SetMealName string `from:"setMealName"`
	SetMealCode string `from:"setMealCode"`
	SysCode     string `from:"sysCode"`
	PerId       string `from:"perId"`
	PerName     string `from:"perName"`
}

//RoleInput 角色信息入参接受
type RoleInput struct {
	Id       int    `from:"id"`
	RoleName string `from:"roleName"`
	AuthText string `from:"authText"`
	SysCode  string `from:"sysCode"`
	PerId    string `from:"perId"`
	PerName  string `from:"perName"`
}
