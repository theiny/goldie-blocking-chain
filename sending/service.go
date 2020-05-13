package sending

import "github.com/theiny/goldie-blocking-chain/blockchain"

type Service interface {
	Send(from, to string, amount int)
}

type service struct {
	bc *blockchain.Blockchain
}

func NewService(bc *blockchain.Blockchain) *service {
	return &service{bc}
}

func (s *service) Send(from, to string, amount int) {
	tx := s.bc.NewTransaction(from, to, amount)
	s.bc.AddBlock([]*blockchain.Transaction{tx})
}
