package poker

type Seller struct {
	wallet []int
}

func NewSeller() *Seller {
	return &Seller{wallet: []int{5}}
}

// 商人有5元 奶茶销售5元 来了A、B、C三人  A支付10元 B支付15元  C支付20元 找钱流程
const CostOfTea = 5 // 引入茶成本常数

// ProcessPayment 在收到付款并找零后重新计算卖家的钱包
func (s *Seller) ProcessPayment(payment int) int {
	changeNeeded := payment - CostOfTea
	newWallet := make([]int, 0) // 为了清晰起见，重命名为 newWallet

	for _, bill := range s.wallet {
		if bill != changeNeeded {
			newWallet = append(newWallet, bill)
		}
	}
	s.wallet = newWallet // 支付过程结束后重新计算钱包
	return changeNeeded
}
