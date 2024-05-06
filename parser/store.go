package parser

import "ethparser/internal/api"

type TransactionStore interface {
	ObserveTransactions(address string, tx ...api.Transaction)
	GetTransactions(address string) []api.Transaction
}

type inMemoryStore struct {
	transactions map[string][]api.Transaction
}

func (i *inMemoryStore) ObserveTransactions(address string, tx ...api.Transaction) {
	txs, _ := i.transactions[address]
	i.transactions[address] = append(txs, tx...)
}

func (i *inMemoryStore) GetTransactions(address string) []api.Transaction {
	return i.transactions[address]
}
