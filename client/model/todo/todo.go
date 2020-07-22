package todo

type AddTodoRequest struct {
	Color   string `form:"color" binding:"required"`
	Content string `form:"content" binding:"required"`
}

type IsAscend struct {
	Value bool `form:"value"`
}

type Status struct {
	Value int `form:"value"`
}

type GetTodoRequest struct {
	IsAscend IsAscend `form:"isAscend"`
	PageSize int32    `form:"pageSize"`
	SortBy   string   `form:"sortBy"`
	Page     int32    `form:"page"`
	Color    string   `form:"color"`
	Content  string   `form:"content"`
	Status   Status   `form:"status"`
	Keyword  string   `form:"keyword"`
	StartAt  int64    `form:"startAt"`
	EndAt    int64    `form:"endAt"`
}

type UpdateTodoRequest struct {
	Color   string `form:"color"`
	Content string `form:"content"`
	Status  Status `form:"status" binding:"-"`
}
