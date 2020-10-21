package main

import (
	"be04gomy/handler"
	"be04gomy/handler/hGuest"
	"be04gomy/handler/hTeller"
	"be04gomy/handler/hUser"
)

const PORT = `:7000`

func main() {
	server := handler.InitServer(`views/`)
	server.Handle(`/guest/student/list`, hGuest.StudentList)
	server.Handle(`/guest/student/create`, hGuest.StudentCreate)
	server.Handle(`/guest/student/update`, hGuest.StudentUpdate)
	server.Handle(`/guest/student/delete`, hGuest.StudentDelete)

	//khairul anas
	server.Handle(`/user/antrian/create`, hUser.AntrianCreate)
	server.Handle(`/user/antrian/getlast`, hUser.AntrianGetlast)
	server.Handle(`/teller/antrian/update`, hTeller.AntrianUpdate)
	server.Handle(`/teller/antrian/getlast`, hTeller.AntrianGetlast)
	server.Listen(PORT)
}
