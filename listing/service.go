package listing

import (
	"log"
	"github.com/theiny/goldie-blocking-chain/blockchain"
)

type Service interface {
	GetBlockChain() *blockchain.Blockchain
}

type service struct {
	bc *blockchain.Blockchain
}

func NewService(bc *blockchain.Blockchain) *service {
	return &service{bc}
}

func (s *service) GetBlockChain() *blockchain.Blockchain{
	return s.bc
}

func (s *service) ShowBalance() {
	address := "aRandomAddress"
	balance := 0
	UTXOs := s.bc.FindUTXO(address)

	for _, out := range UTXOs {
		balance += out.Value
	}

	log.Printf("Balance of '%s': %d\n", address, balance)
}
