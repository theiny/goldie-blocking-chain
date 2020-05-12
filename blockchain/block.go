package blockchain

// Block defines a since block of the blockchain.
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

// BlockChain is made up of a slice of blocks.
type BlockChain struct {
	Blocks []*Block
}

type Iterator struct {
	currentHash []byte
	db          *BlockChain
}

func (i *Iterator) Next() *Block {
	var block *Block

	return block
}

// creates a new block
func createBlock(data []byte, prevHash []byte) *Block {
	block := &Block{[]byte{}, data, prevHash, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock adds the newly created block to the blockchain.
func (chain *BlockChain) AddBlock(data []byte) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis creates the first 'genesis' block of the blockchain.
func Genesis() *Block {
	return createBlock([]byte("Genesis"), []byte{})
}

// InitBlockChain initializes a new blockchain.
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
