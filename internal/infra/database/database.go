package database

import (
	"github.com/test_server/internal/domain"
	"github.com/upper/db/v4"
	"time"
)

func isDeleted(cond db.LogicalExpr, showDeleted bool) db.LogicalExpr {
	if showDeleted {
		return cond
	}
	delCond := db.Cond{"deletedDate IS": nil}
	if cond == nil {
		return delCond
	} else {
		return db.And(cond, delCond)
	}
}
func getTimePtrFromTime(t time.Time) *time.Time {
	empty := time.Time{}
	if t == empty {
		return nil
	} else {
		return &t
	}
}
func getTimeFromTimePtr(t *time.Time) time.Time {
	if t != nil {
		return *t
	} else {
		return time.Time{}
	}
}
func mapDomainToDbQueryParams(p *domain.UrlQueryParams) *dbQueryParams {
	if p == nil {
		return &dbQueryParams{}
	}
	return &dbQueryParams{
		Page:        p.Page,
		PageSize:    p.PageSize,
		ShowDeleted: p.ShowDeleted,
	}
}

type dbQueryParams struct {
	Page        uint
	PageSize    uint
	ShowDeleted bool
}

func (q *dbQueryParams) ApplyToResult(r db.Result) db.Result {
	if !q.ShowDeleted {
		r = r.And(db.Cond{"deletedDate IS": nil})
	}
	if q.PageSize != 0 {
		r = r.Paginate(q.PageSize).Page(q.Page)
	}
	return r
}
