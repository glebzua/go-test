package database

import (
	"github.com/upper/db/v4"
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
