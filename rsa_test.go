package p2p

import "testing"

func TestRSA(t *testing.T) {
	origin := []byte("a special secret message")

	rsa, err := NewRSA()
	if err != nil {
		t.Fatal(err)
	}

	var encoded []byte
	encoded, err = rsa.PublicKey().Encode(origin)
	if err != nil {
		t.Fatal(err)
	}

	var decoded []byte
	decoded, err = rsa.PrivateKey().Decode(encoded)
	if err != nil {
		t.Fatal(err)
	}

	if string(origin) != string(decoded) {
		t.Fatal("Origin and Decoded are not equal")
	}
}
