package pagination

import (
	"context"
	"gorm.io/gorm"
	"math"
	"vab-admin/go/pkg/db"
)

// Paging
// @param db
// @param p
// @date 2022-07-29 00:25:28
func Paging[T any](db *gorm.DB, p *Param) (*Paginator[T], error) {

	if p.Page < 1 {
		p.Page = 1
	}

	if p.Limit == 0 {
		p.Limit = 10
	}

	// done := make(chan bool, 1)
	var (
		offset    int
		count     int64
		paginator Paginator[T]
	)

	var result T

	if p.Page == 1 {
		offset = 0
	} else {
		offset = (p.Page - 1) * p.Limit
	}

	join := db.Statement.Joins

	err := db.Limit(p.Limit).Offset(offset).Find(&result).Error
	if err != nil {

		return nil, err
	}

	if len(join) > 0 {
		db.Statement.Joins = join
	}

	countRecords(db, &count)

	paginator.TotalItem = count

	paginator.Page = p.Page
	paginator.Items = result
	paginator.TotalPage = int(math.Ceil(float64(count) / float64(p.Limit)))

	return &paginator, err
}

// countRecords
// @param tx
// @param done
// @param count
// @date 2022-07-29 00:25:27
func countRecords(tx *gorm.DB, count *int64) {

	ctx := context.Background()

	query := db.Session(ctx)
	if m := tx.Statement.Model; m != nil {
		query.Statement.Model = tx.Statement.Model
	}

	if table := tx.Statement.Table; table != "" {
		query.Statement.Table = tx.Statement.Table
	}

	if tableExpr := tx.Statement.TableExpr; tableExpr != nil {
		query.Statement.TableExpr = tableExpr
	}

	if joins := tx.Statement.Joins; len(joins) > 0 {

		query.Statement.Joins = joins
	}

	if where := tx.Statement.Clauses["WHERE"]; where.Name != "" {
		query.Statement.Clauses["WHERE"] = where
	}

	if groupBy := tx.Statement.Clauses["GROUP BY"]; groupBy.Name != "" {
		query.Statement.Clauses["GROUP BY"] = groupBy
	}

	query.Count(count)
}
