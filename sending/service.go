package sending

import "github.com/theiny/goldie-blocking-chain/blockchain"

type Service interface {
	Send(from, to string, amount int) error
}

type service struct {
	bc *blockchain.Blockchain
}

func NewService(bc *blockchain.Blockchain) *service {
	return &service{bc}
}

func (s *service) Send(from, to string, amount int) error {
	tx, err := s.bc.NewTransaction(from, to, amount)
	if err != nil {
		return err
	}

	s.bc.AddBlock([]*blockchain.Transaction{tx})
	return nil
}
