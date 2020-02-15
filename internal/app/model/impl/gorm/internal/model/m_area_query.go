package model

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/westfly/gin-admin/internal/app/errors"
	"github.com/westfly/gin-admin/internal/app/model/impl/gorm/internal/entity"
	"github.com/westfly/gin-admin/internal/app/schema"
)

// NewAreaQuery 创建地域查询存储实例
func NewAreaQuery(db *gorm.DB) *AreaQuery {
	return &AreaQuery{db}
}

// AreaQuery 地域查询存储
type AreaQuery struct {
	db *gorm.DB
}

func (a *AreaQuery) getQueryOption(opts ...schema.AreaQueryQueryOptions) schema.AreaQueryQueryOptions {
	var opt schema.AreaQueryQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

// Query 查询数据
func (a *AreaQuery) Query(ctx context.Context, params schema.AreaQueryQueryParam, opts ...schema.AreaQueryQueryOptions) (*schema.AreaQueryQueryResult, error) {
	opt := a.getQueryOption(opts...)
	db := entity.GetAreaQueryDB(ctx, a.db)
	// TODO: 查询条件
	db = db.Order("id DESC")

	var list entity.AreaQueries
	pr, err := WrapPageQuery(ctx, db, opt.PageParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.AreaQueryQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaAreaQueries(),
	}

	return qr, nil
}

// Get 查询指定数据
func (a *AreaQuery) Get(ctx context.Context, recordID string, opts ...schema.AreaQueryQueryOptions) (*schema.AreaQuery, error) {
	db := entity.GetAreaQueryDB(ctx, a.db).Where("record_id=?", recordID)
	var item entity.AreaQuery
	ok, err := FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaAreaQuery(), nil
}

// Create 创建数据
func (a *AreaQuery) Create(ctx context.Context, item schema.AreaQuery) error {
	eitem := entity.SchemaAreaQuery(item).ToAreaQuery()
	result := entity.GetAreaQueryDB(ctx, a.db).Create(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Update 更新数据
func (a *AreaQuery) Update(ctx context.Context, recordID string, item schema.AreaQuery) error {
	eitem := entity.SchemaAreaQuery(item).ToAreaQuery()
	result := entity.GetAreaQueryDB(ctx, a.db).Where("record_id=?", recordID).Omit("record_id").Updates(eitem)
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete 删除数据
func (a *AreaQuery) Delete(ctx context.Context, recordID string) error {
	result := entity.GetAreaQueryDB(ctx, a.db).Where("record_id=?", recordID).Delete(entity.AreaQuery{})
	if err := result.Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}
