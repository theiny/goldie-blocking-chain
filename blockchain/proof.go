package blockchain

import (
	"bytes"
	"log"
	"crypto/sha256"
	"encoding/binary"

	"math"
	"math/big"
)

// difficulty is an arbitrary number to calculate the target. This value would usually be determined and adjusted over time using an algorithm.
const difficulty = 12

type proofOfWork struct {
	Block  *Block
	Target *big.Int
}

// newProof takes a block and creates a new proof of work with the given target.
func newProof(b *Block) *proofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))

	pow := &proofOfWork{b, target}

	return pow
}

// prepareData encapsulates all the data which will eventually get hashed.
func (pow *proofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransactions(),
			toHex(int64(nonce)),
			toHex(int64(difficulty)),
		},
		[]byte{},
	)

	return data
}

// executes the algorithm to sequentially find a nonce that is appropriate to sign the block.
func (pow *proofOfWork) run() (int, []byte) {
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
func (pow *proofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.prepareData(pow.Block.Nonce)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

// toHex is a helper function to convert an int64 to []byte.
func toHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Println(err)
	}
	return buff.Bytes()
}
