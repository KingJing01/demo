package outmodels

//系统操作的返回结果
type OperResult struct {
	Result  int // 1 表示成功  0表示失败  2 token失效 跳转登陆
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
	SysCode  string
	MenuText string
}

type SelectInfo struct {
	code  string
	label string
}

// 权限系统 系统配置返回结构
type SysInfo struct {
	Id      int64
	IsValid int
	SysCode string
	SysName string
	SysUrl  string
}

//套餐返回结构
type SetMealInfo struct {
	Id             int64
	SetMealName    string
	SetMealCode    string
	SysCode        string
	SysName        string
	PermissionText string
	IsDeleted      int64
}

type PermissionCheckInfo struct {
	Name        string // 菜单缩写 tms.order 或所属 菜单的编码
	DisplayName string // 菜单显示的中文名  订单管理
	Code        string // 权限缩写  tms.order.update
	CodeName    string // 权限中文  更新订单信息
	Flag        string // 判断是否选中
}

// 平台用户管理 列表
type UserManageInfo struct {
	Id         int
	TenantName string
	SysName    string
	SysCode    string
	MenuText   string
	Operator   string
}

type PermissonInfo struct {
	Id          int
	SysName     string
	SysCode     string
	DisplayName string
	Name        string
	MenuText    string
	MenuCode    int
	PerData     []PerInfo
}

type PerInfo struct {
	DisplayName string
	Name        string
}

type RoleInfo struct {
	Id       int
	RoleName string
	RoleCode string
	SysName  string
	SysCode  string
	IsValid  int
	AuthText string
}

type UserInfo struct {
	Id           int
	RoleName     string
	RoleCode     string
	SysName      string
	SysCode      string
	IsValid      int
	AuthText     string
	EmailAddress string
	PhoneNumber  string
	UserName     string
}

// ComponentData 组件数据
type ComponentData struct {
	// 循环条件的的判断依据
	ParentKey string
	// 显示的父节点名称
	ParentName string
	// 子组件的键值
	ChildKey string
	// 子组件的显示名
	ChildName string
}
