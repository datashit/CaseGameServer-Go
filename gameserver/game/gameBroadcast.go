package game

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
)

// broadcastGameObject : verilen gameobjeyi oyunculara yayınlar
// game : oyun bilgisi
// gameOBJ : yayınlanacak obje
func broadcastGameObject(game *Game, gameOBJ *gameObject) {

	buf := new(bytes.Buffer)
	// Nesne byte' a çevriliyor.
	buf.WriteByte(byte(game.ID))
	binary.Write(buf, binary.LittleEndian, gameOBJ.ID)
	binary.Write(buf, binary.LittleEndian, gameOBJ.X)
	binary.Write(buf, binary.LittleEndian, gameOBJ.Y)
	binary.Write(buf, binary.LittleEndian, gameOBJ.VX)
	binary.Write(buf, binary.LittleEndian, gameOBJ.VY)
	binary.Write(buf, binary.LittleEndian, gameOBJ.Timestamp)

	// Nesne oyunculara yollanıyor.
	for _, p := range game.Players {
		sendPacket(p.UDPCon, p.UDPAddr, buf)
	}
}

// sendPacket : verilen parametreler ile udp üzerinden buferı yollar.
// conn : UDP socket
// addr : Yollanıcak UDP adresi
// buf : yollanacak veri
func sendPacket(conn *net.UDPConn, addr *net.UDPAddr, buf *bytes.Buffer) {

	if conn == nil {
		return // socket yok devam etme.
	}
	_, err := conn.WriteToUDP(buf.Bytes(), addr)

	if err != nil {
		fmt.Println("Error:  ", err)
	}
}
