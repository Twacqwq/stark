package common

type Pager struct {
	// 页码
	Page int64
	// 页数
	PageSize int64
	// 总行数
	TotalRows int64
}

func (m *Pager) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *Pager) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *Pager) GetTotalRows() int64 {
	if m != nil {
		return m.TotalRows
	}
	return 0
}
