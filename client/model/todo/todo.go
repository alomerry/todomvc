package todo

type AddTodoRequest struct {
	Color   string `form:"color" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type GetTodoRequest struct {
	IsAscend int    `form:"isAscend"`
	PageSize int32  `form:"pageSize"`
	SortBy   string `form:"sortBy"`
	Page     int32  `form:"page"`
	Color    string `form:"color"`
	Content  string `form:"content"`
	Status   string `form:"status"`
	Keyword  string `form:"keyword"`
	StartAt  int64  `form:"startAt"`
	EndAt    int64  `form:"endAt"`
}

type UpdateTodoRequest struct {
	Color   string `form:"color" binding:"required"`
	Content string `form:"content" binding:"required"`
	Status  string `form:"status" binding:"-"`
}
