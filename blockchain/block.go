package blockchain

const (
	genesisData    = "This is a reference to the genesis block"
	GenesisAddress = "adam"
)

// Block defines a since block of the blockchain.
type Block struct {
	Hash         []byte         `json:"hash"`
	Transactions []*Transaction `json:"transactions"`
	PrevHash     []byte         `json:"previous_hash"`
	Nonce        int            `json:"nonce"`
}

// Blockchain is made up of a slice of blocks.
type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

// creates a new block
func newBlock(txs []*Transaction, prevHash []byte) *Block {
	block := &Block{[]byte{}, txs, prevHash, 0}

	pow := newProof(block)
	nonce, hash := pow.run()

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
func createGenesis(coinbase *Transaction) *Block {
	return newBlock([]*Transaction{coinbase}, []byte{})
}

// InitBlockChain initializes a new blockchain.
func InitBlockChain(address string) *Blockchain {
	cbtx := coinbaseTx(address, genesisData)
	genesis := createGenesis(cbtx)
	return &Blockchain{[]*Block{genesis}}
}
