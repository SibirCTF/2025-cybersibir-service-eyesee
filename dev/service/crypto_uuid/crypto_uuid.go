package __

import (
	"crypto/rand"
	"crypto/sha512"
	"fmt"
	"io"
	"math/big"
)

func NewDetermRand(seed []byte) io.Reader {
	return &DetermRand{next: seed}
}

type DetermRand struct {
	next []byte
}

func (d *DetermRand) cycle() []byte {
	result := sha512.Sum512(d.next)
	d.next = result[:sha512.Size/2]
	return result[sha512.Size/2:]
}

func (d *DetermRand) Read(b []byte) (int, error) {
	n := 0
	for n < len(b) {
		out := d.cycle()
		n += copy(b[n:], out)
	}
	return n, nil
}

func UUIDFromInt(int *big.Int) string {
	HexStr := fmt.Sprintf("%032x", int)
	//         return '%s-%s-%s-%s-%s' % (
	//            hex[:8], hex[8:12], hex[12:16], hex[16:20], hex[20:])
	return fmt.Sprintf(
		"%s-%s-%s-%s-%s",
		HexStr[:8],
		HexStr[8:12],
		HexStr[12:16],
		HexStr[16:20],
		HexStr[20:],
	)
}

func CryptoUUID(reader io.Reader) string {
	MaxInt := new(big.Int)
	MaxInt.SetString("340282366920938463463374607431768211454", 10)
	r, _ := rand.Int(reader, MaxInt)
	return UUIDFromInt(r)
}
