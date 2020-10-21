package hTeller

import (
	"be04gomy/handler"
	"be04gomy/model/mAntrian"
	"encoding/json"
)

func AntrianUpdate(ctx *handler.Ctx) {
	if ctx.Request.Method == `PUT` {
		m := mAntrian.Antrian{}
		err := json.NewDecoder(ctx.Request.Body).Decode(&m)
		if ctx.IsError(err) {
			return
		}
		affectedRow, err := mAntrian.Update(ctx.Db, &m)
		if ctx.IsError(err) {
			return
		}
		ctx.End(affectedRow)
	}
}

func AntrianGetlast(ctx *handler.Ctx) {
	nomor, err := mAntrian.Getlast(ctx.Db)
	if ctx.IsError(err) {
		return
	}
	ctx.End(nomor)
}
