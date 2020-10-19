package main

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
