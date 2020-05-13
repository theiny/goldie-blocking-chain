package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"

	"math"
	"math/big"
	log "github.com/theiny/slog"
)

// difficulty is an arbitrary number to calculate the target. This value would usually be determined and adjusted over time using an algorithm.
const difficulty = 12

type ProofOfWork struct {
	Block  *Block
	Target *big.Int
}

// NewProof takes a block and creates a new proof of work with the given target.
func NewProof(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	pow := &ProofOfWork{b, target}

	return pow
}

// prepareData encapsulates all the data which will eventually get hashed.
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			ToHex(int64(nonce)),
			ToHex(int64(difficulty)),
		},
		[]byte{},
	)

	return data
}

// Run executes the algorithm to sequentially find a nonce that is appropriate to sign the block.
func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	// math.MaxInt64 is used to avoid a possible overflow of nonce.
	for nonce < math.MaxInt64 {
		data := pow.prepareData(nonce)
		hash = sha256.Sum256(data)

		intHash.SetBytes(hash[:])

		// -1 if intHash < pow.Target
		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return nonce, hash[:]
}

// Validate takes the nonce assigned to the block produced by the proof of work, and runs it through the hash function to validate it.
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.prepareData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// ToHex is a helper function to convert an int64 to []byte.
func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Error(err)
	}
	return buff.Bytes()
}
