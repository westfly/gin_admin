package internal

import (
	"context"

	"github.com/westfly/gin-admin/internal/app/errors"
	"github.com/westfly/gin-admin/internal/app/model"
	"github.com/westfly/gin-admin/internal/app/schema"
	"github.com/westfly/gin-admin/pkg/util"
)

// NewAreaQuery 创建地域查询
func NewAreaQuery(mAreaQuery model.IAreaQuery) *AreaQuery {
	return &AreaQuery{
		AreaQueryModel: mAreaQuery,
	}
}

// AreaQuery 地域查询业务逻辑
type AreaQuery struct {
	AreaQueryModel model.IAreaQuery
}

// Query 查询数据
func (a *AreaQuery) Query(ctx context.Context, params schema.AreaQueryQueryParam, opts ...schema.AreaQueryQueryOptions) (*schema.AreaQueryQueryResult, error) {
	return a.AreaQueryModel.Query(ctx, params, opts...)
}

// Get 查询指定数据
func (a *AreaQuery) Get(ctx context.Context, recordID string, opts ...schema.AreaQueryQueryOptions) (*schema.AreaQuery, error) {
	item, err := a.AreaQueryModel.Get(ctx, recordID, opts...)
	if err != nil {
		return nil, err
	} else if item == nil {
		return nil, errors.ErrNotFound
	}

	return item, nil
}

func (a *AreaQuery) getUpdate(ctx context.Context, recordID string) (*schema.AreaQuery, error) {
	return a.Get(ctx, recordID)
}

// Create 创建数据
func (a *AreaQuery) Create(ctx context.Context, item schema.AreaQuery) (*schema.AreaQuery, error) {
	item.RecordID = util.MustUUID()
	err := a.AreaQueryModel.Create(ctx, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, item.RecordID)
}

// Update 更新数据
func (a *AreaQuery) Update(ctx context.Context, recordID string, item schema.AreaQuery) (*schema.AreaQuery, error) {
	oldItem, err := a.AreaQueryModel.Get(ctx, recordID)
	if err != nil {
		return nil, err
	} else if oldItem == nil {
		return nil, errors.ErrNotFound
	}

	err = a.AreaQueryModel.Update(ctx, recordID, item)
	if err != nil {
		return nil, err
	}
	return a.getUpdate(ctx, recordID)
}

// Delete 删除数据
func (a *AreaQuery) Delete(ctx context.Context, recordID string) error {
	oldItem, err := a.AreaQueryModel.Get(ctx, recordID)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrNotFound
	}

	return a.AreaQueryModel.Delete(ctx, recordID)
}
