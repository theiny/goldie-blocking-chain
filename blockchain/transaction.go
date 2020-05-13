package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"log"
)

const subsidy = 100

type Transaction struct {
	ID      []byte
	Inputs  []TxInput
	Outputs []TxOutput
}

type TxInput struct {
	ID  []byte
	Out int
	Sig string
}

type TxOutput struct {
	Value  int
	PubKey string
}

func (tx *Transaction) setID() {
	var encoded bytes.Buffer
	var hash [32]byte

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	if err != nil {
		log.Println(err)
	}

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// CoinbaseTx creates a seed transaction for the genesis block.
func CoinbaseTx(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := TxInput{[]byte{}, -1, data}
	txout := TxOutput{subsidy, to}

	tx := Transaction{nil, []TxInput{txin}, []TxOutput{txout}}
	tx.setID()

	return &tx
}

func (tx *Transaction) isCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (in *TxInput) canUnlockOutputWith(data string) bool {
	return in.Sig == data
}

// checks if the address matches the output's public key.
func (out *TxOutput) canBeUnlockedWith(data string) bool {
	return out.PubKey == data
}

// HashTransactions produces a single hash representing all transactions within a block.
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}

	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))

	return txHash[:]
}

// FindUnspentTransactions finds outputs of transactions that aren't referenced by inputs of other transactions.
func (bc *Blockchain) FindUnspentTransactions(address string) []Transaction {
	var unspent []Transaction
	spent := make(map[string][]int)

	// need to loop through the blocks in reverse order starting with the latest.
	for i := len(bc.Blocks) - 1; i >= 0; i-- {
		for _, tx := range bc.Blocks[i].Transactions {
			id := hex.EncodeToString(tx.ID)
		Outputs:
			for outID, out := range tx.Outputs {
				// check if an output was already referenced in an input. If so, then skip.
				if spent[id] != nil {
					for _, spentOut := range spent[id] {
						if spentOut == outID {
							continue Outputs
						}
					}
				}

				if out.canBeUnlockedWith(address) {
					unspent = append(unspent, *tx)
				}
			}

			if tx.isCoinbase() == false {
				for _, in := range tx.Inputs {
					if in.canUnlockOutputWith(address) {
						inID := hex.EncodeToString(in.ID)
						spent[inID] = append(spent[inID], in.Out)
					}
				}
			}
		}
	}

	return unspent
}

func (bc *Blockchain) FindUTXO(address string) []TxOutput {
	var UTXOs []TxOutput
	unspent := bc.FindUnspentTransactions(address)

	for _, tx := range unspent {
		for _, out := range tx.Outputs {
			if out.canBeUnlockedWith(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

func (bc *Blockchain) NewTransaction(from, to string, amount int) (*Transaction, error) {
	var inputs []TxInput
	var outputs []TxOutput

	acc, validOutputs := bc.findSpendableOutputs(from, amount)
	if acc < amount {
		return nil, fmt.Errorf("%s does not have enough funds. Current balance: %d", from, acc)
	}

	// Build a list of inputs based on the outputs.
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			return nil, err
		}

		for _, out := range outs {
			input := TxInput{txID, out, from}
			inputs = append(inputs, input)
		}
	}

	// Build a list of outputs
	outputs = append(outputs, TxOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, TxOutput{acc - amount, from}) // a change
	}

	tx := Transaction{nil, inputs, outputs}
	tx.setID()

	return &tx, nil
}

// find all unspent outputs and ensure that they store enough value before creating an input referencing that output.
func (bc *Blockchain) findSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOutputs := make(map[string][]int)
	unspentTXs := bc.FindUnspentTransactions(address)

	accumulated := 0

Work:
	for _, tx := range unspentTXs {
		txID := hex.EncodeToString(tx.ID)

		for outID, out := range tx.Outputs {
			if out.canBeUnlockedWith(address) && accumulated < amount {
				accumulated += out.Value
				unspentOutputs[txID] = append(unspentOutputs[txID], outID)
				if accumulated >= amount {
					break Work
				}
			}
		}
	}

	return accumulated, unspentOutputs
}
