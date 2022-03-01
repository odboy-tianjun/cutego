package v1

import (
	"cutego/core/api/v1/request"
	"cutego/core/entity"
	"cutego/core/service"
	"cutego/pkg/resp"
	"cutego/pkg/tree/tree_dept"
	"cutego/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type DeptApi struct {
	deptService service.DeptService
}

// TreeSelect 查询部门菜单树
func (a DeptApi) DeptTreeSelect(c *gin.Context) {
	query := request.DeptQuery{}
	if c.BindQuery(&query) == nil {
		treeSelect := a.deptService.FindTreeSelect(query)
		list := tree_dept.DeptList{}
		c.JSON(http.StatusOK, resp.Success(list.GetTree(treeSelect)))
	} else {
		c.JSON(http.StatusInternalServerError, resp.ErrorResp("参数绑定错误"))
	}
}

// RoleDeptTreeSelect 加载对应角色部门列表树
func (a DeptApi) RoleDeptTreeSelect(c *gin.Context) {
	m := make(map[string]interface{})
	param := c.Param("roleId")
	roleId, _ := strconv.ParseInt(param, 10, 64)
	checkedKeys := a.deptService.FindDeptListByRoleId(roleId)
	m["checkedKeys"] = checkedKeys
	treeSelect := a.deptService.FindTreeSelect(request.DeptQuery{})
	list := tree_dept.DeptList{}
	tree := list.GetTree(treeSelect)
	m["depts"] = tree
	resp.OK(c, m)
}

// Find 查询部门列表
func (a DeptApi) Find(c *gin.Context) {
	query := request.DeptQuery{}
	if c.BindQuery(&query) != nil {
		resp.ParamError(c)
		return
	}
	resp.OK(c, a.deptService.FindDeptList(query))
}

// ExcludeChild 查询部门列表（排除节点)
func (a DeptApi) ExcludeChild(c *gin.Context) {
	param := c.Param("deptId")
	deptId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c)
		return
	}
	list := a.deptService.FindDeptList(request.DeptQuery{})
	var depts = *list
	deptList := make([]entity.SysDept, 0)
	for _, dept := range depts {
		if dept.DeptId == deptId || strings.Contains(dept.Ancestors, strconv.Itoa(deptId)) {
			continue
		}
		deptList = append(deptList, dept)
	}
	resp.OK(c, deptList)
}

// GetInfo 根据部门编号获取详细信息
func (a DeptApi) GetInfo(c *gin.Context) {
	param := c.Param("deptId")
	deptId, err := strconv.Atoi(param)
	if err != nil {
		resp.ParamError(c)
		return
	}
	resp.OK(c, a.deptService.GetDeptById(deptId))
}

// Add 添加部门
func (a DeptApi) Add(c *gin.Context) {
	dept := entity.SysDept{}
	if c.Bind(&dept) != nil {
		resp.ParamError(c)
		return
	}
	// 校验部门名称是否唯一
	unique := a.deptService.CheckDeptNameUnique(dept)
	if unique {
		resp.Error(c, "新增部门'"+dept.DeptName+"'失败, 部门名称已存在")
		return
	}
	info := a.deptService.GetDeptById(dept.ParentId)
	if info.Status == "1" {
		resp.Error(c, "部门停用, 不允许新增")
		return
	}
	dept.Ancestors = info.Ancestors + "," + strconv.Itoa(dept.ParentId)
	dept.CreateBy = util.GetUserInfo(c).UserName
	if a.deptService.Save(dept) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Delete 删除部门
func (a DeptApi) Delete(c *gin.Context) {
	param := c.Param("deptId")
	deptId, _ := strconv.Atoi(param)
	// 是否存在部门子节点
	if a.deptService.HasChildByDeptId(deptId) > 0 {
		resp.Error(c, "存在下级部门,不允许删除")
		return
	}
	if a.deptService.CheckDeptExistUser(deptId) > 0 {
		resp.Error(c, "部门存在用户,不允许删除")
		return
	}
	if a.deptService.Remove(deptId) > 0 {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}

// Edit 修改部门数据接口
func (a DeptApi) Edit(c *gin.Context) {
	dept := entity.SysDept{}
	if c.Bind(&dept) != nil {
		resp.ParamError(c)
		return
	}
	dept.UpdateTime = time.Now()
	dept.UpdateBy = util.GetUserInfo(c).UserName
	if a.deptService.Edit(dept) {
		resp.OK(c)
	} else {
		resp.Error(c)
	}
}
