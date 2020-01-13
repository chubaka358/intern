package factory_method

import "testing"

func TestCash(t *testing.T) {
	cash := NewCash()
	if cash == nil {
		t.Fatalf("successful creation of the object was expected")
	}
	if cash.Balance() != 0 {
		t.Errorf("expected 1000 on balance")
	}

	cash.Replenish(1000)
	if cash.Balance() != 1000 {
		t.Errorf("expected 1000 on balance")
	}

	cash.Pay(300)
	if cash.Balance() != 700 {
		t.Errorf("expected 700 on balance")
	}
}

func TestCreditCard(t *testing.T) {
	creditCard := NewCreditCard()
	if creditCard == nil {
		t.Fatalf("successful creation of the object was expected")
	}
	if creditCard.Balance() != 0 {
		t.Errorf("expected 1000 on balance")
	}

	creditCard.Replenish(1000)
	if creditCard.Balance() != 1000 {
		t.Errorf("expected 1000 on balance")
	}

	creditCard.Pay(300)
	if creditCard.Balance() != 700 {
		t.Errorf("expected 700 on balance")
	}
}

func TestBitcoin(t *testing.T) {
	bitcoin := NewBitcoin()
	if bitcoin == nil {
		t.Fatalf("successful creation of the object was expected")
	}
	if bitcoin.Balance() != 0 {
		t.Errorf("expected 1000 on balance")
	}

	bitcoin.Replenish(1000)
	if bitcoin.Balance() != 1000 {
		t.Errorf("expected 1000 on balance")
	}

	bitcoin.Pay(300)
	if bitcoin.Balance() != 700 {
		t.Errorf("expected 700 on balance")
	}
}
