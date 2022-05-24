package dboperate

type PermissionInfo struct {
	Id     int    `json:"id"`
	Api    string `json:"api"`
	Method string `json:"method"`
	Desc   string `json:"desc"`
}

type PermissionInfos struct {
	Id             int               `json:"id"`
	PermissionName string            `json:"permission_name"`
	Infos          []*PermissionInfo `json:"infos"`
}

type PermissionGroup struct {
	Id        int                `json:"id"`
	GroupName string             `json:"group_name"`
	Infos     []*PermissionInfos `json:"infos"`
}

var arrPermission = make([]*PermissionGroup, 0, 10)

// 存储权限管理
var mpPermission = make(map[int]*PermissionInfo, 100)

func init() {
	arrPermission = append(arrPermission, newPermissionGroup(101, "系统管理", []*PermissionInfos{
		newPermissionInfos(10101, "用户管理", []*PermissionInfo{
			newPermissionInfo(10101001, "/api/users/user/list", "PUT", "查看用户列表"),
			newPermissionInfo(10101002, "/api/users/user/create", "POST", "创建新用户"),
			newPermissionInfo(10101003, "/api/users/user/update/*", "PUT", "修改用户数据"),
		}),
		newPermissionInfos(10102, "角色管理", []*PermissionInfo{
			newPermissionInfo(10102001, "/api/users/role/list", "PUT", "查看角色列表"),
			newPermissionInfo(10102002, "/api/users/role/create", "POST", "创建新角色"),
			newPermissionInfo(10102003, "/api/users/role/update/*", "PUT", "修改角色数据"),
		}),
		newPermissionInfos(10103, "部门管理", []*PermissionInfo{
			newPermissionInfo(10103001, "/api/users/department/list", "PUT", "查看部门列表"),
			newPermissionInfo(10103002, "/api/users/department/create", "POST", "创建新部门"),
			newPermissionInfo(10103003, "/api/users/department/update/*", "PUT", "修改部门数据"),
		}),
	}))
}

// 构造权限对象
func newPermissionInfo(id int, api, method, desc string) *PermissionInfo {
	p := &PermissionInfo{
		Id:     id,
		Api:    api,
		Method: method,
		Desc:   desc,
	}
	mpPermission[id] = p
	return p
}

// 构造二级分组
func newPermissionInfos(id int, name string, arrInfo []*PermissionInfo) *PermissionInfos {
	return &PermissionInfos{
		Id:             id,
		PermissionName: name,
		Infos:          arrInfo,
	}
}

// 构造组别
func newPermissionGroup(id int, name string, arrInfo []*PermissionInfos) *PermissionGroup {
	return &PermissionGroup{
		Id:        id,
		GroupName: name,
		Infos:     arrInfo,
	}
}

// 获取权限列表
func GetPermissionInfos() []*PermissionGroup {
	return arrPermission
}
