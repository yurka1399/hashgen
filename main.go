package hashgen

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// HashGenerator Class
type HashGenerator struct {
	key string
}

// Generate generates hash from string passd in toHash
func (h *HashGenerator) Generate(toHash string) string {
	algo := md5.New()
	algo.Write([]byte(fmt.Sprintf("%s%s", toHash, h.key)))
	return hex.EncodeToString(algo.Sum(nil))
}

// Init returns new HashGenerator object
func Init(skey string) *HashGenerator {
	return &HashGenerator{
		key: skey,
	}
}
