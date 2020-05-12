package listing

import "github.com/theiny/goldie-blocking-chain/blockchain"

type Service interface {
	GetBlockChain() *blockchain.BlockChain
}

type service struct {
	bc *blockchain.BlockChain
}

func NewService(bc *blockchain.BlockChain) *service {
	return &service{bc}
}

func (s *service) GetBlockChain() *blockchain.BlockChain {
	return s.bc
}
