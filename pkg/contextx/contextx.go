package contextx

import "context"

type (
	userIdCtx struct{}
	transCtx  struct{}
)

// NewTrans
// @param ctx
// @param trans
// @date 2022-05-20 23:01:09
func NewTrans(ctx context.Context, trans interface{}) context.Context {
	return context.WithValue(ctx, transCtx{}, trans)
}

// FromTrans
// @param ctx
// @date 2022-05-20 23:01:08
func FromTrans(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(transCtx{})
	return v, v != nil
}

func NewUserId(ctx context.Context, userId uint64) context.Context {
	return context.WithValue(ctx, userIdCtx{}, userId)
}

func FromUserID(ctx context.Context) uint64 {
	v := ctx.Value(userIdCtx{})
	if v != nil {
		if s, ok := v.(uint64); ok {
			return s
		}
	}
	return 0
}
