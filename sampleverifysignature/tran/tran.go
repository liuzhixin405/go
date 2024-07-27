package tran

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type Transaction struct {
	Sender    []byte
	Receiver  []byte
	Amount    float64
	Signature []byte
	DTime     int64
}

func GenerateKeyPair() (*ecdsa.PrivateKey, error) {
	priv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		return nil, err
	}
	return priv, nil
}

func PublicKeyToBytes(pub *ecdsa.PublicKey) []byte {
	return elliptic.MarshalCompressed(elliptic.P256(), pub.X, pub.Y)
}

func BytesToPublicKey(publicBytes []byte) *ecdsa.PublicKey {
	x, y := elliptic.UnmarshalCompressed(elliptic.P256(), publicBytes)
	return &ecdsa.PublicKey{Curve: elliptic.P256(), X: x, Y: y}
}

func SignTransaction(tx *Transaction, priv *ecdsa.PrivateKey) ([]byte, error) {
	txHash := sha256.Sum256([]byte(fmt.Sprintf("%x:%x:%f:%d", tx.Sender, tx.Receiver, tx.Amount, tx.DTime)))
	r, s, err := ecdsa.Sign(rand.Reader, priv, txHash[:])
	if err != nil {
		return nil, err
	}
	signature := append(r.Bytes(), s.Bytes()...)
	return signature, nil
}

func VerifySignature(tx *Transaction, pub *ecdsa.PublicKey) bool {
	txHash := sha256.Sum256([]byte(fmt.Sprintf("%x:%x:%f:%d", tx.Sender, tx.Receiver, tx.Amount, tx.DTime)))
	r := big.Int{}
	s := big.Int{}
	sigLen := len(tx.Signature)
	r.SetBytes(tx.Signature[:(sigLen / 2)])
	s.SetBytes(tx.Signature[(sigLen / 2):])
	return ecdsa.Verify(pub, txHash[:], &r, &s)

}
