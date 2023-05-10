package pagination

type (
	Paginator[T any] struct {
		TotalItem int64 `json:"totalItem"`
		TotalPage int   `json:"totalPage"`
		Page      int   `json:"page"`
		Items     T     `json:"items"`
	}

	Param struct {
		Limit int `query:"limit"`
		Page  int `query:"page"`
	}
)
