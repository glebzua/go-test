package database

import (
	"github.com/upper/db/v4"
	"log"
)

func isDeleted(cond db.LogicalExpr, showDeleted bool) db.LogicalExpr {
	if showDeleted {
		log.Print("showDeleted", showDeleted, "cond", cond)
		return cond
	}
	delCond := db.Cond{"deleted_date IS": nil}
	if cond == nil {
		log.Print("delCond", delCond, "cond", cond)
		return delCond
	} else {
		log.Print("else delCond", delCond, "cond", cond)
		return db.And(cond, delCond)
	}
}
