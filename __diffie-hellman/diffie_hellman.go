package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// PrivateKey returns 'a'
func PrivateKey(p *big.Int) *big.Int {
	nBig, err := rand.Int(rand.Reader, p)
	if err != nil {
		panic(err)
	}
	a := nBig.Add(nBig, big.NewInt(1))
	if a.Cmp(big.NewInt(1)) <= 0 || p.Cmp(a) <= 0 {
		return PrivateKey(p)
	}
	return a
}

func PublicKey(a, p *big.Int, g int64) *big.Int {
	var temp, y big.Int
	temp.SetInt64(g)
	A := *y.Exp(&temp, a, p)
	return &A
}

func SecretKey(a, b, p *big.Int) *big.Int {
	var temp big.Int
	s := *temp.Exp(b, a, p)
	return &s
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	a := PrivateKey(p)
	A := PublicKey(a, p, g)
	return a, A
}
