package poker

import (
	"testing"
)

func TestProcessPayment(t *testing.T) {
	seller := NewSeller()
	payment := seller.ProcessPayment(20)
	t.Log(payment)
}
