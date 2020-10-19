package main

import (
	"be04gomy/handler"
	"be04gomy/handler/hGuest"
)

const PORT = `:7000`
func main() {
	server := handler.InitServer()
	server.Handle(`/guest/student/list`,hGuest.StudentList)
	server.Handle(`/guest/student/create`,hGuest.StudentCreate)
	server.Handle(`/guest/student/update`,hGuest.StudentUpdate)
	server.Handle(`/guest/student/delete`,hGuest.StudentDelete)
	server.Listen(PORT) 
}
