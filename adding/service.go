package adding

import "github.com/theiny/goldie-blocking-chain/blockchain"

type Service interface {
	AddBlock(data []byte)
}

type service struct {
	bc *blockchain.BlockChain
}

func NewService(bc *blockchain.BlockChain) *service {
	return &service{bc}
}

func (s *service) AddBlock(data []byte) {
	s.bc.AddBlock(data)
}
