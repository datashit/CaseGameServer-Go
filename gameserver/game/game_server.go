package game

import (
	"encoding/json"
	"fmt"
	"net"
	"sync/atomic"
	"time"
)

var gameIDCounter uint32                               //Benzersiz game ID sayacı
var incomeGameMessageJobs = make(chan IncomeJob, 1000) // Gelen udp mesaj işleme kanalı
var gamesRunJob = make(chan *Game, 1000)               // Oyun Run kanalı

// Game : Oyun yığını
type Game struct {
	ID          uint32        // Oyun ID
	Players     []*Client     // Oyuncu listesi
	GameObjects []*gameObject // Oyun Objesi listesi
}

// IncomeGameMessage : UDP gelen mesaj şablonu
type IncomeGameMessage struct {
	GameID   uint32 // Oyun ID
	PlayerID uint64 // Oyuncu ID
	Command  string // Komut
	Data     string // Komut Verisi
}

// IncomeJob : UDP gelen görev işleme şablonu
type IncomeJob struct {
	conn *net.UDPConn      // UDP socket
	addr *net.UDPAddr      // Komutu yollayan udp adresi
	msg  IncomeGameMessage // Gelen komut mesajı
}

// CreateGame : oyun olusturur.
func CreateGame(player1, player2 *Client) *Game {

	g := &Game{
		ID:          atomic.AddUint32(&gameIDCounter, 1), // Benzersiz game ID ataması yapılıyor.
		Players:     make([]*Client, 0),
		GameObjects: make([]*gameObject, 0),
	}

	// Oyuncu bilgisi oyuna ekleniyor.
	g.Players = append(g.Players, player1)
	g.Players = append(g.Players, player2)

	// Test için 2 adet gameobject ekleniyor.
	g.GameObjects = append(g.GameObjects, &gameObject{ID: 0, X: 0, Y: 0})
	g.GameObjects = append(g.GameObjects, &gameObject{ID: 1, X: 10, Y: 0})

	return g
}

// Games : aktif oyun dizisi
var Games map[uint32]*Game

// InitGames : fonksiyonu oyunlar döngü işleyicisini aktif eder.
func InitGames() {
	Games = make(map[uint32]*Game)
	go GamesLoop()
}

// AddGame : fonksiyonu oyunu run moda geçmesi için işleme listesine ekler.
func AddGame(g *Game) {
	Games[g.ID] = g
}

// GamesLoop : işleme zamanı geldiğinde oyunları run mode işleme kanalına aktarır.
func GamesLoop() {
	for _ = range time.Tick(time.Second / 60) {
		for _, g := range Games {
			gamesRunJob <- g // Saniyede 60 kere oyunları işlenmesi için göreve eklenir.
		}
	}

}

// CreateGameRunWorker : Oyun işleyici pool workerlar oluşturur.
// workerSize : worker sayısı
func CreateGameRunWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go gameRunHandle(gamesRunJob)
	}
}

func gameRunHandle(job <-chan *Game) {
	for j := range job {
		j.Run() // Oyun işleniyor.
	}
}

// Run : oyun anadöngüsünün koştuğu metod
func (g *Game) Run() {

	for _, p := range g.Players {
		if p.UDPAddr == nil {
			return // Oyuncular hazır değilse devam etme.
		}
	}

	//Update islemleri burada yapılıyor.

	// Oyun Playerlara Broadcast ediliyor.
	for _, obj := range g.GameObjects {
		broadcastGameObject(g, obj)
	}

}

// StartGameUDPListen : Udp üzerinden oyunu dinlemeye baslar.
func StartGameUDPListen(address string) {
	serverAddr, err := net.ResolveUDPAddr("udp", address) // UDP Listen address oluşturuluyor.
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	serverConn, err := net.ListenUDP("udp", serverAddr) // UDP Listen socket oluşturuluyor.
	if err != nil {
		fmt.Println("Error: ", err)
	}
	defer serverConn.Close()

	buf := make([]byte, 1024)
	for {
		n, udpAddr, err := serverConn.ReadFromUDP(buf) // UDP Socket dinleniyor.
		fmt.Println("Received ", string(buf[0:n]), " from ", udpAddr)

		if err != nil {
			fmt.Println("Error: ", err)
		}

		// Gelen data işleniyor.
		var inMsg IncomeGameMessage
		err = json.Unmarshal(buf[0:n], &inMsg)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// Görev oluşturuluyor.
		var job = IncomeJob{conn: serverConn, addr: udpAddr, msg: inMsg}

		// Görev işlenmesi için channela aktarılıyor.
		incomeGameMessageJobs <- job
	}
}

// CreateGameIncomeMessageWorker : fonksiyonu asenkron udp income görev işleyici  worker olusturur.
// workerSize : worker sayısı
func CreateGameIncomeMessageWorker(workerSize int) {
	for w := 1; w <= workerSize; w++ {
		go incomeMessageHandle(incomeGameMessageJobs)
	}
}

// incomeMessageHandle : UDP üzerinden gelen görevlerinin işlendiği handle
// job : channel olarak görev alır.
// daha önce oluşturulmuş workerlar bu handle' ı işler.
func incomeMessageHandle(job <-chan IncomeJob) {
	for j := range job {

		switch j.msg.Command {
		case "CONNECT_GAME": // Oyuncu oyun ile eşleştirilir.
			for _, p := range Games[j.msg.GameID].Players {
				if p.PlayerID == j.msg.PlayerID {
					p.UDPCon = j.conn
					p.UDPAddr = j.addr
				}
			}
		case "GAMEOBJECT":
			// Gelen obje bilgisi burada işlenir.
		default:
		}
	}
}
