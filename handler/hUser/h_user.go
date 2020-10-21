package hUser

import (
	"be04gomy/handler"
	"be04gomy/model/mAntrian"
)

func AntrianCreate(ctx *handler.Ctx) {
	if ctx.Request.Method == `POST` {
		antrian, err := mAntrian.Create(ctx.Db)
		if ctx.IsError(err) {
			return
		}
		ctx.End(antrian)
	}
}

func AntrianGetlast(ctx *handler.Ctx) {
	nomor, err := mAntrian.Getlast(ctx.Db)
	if ctx.IsError(err) {
		return
	}
	ctx.End(nomor)
}
