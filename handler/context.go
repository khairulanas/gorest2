package handler

import "net/http"

type Ctx struct {
	http.ResponseWriter
	Server *Server
	Request *http.Request
}
