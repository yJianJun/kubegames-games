package poker

var wallet = []int{5}

const price = 5

// 商人有5元 奶茶价格是5元 来了A、B、C三人  A支付10元 B支付15元  C支付20元 找钱流程
func giveMoney(pay int) int {
	change := pay - price
	if change < 0 {
		return pay
	}
	newWallet := make([]int, 0)
	for i := 0; i < len(wallet); i++ {
		if wallet[i] != change {
			newWallet = append(newWallet, wallet[i])
		}
	}
	if len(newWallet) == len(wallet) {
		return pay
	}
	newWallet = append(newWallet, pay)
	wallet = newWallet
	return change
}
