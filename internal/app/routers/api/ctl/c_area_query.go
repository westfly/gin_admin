package ctl

import (
	"github.com/gin-gonic/gin"
	"github.com/westfly/gin-admin/internal/app/bll"
	"github.com/westfly/gin-admin/internal/app/ginplus"
	"github.com/westfly/gin-admin/internal/app/schema"
)

// NewAreaQuery 创建地域查询控制器
func NewAreaQuery(bAreaQuery bll.IAreaQuery) *AreaQuery {
	return &AreaQuery{
		AreaQueryBll: bAreaQuery,
	}
}

// AreaQuery 地域查询控制器
type AreaQuery struct {
	AreaQueryBll bll.IAreaQuery
}

// Query 查询数据
// @Tags 地域查询
// @Summary 查询数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param current query int true "分页索引" default(1)
// @Param pageSize query int true "分页大小" default(10)
// @Success 200 {array} schema.AreaQuery "查询结果：{list:列表数据,pagination:{current:页索引,pageSize:页大小,total:总数量}}"
// @Failure 401 {object} schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/area-queries [get]
func (a *AreaQuery) Query(c *gin.Context) {
	var params schema.AreaQueryQueryParam

	result, err := a.AreaQueryBll.Query(ginplus.NewContext(c), params, schema.AreaQueryQueryOptions{
		PageParam: ginplus.GetPaginationParam(c),
	})
	if err != nil {
		ginplus.ResError(c, err)
		return
	}

	ginplus.ResPage(c, result.Data, result.PageResult)
}

// Get 查询指定数据
// @Tags 地域查询
// @Summary 查询指定数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 {object} schema.AreaQuery
// @Failure 401 {object} schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 404 {object} schema.HTTPError "{error:{code:0,message:资源不存在}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/area-queries/{id} [get]
func (a *AreaQuery) Get(c *gin.Context) {
	item, err := a.AreaQueryBll.Get(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, item)
}

// Create 创建数据
// @Tags 地域查询
// @Summary 创建数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param body body schema.AreaQuery true "创建数据"
// @Success 200 {object} schema.AreaQuery
// @Failure 400 {object} schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/area-queries [post]
func (a *AreaQuery) Create(c *gin.Context) {
	var item schema.AreaQuery
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	item.Creator = ginplus.GetUserID(c)
	nitem, err := a.AreaQueryBll.Create(ginplus.NewContext(c), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Update 更新数据
// @Tags 地域查询
// @Summary 更新数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Param body body schema.AreaQuery true "更新数据"
// @Success 200 {object} schema.AreaQuery
// @Failure 400 {object} schema.HTTPError "{error:{code:0,message:无效的请求参数}}"
// @Failure 401 {object} schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/area-queries/{id} [put]
func (a *AreaQuery) Update(c *gin.Context) {
	var item schema.AreaQuery
	if err := ginplus.ParseJSON(c, &item); err != nil {
		ginplus.ResError(c, err)
		return
	}

	nitem, err := a.AreaQueryBll.Update(ginplus.NewContext(c), c.Param("id"), item)
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResSuccess(c, nitem)
}

// Delete 删除数据
// @Tags 地域查询
// @Summary 删除数据
// @Param Authorization header string false "Bearer 用户令牌"
// @Param id path string true "记录ID"
// @Success 200 {object} schema.HTTPStatus "{status:OK}"
// @Failure 401 {object} schema.HTTPError "{error:{code:0,message:未授权}}"
// @Failure 500 {object} schema.HTTPError "{error:{code:0,message:服务器错误}}"
// @Router /api/v1/area-queries/{id} [delete]
func (a *AreaQuery) Delete(c *gin.Context) {
	err := a.AreaQueryBll.Delete(ginplus.NewContext(c), c.Param("id"))
	if err != nil {
		ginplus.ResError(c, err)
		return
	}
	ginplus.ResOK(c)
}
