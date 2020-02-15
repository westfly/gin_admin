package schema

import (
	"time"
)

// AreaQuery 地域查询对象
type AreaQuery struct {
	RecordID  string    `json:"record_id"`  // 记录ID
	Creator   string    `json:"creator"`    // 创建者
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

// AreaQueryQueryParam 查询条件
type AreaQueryQueryParam struct {
}

// AreaQueryQueryOptions 查询可选参数项
type AreaQueryQueryOptions struct {
	PageParam *PaginationParam // 分页参数
}

// AreaQueryQueryResult 查询结果
type AreaQueryQueryResult struct {
	Data       AreaQueries
	PageResult *PaginationResult
}

// AreaQueries 地域查询列表
type AreaQueries []*AreaQuery
