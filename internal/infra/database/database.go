package database

import "github.com/upper/db/v4"

func softDelCond(cond db.LogicalExpr, showDeleted bool) db.LogicalExpr {
	if showDeleted {
		return cond
	}
	delCond := db.Cond{"deleted_date IS": nil}
	if cond == nil {
		return delCond
	} else {
		return db.And(cond, delCond)
	}
}
