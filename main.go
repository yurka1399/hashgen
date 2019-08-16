package hashgen

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash"
)

// HashGenerator Class
type HashGenerator struct {
	key  string
	algo hash.Hash
}

// Generate generates hash from string passd in toHash
func (h *HashGenerator) Generate(toHash string) string {
	h.algo.Write([]byte(fmt.Sprintf("%s%s", toHash, h.key)))
	return hex.EncodeToString(h.algo.Sum(nil))
}

// Init returns new HashGenerator object
func Init(skey string) *HashGenerator {
	return &HashGenerator{
		key:  skey,
		algo: md5.New(),
	}
}
