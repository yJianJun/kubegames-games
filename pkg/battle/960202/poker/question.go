package poker

var wallet = []int{5}

const price = 5

// 商人有5元 奶茶价格是5元 来了A、B、C三人  A支付10元 B支付15元  C支付20元 找钱流程
func giveMoney(pay int) int {
	change := pay - price
	newWallet := make([]int, 0)
	for _, money := range wallet {
		if money != change {
			newWallet = append(newWallet, money)
		}
	}
	if len(newWallet) == len(wallet) {
		// 如果没有找零则退款
		return pay
	} else {
		// 将付款添加到钱包并找零
		newWallet = append(newWallet, pay)
		wallet = newWallet
		return change
	}
}
