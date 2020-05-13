package listing

import (
	"github.com/theiny/goldie-blocking-chain/blockchain"
)

type Service interface {
	GetBlockChain() *blockchain.Blockchain
	GetBalance(addr string) int
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

func (s *service) GetBalance(addr string) int {
	balance := 0
	UTXOs := s.bc.FindUTXO(addr)

	for _, out := range UTXOs {
		balance += out.Value
	}

	return balance
}
