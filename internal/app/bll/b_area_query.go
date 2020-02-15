package bll

import (
	"context"

	"github.com/westfly/gin-admin/internal/app/schema"
)

// IAreaQuery 地域查询业务逻辑接口
type IAreaQuery interface {
	// 查询数据
	Query(ctx context.Context, params schema.AreaQueryQueryParam, opts ...schema.AreaQueryQueryOptions) (*schema.AreaQueryQueryResult, error)
	// 查询指定数据
	Get(ctx context.Context, recordID string, opts ...schema.AreaQueryQueryOptions) (*schema.AreaQuery, error)
	// 创建数据
	Create(ctx context.Context, item schema.AreaQuery) (*schema.AreaQuery, error)
	// 更新数据
	Update(ctx context.Context, recordID string, item schema.AreaQuery) (*schema.AreaQuery, error)
	// 删除数据
	Delete(ctx context.Context, recordID string) error
}
