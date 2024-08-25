package game

import (
	"github.com/kubegames/kubegames-games/pkg/battle/landlord/msg"
)

// FlowerType 表示花的类型。它是一个整数类型。
// 使用示例：
// 钻石 FlowerType = iota + 1
type FlowerType int

// landlord 从扑克牌开始
const (
	//方块
	DIAMOND FlowerType = iota + 1
	//梅花
	PLUM
	//红桃
	HEART
	//黑桃
	SPADE
)

// CardPoint代表卡点值。它是一个字节类型。
//
// 用法示例：
// C3 CardPoint = iota + 1
type CardPoint byte

const (
	C3 CardPoint = iota + 1
	C4
	C5
	C6
	C7
	C8
	C9
	C10
	CJ
	CQ
	CK
	CA
	C2

	// CBlackJoker 表示黑色小丑牌的常量值。
	// 它的值为 0xe1。
	//
	// 使用示例：
	// 如果 g.Cards[0]>>4 == byte(CBlackJoker) {
	// // 做某事
	// }
	CBlackJoker = 0xe1

	// CRedJoker 代表红色小丑牌点值。它是 CardPoint 类型的常量。
	CRedJoker = 0xf1
)

// Sort111 根据提供的卡片类型对输入字节切片进行排序。
// 该函数返回带有已排序卡片的新字节片。如果卡片类型
// 无法识别或者不需要排序，函数返回原始切片。
//
// 该函数遵循以下排序规则：
// 1. 对于卡片类型 msg.CardsType_TripletWithSingle，如果输入切片的第一张和第二张卡片
// 有不同的花色，卡片的排序方式是第一张和第四张卡片交换位置。
// 2. 对于卡片类型 msg.CardsType_TripletWithPair，如果输入切片的第一张和第三张卡片有
// 不同花色，牌按照第一张和第二张三联牌的放置方式排序
// 在结果切片的末尾，后面是第一对和第二对卡片。
// 3. 对于卡片类型 msg.CardsType_SerialTripletWithOne 和 msg.CardsType_SerialTripletWithWing，该函数
// 在输入切片中查找由三张卡片组成的序列，并将该序列附加到结果切片中。
// 然后将剩余的卡片按以下顺序附加：
// - 如果序列从索引 0 开始，则附加该序列之后的剩余卡片。
// - 如果序列从输入切片中间的某个位置开始并在最后一张卡片之前结束，
// 序列之前的卡片和序列之后的卡片都追加到结果切片中。
// - 如果序列从输入切片中间的某个位置开始并在最后一张卡片处结束，
// 序列之前的卡片将附加到结果切片中。
// 4. 对于 msg.CardsType_QuartetWithTwo 类型的卡片，该函数在输入中查找四张卡片的序列
// 切片，并将序列附加到结果切片中。然后按以下顺序附加剩余的卡片：
// - 如果序列从索引 0 开始，则附加该序列之后的剩余卡片。
// - 如果序列从输入切片中间的某个位置开始并在最后一张卡片之前结束，
// 序列之前的卡片和序列之后的卡片都追加到结果切片中。
// - 如果序列从输入切片中间的某个位置开始并在最后一张卡片处结束，
// 序列之前的卡片将附加到结果切片中。
// 如果以上卡片类型均不匹配，则该函数返回原始输入切片。
//
// 该函数不会修改输入切片。
func Sort111(c []byte, cardsType msg.CardsType) []byte {
	var res []byte
	switch cardsType {
	case msg.CardsType_TripletWithSingle:
		if c[0]>>4 != c[1]>>4 {
			res = make([]byte, len(c))
			copy(res, c)
			res[0], res[3] = res[3], res[0]
		}
	case msg.CardsType_TripletWithPair:
		if c[0]>>4 != c[2]>>4 {
			res = make([]byte, 0)
			res = append(res, c[2:5]...)
			res = append(res, c[0:2]...)
		}
	case msg.CardsType_SerialTripletWithOne:
		fallthrough
	case msg.CardsType_SerialTripletWithWing:
		threeIdx := 0
		for i := range c {
			if c[i]>>4 == c[i+1]>>4 && c[i+1]>>4 == c[i+2]>>4 {
				threeIdx = i
				break
			}
		}
		threeEndIdx := 0
		for i := threeIdx; i < len(c)-2; {
			if c[i]>>4 == c[i+1]>>4 && c[i+1]>>4 == c[i+2]>>4 {
				i += 3
				threeEndIdx = i
				continue
			} else {
				break
			}
		}

		res = make([]byte, 0)
		res = append(res, c[threeIdx:threeEndIdx]...)
		if threeIdx == 0 {
			res = append(res, c[threeEndIdx:]...)
		} else if threeIdx > 0 && threeEndIdx < len(c) {
			res = append(res, c[0:threeIdx]...)
			res = append(res, c[threeEndIdx:]...)
		} else if threeIdx > 0 && threeEndIdx == len(c) {
			res = append(res, c[0:threeIdx]...)
		}
	case msg.CardsType_QuartetWithTwo:
		fourIdx := 0
		for i := range c {
			if c[i]>>4 == c[i+1]>>4 && c[i+1]>>4 == c[i+2]>>4 && c[i+2]>>4 == c[i+3]>>4 {
				fourIdx = i
				break
			}
		}
		res = make([]byte, 0)
		res = append(res, c[fourIdx:fourIdx+4]...)
		if fourIdx == 0 {
			res = append(res, c[fourIdx+4:]...)
		} else if fourIdx > 0 && fourIdx < len(c) {
			res = append(res, c[:fourIdx]...)
			res = append(res, c[fourIdx+4:]...)
		} else {
			res = append(res, c[:fourIdx]...)
		}

	default:

		return c
	}

	return res
}

// CardsType 表示一副牌中的牌的类型。它是一个整数类型。
// 使用 CardsType 定义一副牌中不同类型的牌。
// 用法示例：
// var ct CardsType = 0
// ct = 1
// ct = 2
// ct = 3
// ct = 4
type CardsType int

// 取较小的整数
func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 取较大的整数
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 取较小的整数
func MinInt64(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

// 取较大的整数
func MaxInt64(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
