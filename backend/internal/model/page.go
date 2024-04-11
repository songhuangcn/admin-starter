package model

// 因为有默认值，因此这里不需要 omitempty，如果使用了 omitempty，那前端就能传入零值 0 了
type PageRequest struct {
	Page     int `form:"page"      binding:"min=1"`
	PageSize int `form:"per_page"  binding:"min=1,max=1000"`
}

func NewPageRequest() *PageRequest {
	return &PageRequest{1, 25}
}
