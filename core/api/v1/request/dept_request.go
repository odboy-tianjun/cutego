package request

// DeptQuery 部门查询参数结构体 GET请求
type DeptQuery struct {
	ParentId int    `form:"parentId"`
	DeptName string `form:"deptName"`
	Status   string `form:"status"`
}
