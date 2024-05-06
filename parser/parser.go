package parser

import (
	"ethparser/internal/api"
	"log"
	"time"
)

// default value of 5s, considering block interval of 15s
const pollIntervalSeconds = 5

type Parser interface {
	// last parsed block
	GetCurrentBlock() int64
	// add address to observer
	Subscribe(address string) bool
	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []api.Transaction
	// whether an address is subscribed
	IsSubscribed(address string) bool
}

func NewParser() *DefaultParser {
	return &DefaultParser{
		subscribedAddresses: make(map[string]struct{}),
		store: &inMemoryStore{
			transactions: make(map[string][]api.Transaction),
		},
	}
}

type DefaultParser struct {
	currentBlock        int64
	subscribedAddresses map[string]struct{}
	store               TransactionStore
}

func (p *DefaultParser) Start() chan struct{} {
	exit := make(chan struct{})
	go func() {
		for {
			select {
			case <-exit:
				return
			case <-time.After(time.Second * pollIntervalSeconds):
				log.Printf("polling...\n")
				currentBlock, err := api.GetEthBlockNumber()
				if err != nil {
					log.Printf("error while getting block number: %v\n", err)
					continue
				}
				if currentBlock != p.currentBlock {
					p.currentBlock = currentBlock
					log.Printf("processing new block: %d\n", currentBlock)
					block, err := api.GetEthBlockByNumber(currentBlock)
					if err != nil {
						log.Printf("error while getting block data: %v\n", err)
						continue
					}
					for _, tx := range block.Result.Transactions {
						p.processTransaction(tx)
					}
				}
			}
		}
	}()
	return exit
}

func (p *DefaultParser) processTransaction(tx api.Transaction) {
	var address string
	if p.IsSubscribed(tx.From) {
		address = tx.From
	} else if tx.To != nil && p.IsSubscribed(*tx.To) {
		address = *tx.To
	}
	if address != "" {
		p.store.ObserveTransactions(address, tx)
	}
}

func (p *DefaultParser) GetCurrentBlock() int64 {
	return p.currentBlock
}

func (p *DefaultParser) Subscribe(address string) bool {
	p.subscribedAddresses[address] = struct{}{}
	return true
}

func (p *DefaultParser) IsSubscribed(address string) bool {
	_, ok := p.subscribedAddresses[address]
	return ok
}

func (p *DefaultParser) GetTransactions(address string) []api.Transaction {
	if _, ok := p.subscribedAddresses[address]; !ok {
		return nil
	}
	return p.store.GetTransactions(address)
}
