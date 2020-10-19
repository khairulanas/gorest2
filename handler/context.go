package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Ctx struct {
	http.ResponseWriter
	*Server
	Request *http.Request
}

func (c *Ctx) IsError(err error) bool {
	if err != nil {
		c.ResponseWriter.WriteHeader(200)
		c.ResponseWriter.Header().Set(`content-type`,`application/json`)
		fmt.Println(err)
		js, err := json.Marshal(map[string]string{`error`:err.Error()})
		if err != nil {
			fmt.Println(err)
		}
		c.ResponseWriter.Write(js)
		return true
	}
	return false
}

func (c *Ctx) End(any interface{}) {
	js, err := json.Marshal(any)
	if err !=nil {
		fmt.Println(err)
	}
	c.ResponseWriter.WriteHeader(200)
	c.ResponseWriter.Header().Set(`content-type`,`application/json`)
	c.ResponseWriter.Write(js)
}
