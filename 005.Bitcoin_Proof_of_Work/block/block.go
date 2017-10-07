package block

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
)

// NonceData returns a string created by concatenation of data and nonce
func (o *Block) NonceData() string {
	return string(append(o.Data, []byte(strconv.FormatUint(o.Nonce, 10))...))
}

// Block A prof of work representation
type Block struct {
	Nonce uint64
	Data  []byte
	Hash  string
}

// Mine returns a mined block which data hash has got n leading zeros
func Mine(data []byte, prefix string) Block {
	var hash [32]byte
	nonce := uint64(0)
	bPrefix, err := hex.DecodeString(prefix)

	if err != nil {
		panic(err)
	}

	for {
		nonceData := append(data, []byte(strconv.FormatUint(nonce, 10))...)
		hash = sha256.Sum256(nonceData)

		if bytes.HasPrefix(hash[:], bPrefix) {
			break
		}
		nonce++
	}

	return Block{
		Nonce: nonce,
		Data:  data,
		Hash:  string(hash[:])}
}
