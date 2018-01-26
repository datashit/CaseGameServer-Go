package security

import "testing"

var noCryptodataTest = []byte{0xF, 0x5, 0x0, 0xFF}

func TestMessageNoCryptoEncode(t *testing.T) {
	nocpt := &Nocrypto{}
	encode := nocpt.Encode(noCryptodataTest)

	if len(noCryptodataTest) != len(encode) {
		t.Error("NoCrypto encode function: data len is not equal")
	}

	for i := 0; i < len(encode); i++ {
		if encode[i] != noCryptodataTest[i] {
			t.Error("NoCrypto encode function: data is not equal")
		}
	}

}

func TestMessageNoCryptoDecode(t *testing.T) {
	nocpt := &Nocrypto{}

	decode := nocpt.Decode(noCryptodataTest)

	if len(noCryptodataTest) != len(decode) {
		t.Error("NoCrypto decode function: data len is not equal")
	}

	for i := 0; i < len(decode); i++ {
		if decode[i] != noCryptodataTest[i] {
			t.Error("NoCrypto decode function: data is not equal")
		}
	}

}
