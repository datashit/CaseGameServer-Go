package gameserver

import "fmt"

// Server yığını
// Server bilgilerini tutar
type ServerCore struct {
}

// Bir server instance olusturup geri doner.
func NewInstance(address string) *ServerCore {
	server := &ServerCore{}

	return server
}

// Serveri baslatir
func (server *ServerCore) Start() {
	fmt.Println("Server Starting...")
}

func (server *ServerCore) Shutdown() {
	fmt.Println("Server Shutdown!")
}
