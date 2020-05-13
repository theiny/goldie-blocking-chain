package blockchain

const genesisData = "This is a reference to the genesis block"

// Block defines a since block of the blockchain.
type Block struct {
	Hash         []byte
	Transactions []*Transaction
	PrevHash     []byte
	Nonce        int
}

// Blockchain is made up of a slice of blocks.
type Blockchain struct {
	Blocks []*Block
}

// creates a new block
func newBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}

// AddBlock adds the newly created block to the blockchain.
func (chain *Blockchain) AddBlock(tx []*Transaction) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := newBlock(tx, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis creates the first 'genesis' block of the blockchain.
func Genesis(coinbase *Transaction) *Block {
	return newBlock([]*Transaction{coinbase}, []byte{})
}

// InitBlockChain initializes a new blockchain.
func InitBlockChain(address string) *Blockchain {
	cbtx := CoinbaseTx(address, genesisData)
	genesis := Genesis(cbtx)
	return &Blockchain{[]*Block{genesis}}
}
