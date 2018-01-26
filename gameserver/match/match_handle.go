package match

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"

	"github.com/datashit/CaseGameServer-Go/gameserver/game"
)

var simTCPsize uint64 //bağlı TCP socket sayısı
var incomeMessageJobs = make(chan incomeJobs, 1000)
var partner = make(chan *game.Client)

// CreateMatchHandleWorker : fonksiyonu asenkron matchHandle worker olusturur.
// workerSize : worker sayısı
func CreateMatchHandleWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go matchHandle(incomeMessageJobs)
	}
}

func matchHandle(job <-chan incomeJobs) {
	for j := range job {
		switch j.inMsg.Command {
		case "FIND_GAME":
			fmt.Println("Finding game")
			j.inMsg.Command = "SEARCH_GAME"
			j.encoder.Encode(j.inMsg)
			match(j.inClient)
		default:
		}
	}
}

// match : oyuncu eşleştrime fonksiyonu
// client parametresi : eşleştirlecek oyuncu
func match(client *game.Client) {
	fmt.Printf("Player %v : Waiting for a player... \r\n", client.PlayerID)
	select {
	case partner <- client: // Eşleştirilmek için sıraya alınıyor.
	case p := <-partner: // Oyuncular eşleştirildi.
		g := game.CreateGame(p, client) // Oyun kuruluyor.
		fmt.Println("Game Create: ", g.ID)
		game.AddGame(g) // Oyun pool workera ekleniyor.

		// Oyunun oluşturulduğu bilgisi clientlara bildiriliyor.
		for _, p := range g.Players {
			encoder := json.NewEncoder(*(p.Conn))
			str := fmt.Sprint(g.ID)
			var msg = message{PlayerID: p.PlayerID, Command: "GAME_CREATED", Data: "{\"GameID\":" + str + "}"} //Mesaj oluşturuluyor.
			err := encoder.Encode(msg)                                                                         // Mesaj yollandı.
			if err != nil {
				fmt.Println("Error: ", err)
				return
			}
		}
	}
}

// welcomeMsg : oyuncuya severa bağlandı bilgisi ve oyuncu ID ataması yapar
// client parametresi : atanmış oyuncu kullanıcısı
// encoder parametrei : sokete baglanmış json encoderı
func welcomeMsg(client *game.Client, encoder *json.Encoder) {
	var wel = message{
		PlayerID: client.PlayerID,
		Command:  "WELCOME",
		Data:     "",
	}
	fmt.Printf("Player %v is connected \r\n", wel.PlayerID)
	(*encoder).Encode(wel)
}

// clientReadHandle : Client üzerinden gelen match verilerini dinleyen metod.
// client parametresi : dinlenecek client nesnesi.
func clientReadHandle(client *game.Client) {
	defer (*client.Conn).Close()

	// Reader ve encoder oluşturuluyor.
	encoder := json.NewEncoder(*client.Conn)
	reader := bufio.NewReader(*client.Conn)
	// Client kabul bilgisi yollanıyor.
	welcomeMsg(client, encoder)

	for {
		lineData, isPrefix, err := reader.ReadLine() // Client socket dinleniyor.
		if err == io.EOF {
			fmt.Println("Server disconnected: " + (*client.Conn).RemoteAddr().String())
			return
		}

		if err != nil {
			fmt.Println("Reader error: ", err)
			return
		}

		if isPrefix {

			continue
		}

		var msg message
		err = json.Unmarshal(lineData, &msg) // Gelen metod encode ediliyor.
		if err != nil {
			fmt.Println(err)
			continue
		}

		if msg.PlayerID != client.PlayerID {
			fmt.Println("Player id does not match")
			return // Gelen id ile socket id uyuşmadı.
		}

		fmt.Println(msg)
		// Mesaj işlenmesi için görev oluşturuluyor.
		ijob := incomeJobs{inClient: client, inMsg: msg, encoder: encoder}
		incomeMessageJobs <- ijob
	}
}
