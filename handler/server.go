package handler

import (
	"be04gomy/config"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	Db *sql.DB
}

func (s *Server) Handle(route string,f func(*Ctx)) {
	http.HandleFunc(route,func(writer http.ResponseWriter, request *http.Request) {
		f(&Ctx{
			Server: s,
			ResponseWriter: writer,
			Request: request,
		})
	})
	
} 

func (s *Server) Listen(port string) {
	fmt.Println(`listen server at `+port)
	http.ListenAndServe(port,nil)
}

func InitServer() *Server {
	db, err := config.ConnectMysql()
	if err != nil {
		log.Fatal(err)
	}
	return &Server{
		Db: db,	
	}
}
