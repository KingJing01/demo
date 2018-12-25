package inputmodels

//user login  from info
type LoginInfo struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type SetMeatInput struct {
	Id          int    `from:"id"`
	SetMealName string `from:"setMealName"`
	SetMealCode string `from:"setMealCode"`
	SysCode     string `from:"sysCode"`
	PerId       string `from:"perId"`
	PerName     string `from:"perName"`
}
