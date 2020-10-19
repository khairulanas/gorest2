package main

const PORT = `:7000`
func main() {
	server := InitServer()
	server.Listen(PORT) 
}
