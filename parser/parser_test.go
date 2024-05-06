package parser

import (
	"ethparser/internal/api"
	"reflect"
	"testing"
)

func TestDefaultParser(t *testing.T) {
	address1 := "foo"
	address2 := "bar"
	t1 := api.Transaction{From: address1, Value: "0x1", S: "some data"}
	t2 := api.Transaction{From: address2, To: &address1, Value: "0x2", S: "some more data"}

	p := NewParser()
	p.processTransaction(t1)

	// at this point, address1 wasn't subscribed, so t1 shouldn't be recorded
	if got := p.GetTransactions(address1); len(got) > 0 {
		t.Errorf("GetTransactions should not return any transactions on unsubscribed addresses")
	}

	p.Subscribe(address1)
	p.processTransaction(t1)

	// after subscribing to address1, we expect to see t1 recorded as an outgoing txn
	if got := p.GetTransactions(address1); len(got) != 1 || !reflect.DeepEqual(got[0], t1) {
		t.Errorf("GetTransactions did not return transactions on subscribed addresses")
	}

	// t2 should also be recorded as an incoming txn
	p.store.ObserveTransactions(address1, t2)

	if got := p.GetTransactions(address1); len(got) != 2 || !reflect.DeepEqual(got[1], t2) {
		t.Errorf("GetTransactions did not return transactions on subscribed addresses ")
	}
}
