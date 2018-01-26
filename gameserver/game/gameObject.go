package game

// GameObject : Sahte oyun yığını test amaçlı
type gameObject struct {
	ID        uint32  // Obje ID
	X         float32 // X kordinatı
	Y         float32 // Y kordinatı
	VX        float32 // Velocity X
	VY        float32 // Velocity Y
	Timestamp int32   // Zaman bilgisi
}
