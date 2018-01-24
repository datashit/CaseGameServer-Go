package gameserver

import (
	"fmt"
	"net"
)

//ServerCore : Server yığını
type ServerCore struct {
	HostAddr       string       // Host soket adresi
	ListenerSocket net.Listener // Dinlenen soket
}

//NewInstance : bir server nesnesi oluşturup geri döner.
func NewInstance(address string) *ServerCore {
	server := &ServerCore{
		HostAddr: address,
	}

	return server
}

//Start : metodu serverı başlatır
func (server *ServerCore) Start() {
	fmt.Println("Server Starting...")
	server.listen()
}

// Shutdown : metodu serverı kapatır
func (server *ServerCore) Shutdown() {
	fmt.Println("Server Shutdown!")

}

func (server *ServerCore) listen() {
	l, err := net.Listen("tcp", server.HostAddr)
	if err != nil {
		fmt.Println("Error: ")
		fmt.Println(err)
		return
	}
	defer l.Close()
	server.ListenerSocket = l

	fmt.Printf("Socket is listening: %s \r\n", server.ListenerSocket.Addr())

	for {
		conn, err := server.ListenerSocket.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Incoming connection request <--")
		// conn handla gönderilmeli
		conn.RemoteAddr()
	}

}
