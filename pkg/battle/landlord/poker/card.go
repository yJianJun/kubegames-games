package poker

import (
	"github.com/kubegames/kubegames-games/internal/pkg/rand"
	"github.com/kubegames/kubegames-games/pkg/battle/landlord/msg"
	"github.com/kubegames/kubegames-sdk/pkg/log"
)

// GamePoker 牌堆
type GamePoker struct {
	// 1个字节表示1张牌 字节数组表示1手牌
	Cards []byte
}

// Deck 是一个变量，表示字节数组中的一副纸牌。
// 这是一副标准的 54 张牌，顺序从高到低：
// 大王, 小王, 2, A, K, Q, J, 10, 9, 8, 7, 6, 5, 4, 3,
// 方1, 梅1, 红1, 黑1,
// 方2, 梅2, 红2, 黑2,
// 方3, 梅3, 红3, 黑3,
// 方4, 梅4, 红4, 黑4,
// 方5, 梅5, 红5, 黑5,
// 方6, 梅6, 红6, 黑6,
// 方7, 梅7, 红7, 黑7,
// 方8, 梅8, 红8, 黑8,
// 方9, 梅9, 红9, 黑9,
// 方10, 梅10, 红10, 黑10,
// 方J, 梅J, 红J, 黑J,
// 方Q, 梅Q, 红Q, 黑Q,
// 方K, 梅K, 红K, 黑K
var Deck = []byte{
	//3 , 4   , 5   , 6   , 7   , 8   , 9   , 10  , J   , Q   , K   , A   , 2   , 小王 , 大王
	0x11, 0x21, 0x31, 0x41, 0x51, 0x61, 0x71, 0x81, 0x91, 0xa1, 0xb1, 0xc1, 0xd1, 0xe1, 0xf1,
	0x12, 0x22, 0x32, 0x42, 0x52, 0x62, 0x72, 0x82, 0x92, 0xa2, 0xb2, 0xc2, 0xd2,
	0x13, 0x23, 0x33, 0x43, 0x53, 0x63, 0x73, 0x83, 0x93, 0xa3, 0xb3, 0xc3, 0xd3,
	0x14, 0x24, 0x34, 0x44, 0x54, 0x64, 0x74, 0x84, 0x94, 0xa4, 0xb4, 0xc4, 0xd4,
}

// Unit 是一个字节片，包含表示一个单元的一系列值。
var Unit = []byte{
	0x1, 0x2, 0x3, 0x4, 0x5, 0x6, 0x7, 0x8, 0x9, 0xa, 0xb, 0xc, 0xd,
}

// HandCards 手牌
type HandCards struct {
	Cards       []byte // 手牌
	UserID      int64  // 持有这幅手牌的用户ID
	WeightValue byte   // 手牌权重值，用与比较大小
	CardsType   int32  // 牌型
}

// SolutionCards 代表游戏场景中，出牌的潜在方案。
type SolutionCards struct {
	// Cards 以字节格式表示一手牌。
	Cards []byte

	// CardsType代表手牌的类型。
	CardsType msg.CardsType

	// PutScore 表示手中牌的分数。
	PutScore int
}

// landlord 2一副扑克牌
// InitPoker 一副新牌
func (gamePoker *GamePoker) InitPoker() {

	for _, v := range Deck {
		gamePoker.Cards = append(gamePoker.Cards, v)
	}
}

// landlord 2一副扑克牌
// ShuffleCards 洗牌
func (gamePoker *GamePoker) ShuffleCards() {
	rand.Shuffle(len(gamePoker.Cards), func(i, j int) {
		gamePoker.Cards[i], gamePoker.Cards[j] = gamePoker.Cards[j], gamePoker.Cards[i]
	})
}

// DrawCard 从牌堆中抽取指定数量的牌。
// 参数：
// count：从牌堆中抽出的牌数。
// 返回：
// 卡：包含抽出的卡的字节数组。
func (gamePoker *GamePoker) DrawCard(count int) (cards []byte) {

	length := len(gamePoker.Cards)

	cards = append(cards, gamePoker.Cards[length-count:]...)

	gamePoker.Cards = gamePoker.Cards[:(length - count)]

	return
}

// GetCardValueAndColor 获取一张牌的牌值和花色
func GetCardValueAndColor(value byte) (cardValue, cardColor byte) {
	// 与 1111 1111 进行 & 位运算然后 右移动 4 位
	cardValue = (value & 0xff) >> 4
	// 与 1111 进行 & 位运算
	cardColor = value & 0xf
	return
}

// ReverseSortCards 手牌倒序排序
func ReverseSortCards(cards []byte) []byte {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < (len(cards) - 1 - i); j++ {
			if (cards)[j] < (cards)[j+1] {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}
	return cards
}

// PositiveSortCards 按升序对一手牌进行排序。
// 它使用冒泡排序算法来比较相邻的卡片，并在必要时交换它们。
// 然后返回排序后的手牌。
func PositiveSortCards(cards []byte) []byte {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < (len(cards) - 1 - i); j++ {
			if (cards)[j] > (cards)[j+1] {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}
	return cards
}

// SortCards 将牌按照牌值进行排序
func SortCards(cards []byte) []byte {
	for i := 0; i < len(cards)-1; i++ {
		for j := 0; j < (len(cards) - 1 - i); j++ {
			value1, _ := GetCardValueAndColor(cards[j])
			value2, _ := GetCardValueAndColor(cards[j+1])

			if value1 > value2 || value1 == value2 {
				cards[j], cards[j+1] = cards[j+1], cards[j]
			}
		}
	}
	return cards
}

// InByteArr 在byte数组中找目标值
func InByteArr(v byte, arr []byte) bool {
	for _, value := range arr {
		if v == value {
			return true
		}
	}

	return false
}

// 获取最小的牌
func GetSmallestCard(cards []byte) (smallestCard byte) {

	smallestCard = cards[0]
	for _, card := range cards {
		if card < smallestCard {
			smallestCard = card
		}
	}

	return
}

// HaveBigCard 有没有2或者大小王的牌
func HaveBigCard(cards []byte) bool {
	for _, card := range cards {
		value, _ := GetCardValueAndColor(card)

		if value > 12 {
			return true
		}
	}

	return false
}

// HaveRocket 有双王
func HaveRocket(cards []byte) bool {
	var kingCount int
	for _, card := range cards {
		if card > 0xe0 {
			kingCount++
		}
	}
	if kingCount > 1 {
		return true
	}

	return false
}

// HaveKing 有王牌
func HaveKing(cards []byte) bool {
	for _, card := range cards {
		if card > 0xe0 {
			return true
		}
	}

	return false
}

// GetKingCount 获取王牌 牌数
func GetKingCount(cards []byte) (count int) {
	for _, card := range cards {
		if card > 0xe0 {
			count++
		}
	}

	return count
}

// GetValue2Count 获取牌值2 的 牌数
func GetValue2Count(cards []byte) (count int) {
	for _, card := range cards {
		value, _ := GetCardValueAndColor(card)
		if value == 0xd {
			count++
		}
	}

	return count
}

// landlord 2一副扑克牌
// ContrastCards 对比两组牌的大小 true/false 表示 能/不能 大过
// @curCards 比较牌
// @lastCards 被比较牌
func ContrastCards(curCards []byte, lastCards []byte) bool {

	// 获取当前和上次打出的牌的类型
	curType, lastType := GetCardsType(curCards), GetCardsType(lastCards)

	// 卡的类型相同
	if curType == lastType {

		// 获取新的卡组进行编码
		newCurCards, newLastCards := curCards, lastCards

		// 切换不同卡类型的大小写结构
		switch curType {

		// 四带二,四带二对子
		case msg.CardsType_QuartetWithTwo, msg.CardsType_QuartetWithTwoPair:
			newCurCards, newLastCards = GetRepeatedCards(curCards, 4), GetRepeatedCards(lastCards, 4)
			break

		// 三带一, 三代一对, 飞机带对子
		case msg.CardsType_TripletWithSingle, msg.CardsType_TripletWithPair, msg.CardsType_SerialTripletWithWing:
			newCurCards, newLastCards = GetRepeatedCards(curCards, 3), GetRepeatedCards(lastCards, 3)
			break

		// 飞机带单张
		case msg.CardsType_SerialTripletWithOne:
			newCurCards, _ = GetPlane(curCards)
			newLastCards, _ = GetPlane(lastCards)
			break
		}

		// 检查卡片的长度是否相同
		if len(newCurCards) != len(newLastCards) {
			return false
		}

		// 获取卡片编码并进行比较
		curCode, lastCode := EncodeCard(newCurCards), EncodeCard(newLastCards)

		if curCode > lastCode {
			return true
		}
	} else {
		// 卡牌类型不同
		// 火箭 > 炸弹 > 其他牌
		if curType == msg.CardsType_Rocket || (curType == msg.CardsType_Bomb && lastType < msg.CardsType_Bomb) {
			return true
		}
	}

	// 在所有其他情况下返回 false
	return false
}

// EncodeCard 对手牌进行编码
func EncodeCard(cards []byte) (cardEncode int) {

	// 倒叙排序
	cards = ReverseSortCards(cards)

	for i, card := range cards {
		cardEncode |= (int(card) >> 4) << uint((len(cards)-i-1)*4)
	}
	return
}

// GetRepeatedCards 获取重复次数大于等于目标次数的牌组
func GetRepeatedCards(inCards []byte, repeatCount int) (outCards []byte) {

	for i := 0; i < len(inCards); i++ {

		// 累加器
		counter := 0

		// 值相同的牌组
		repeatCards := []byte{inCards[i]}

		valueI, _ := GetCardValueAndColor(inCards[i])

		for j := i + 1; j < len(inCards); j++ {

			valueJ, _ := GetCardValueAndColor(inCards[j])
			if valueI == valueJ {
				counter++
				repeatCards = append(repeatCards, inCards[j])
			}
		}

		if counter == repeatCount-1 {
			outCards = append(outCards, repeatCards...)
		}

	}
	return
}

// GetRepeatedValueArr 获取有重复次数大于等于目标次数的 值数组
// eg {♥3, ♠3, ♦3} => {3}
func GetRepeatedValueArr(inCards []byte, repeatCount int) (repeatValueArr []byte) {

	for _, valueI := range Unit {

		// 累加器
		counter := 0

		for _, card := range inCards {

			valueJ, _ := GetCardValueAndColor(card)
			if valueI == valueJ {
				counter++
			}

			if counter >= repeatCount {
				repeatValueArr = append(repeatValueArr, valueI)
				break
			}
		}
	}

	return
}

// GetAllValue 获取所有的值
func GetAllValue(cards []byte) (values []byte) {
	for _, card := range cards {
		v, _ := GetCardValueAndColor(card)
		values = append(values, v)
	}
	return
}

// 查询重复牌个数
// repeatedArr
// index:0 重复一次(单张牌)的牌
// index:1 重复二次(对牌)的牌
// index:2 重复三次(三张)的牌
// index:3 重复四次(炸弹)的牌
func CheckRepeatedCards(cards []byte) (repeatedArr [4][]byte) {
	notRepeatedCards := []byte{}
	for i := 0; i < len(cards); i++ {
		repeatedCount := 0
		value1, _ := GetCardValueAndColor(cards[i])

		// 防止循环检测
		isRepeat := false
		for _, cardValue := range notRepeatedCards {
			if value1 == cardValue {
				isRepeat = true
			}
		}
		if isRepeat {
			continue
		} else {
			notRepeatedCards = append(notRepeatedCards, value1)
		}

		for j := 0; j < len(cards); j++ {
			value2, _ := GetCardValueAndColor(cards[j])

			if value1 == value2 {
				repeatedCount++
			}
		}

		if repeatedCount < 1 || repeatedCount > 4 {
			log.Errorf("错误的重复牌个数 %d , 手牌: %v", repeatedCount, cards)
			return
		}
		repeatedArr[repeatedCount-1] = append(repeatedArr[repeatedCount-1], value1)
	}
	return
}

// CardsToString 牌组转字符串
func CardsToString(cards []byte) (cardsStr string) {
	for key, card := range cards {

		value, _ := GetCardValueAndColor(card)

		switch value {
		case 0x1:
			cardsStr += "3"
			break
		case 0x2:
			cardsStr += "4"
			break
		case 0x3:
			cardsStr += "5"
			break
		case 0x4:
			cardsStr += "6"
			break
		case 0x5:
			cardsStr += "7"
			break
		case 0x6:
			cardsStr += "8"
			break
		case 0x7:
			cardsStr += "9"
			break
		case 0x8:
			cardsStr += "10"
			break
		case 0x9:
			cardsStr += "J"
			break
		case 0xa:
			cardsStr += "Q"
			break
		case 0xb:
			cardsStr += "K"
			break
		case 0xc:
			cardsStr += "A"
			break
		case 0xd:
			cardsStr += "2"
			break
		case 0xe:
			cardsStr += "小王"
			break
		case 0xf:
			cardsStr += "大王"
			break
		}

		if key == len(cards)-1 {
			continue
		}
		cardsStr += "/"
	}

	return
}

// landlord 1从扑克牌开始
// TransformCards 转译牌
// 将字节切片表示的牌转译为字符串表示的牌
// 输入参数cards是一个字节切片，其中每个字节表示一张牌
// 返回值transCards是一个字符串切片，表示转译后的牌
// 转译规则如下：
// - 牌值：
//   - 0x1 -> "3", 0x2 -> "4", 0x3 -> "5", 0x4 -> "6", 0x5 -> "7",
//   - 0x6 -> "8", 0x7 -> "9", 0x8 -> "10", 0x9 -> "J", 0xa -> "Q",
//   - 0xb -> "K", 0xc -> "A", 0xd -> "2", 0xe -> "小王", 0xf -> "大王"
//
// - 花色：
//   - 0x1 -> "♦️", 0x2 -> "♣️", 0x3 -> "♥️", 0x4 -> "♠️"
//
// - 如果牌值大于0xd，则将花色设为空字符串
// 注意：函数依赖于GetCardValueAndColor函数来解析牌的牌值和花色
func TransformCards(cards []byte) (transCards []string) {

	for _, v := range cards {
		var decodeValue, decodeColor string

		value, color := GetCardValueAndColor(v)

		switch value {
		case 0x1:
			decodeValue = "3"
			break
		case 0x2:
			decodeValue = "4"
			break
		case 0x3:
			decodeValue = "5"
			break
		case 0x4:
			decodeValue = "6"
			break
		case 0x5:
			decodeValue = "7"
			break
		case 0x6:
			decodeValue = "8"
			break
		case 0x7:
			decodeValue = "9"
			break
		case 0x8:
			decodeValue = "10"
			break
		case 0x9:
			decodeValue = "J"
			break
		case 0xa:
			decodeValue = "Q"
			break
		case 0xb:
			decodeValue = "K"
			break
		case 0xc:
			decodeValue = "A"
			break
		case 0xd:
			decodeValue = "2"
			break
		case 0xe:
			decodeValue = "小王"
			break
		case 0xf:
			decodeValue = "大王"
			break
		}

		switch color {
		case 0x1:
			decodeColor = "♦️"
			break
		case 0x2:
			decodeColor = "♣️"
			break
		case 0x3:
			decodeColor = "♥️"
			break
		case 0x4:
			decodeColor = "♠️"
			break
		}

		if value > 0xd {
			decodeColor = ""
		}

		transCards = append(transCards, decodeColor+decodeValue)
	}

	return
}

// CardsTypeToString 牌型转字符串
func CardsTypeToString(cardsType msg.CardsType) (cardsTypeStr string) {
	switch cardsType {
	case msg.CardsType_Normal:
		cardsTypeStr = "无牌型"
		break
	case msg.CardsType_SingleCard:
		cardsTypeStr = "单张"
		break
	case msg.CardsType_Pair:
		cardsTypeStr = "对子"
		break
	case msg.CardsType_Triplet:
		cardsTypeStr = "三条"
		break
	case msg.CardsType_TripletWithSingle:
		cardsTypeStr = "三带一"
		break
	case msg.CardsType_TripletWithPair:
		cardsTypeStr = "三带二"
		break
	case msg.CardsType_Sequence:
		cardsTypeStr = "顺子"
		break
	case msg.CardsType_SerialPair:
		cardsTypeStr = "连对"
		break
	case msg.CardsType_SerialTriplet:
		cardsTypeStr = "飞机"
		break
	case msg.CardsType_SerialTripletWithWing:
		cardsTypeStr = "飞机带翅膀"
		break
	case msg.CardsType_QuartetWithTwo:
		cardsTypeStr = "四带二"
		break
	case msg.CardsType_Bomb:
		cardsTypeStr = "炸弹"
		break
	case msg.CardsType_Rocket:
		cardsTypeStr = "火箭"
		break

	}
	return
}

type ArrList struct {
	list [][]msg.CardsType
}

func (arrlist *ArrList) getAll(arr []msg.CardsType, m int, n int) {
	if m == n {

		list := []msg.CardsType{}
		for _, cardsType := range arr {
			list = append(list, cardsType)
		}
		arrlist.list = append(arrlist.list, list)

	} else {

		for i := m; i < n; i++ {

			tamp := arr[m]
			arr[m] = arr[i]
			arr[i] = tamp

			arrlist.getAll(arr, m+1, n)
			tamp = arr[m]
			arr[m] = arr[i]
			arr[i] = tamp
		}
	}

}
