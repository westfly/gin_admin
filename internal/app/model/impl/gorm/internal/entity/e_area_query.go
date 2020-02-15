package entity

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/westfly/gin-admin/internal/app/schema"
)

// GetAreaQueryDB 地域查询
func GetAreaQueryDB(ctx context.Context, defDB *gorm.DB) *gorm.DB {
	return getDBWithModel(ctx, defDB, AreaQuery{})
}

// SchemaAreaQuery 地域查询
type SchemaAreaQuery schema.AreaQuery

// ToAreaQuery 转换为地域查询实体
func (a SchemaAreaQuery) ToAreaQuery() *AreaQuery {
	item := &AreaQuery{
		RecordID: &a.RecordID,
		Creator:  &a.Creator,
	}
	return item
}

// AreaQuery 地域查询实体
type AreaQuery struct {
	Model
	RecordID *string `gorm:"column:record_id;size:36;index;"` // 记录ID
	Creator  *string `gorm:"column:creator;size:36;index;"`   // 创建者
}

func (a AreaQuery) String() string {
	return toString(a)
}

// TableName 表名
func (a AreaQuery) TableName() string {
	return a.Model.TableName("area_query")
}

// ToSchemaAreaQuery 转换为地域查询对象
func (a AreaQuery) ToSchemaAreaQuery() *schema.AreaQuery {
	item := &schema.AreaQuery{
		RecordID:  *a.RecordID,
		Creator:   *a.Creator,
		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
	}
	return item
}

// AreaQueries 地域查询列表
type AreaQueries []*AreaQuery

// ToSchemaAreaQueries 转换为地域查询对象列表
func (a AreaQueries) ToSchemaAreaQueries() []*schema.AreaQuery {
	list := make([]*schema.AreaQuery, len(a))
	for i, item := range a {
		list[i] = item.ToSchemaAreaQuery()
	}
	return list
}
