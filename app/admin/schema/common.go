package schema

type Timestamp struct {
	Start int64 `query:"start"`
	End   int64 `query:"end"`
}

func (v Timestamp) IsZero() bool {
	return v.Start <= 0 && v.End <= 0
}
