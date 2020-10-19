package hGuest

import (
	"be04gomy/handler"
	"be04gomy/model/mStudent"
)

func StudentList(ctx *handler.Ctx) {
	students, err := mStudent.SelectAll(ctx.Db)
	if ctx.IsError(err) {
		return
	}
	ctx.End(students)
}

func StudentCreate(ctx *handler.Ctx) {
	
}

func StudentUpdate(ctx *handler.Ctx) {
	
}

func StudentDelete(ctx *handler.Ctx) {
	
}
