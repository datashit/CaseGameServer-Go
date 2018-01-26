package security

type Imessagecrypto interface {
	Decode(data []byte) []byte // decode metodu dayı kriptolar.
	Encode(data []byte) []byte // encode metodu kriptolanmis datanın kriptosunu çözer.
}

type Nocrypto struct {
}

func (cpt *Nocrypto) Decode(data []byte) []byte {
	return data
}

func (cpt *Nocrypto) Encode(data []byte) []byte {
	return data
}
