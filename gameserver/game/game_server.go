package game

import (
	"fmt"
	"net"
	"sync/atomic"
)

// Game : Oyun yığını
type Game struct {
	Id        uint64
	Player1ID uint64
	Player2ID uint64
}

var gameIDCounter uint64 //Benzersiz game id sayacı

// CreateGame : oyun olusturur.
func CreateGame(playerid1, playerid2 uint64) *Game {

	game := &Game{
		Id:        atomic.AddUint64(&gameIDCounter, 1), // Benzersiz game id kullanılmalı
		Player1ID: playerid1,
		Player2ID: playerid2,
	}

	return game
}

// StartListen : Udp üzerinden dinlemeye baslar.
func (g *Game) StartListen() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":10001")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	serverConn, err := net.ListenUDP("udp", serverAddr)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer serverConn.Close()

	buf := make([]byte, 1024)
	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
		serverConn.WriteToUDP(buf[0:n], addr)
	}
}
