package main

import (
	"bufio"
	"ethparser/internal/api"
	"ethparser/parser"
	"log"
	"os"
	"strings"
)

// simple CLI
func main() {
	ethBlockNumber, err := api.GetEthBlockNumber()
	if err != nil {
		panic(err)
	}
	log.Printf("current block number: %d\n", ethBlockNumber)

	p := parser.NewParser()
	exit := p.Start()

	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			address := strings.ToLower(scanner.Text())
			if p.IsSubscribed(address) {
				txns := p.GetTransactions(address)
				log.Printf("%d transactions recorded for %s\n", len(txns), address)
				//for _, txn := range txns {
				//	fmt.Println(txn)
				//}
			} else {
				p.Subscribe(address)
				log.Printf("subscribed to %s\n", address)
			}
		}

		if scanner.Err() != nil {
			break
		}
	}

	close(exit)
}
