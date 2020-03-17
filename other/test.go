package main

import (
	"fmt"
	"math/rand"
)

// Card 扑克牌类
type Card struct {
	val    int // 牌值
	flower int // 花色
	point  int // 点数
}

// 花色
const (
	flowerFANGKUAI = iota + 1 // 方块
	flowerMEIHUA              // 梅花
	flowerHONGXIN             //
	flowerHEITAO              //
	flowerKing                // 小王,大王
)

// 点数
const (
	pointA = iota + 1
	point2
	point3
	point4
	point5
	point6
	point7
	point8
	point9
	point10
	pointJ = point10
	pointQ = point10
	pointK = point10
	pointX = point10 // 小王
	pointY = point10 // 大王
)

// 点数
const (
	value2 = iota + 1
	value3
	value4
	value5
	value6
	value7
	value8
	value9
	value10
	valueJ
	valueQ
	valueK
	valueA
	valueX // 小王
	valueY // 大王
)

type Cards []*Card // 一副牌

// NewCards 新建卡牌
func NewCards() Cards {
	cards := make(Cards, 54)
	// 给牌堆里面的牌赋初始值
	for i := 0; i < 54; i++ {
		pCard := NewCard(i + 1) // 这里是从1开始
		cards[i] = pCard
	}
	// 打乱牌的顺序
	rand.Shuffle(len(cards), func(i, j int) {
		cards[i], cards[j] = cards[j], cards[i]
	})

	return cards

}

// NewCard val: 1-54
func NewCard(val int) *Card {
	if val <= 0 || val > 54 {
		return nil
	}
	ret := &Card{}
	ret.val = toCardValue(val)
	ret.flower = toCardFlower(val)
	ret.point = toCardPoint(val)
	return ret
}

func toCardValue(val int) int {
	if val == 53 {
		return valueX // 小王
	}
	if val == 54 {
		return valueY // 大王
	}
	return ((val - 1) % 13) + 1
}

// 从牌值获取花色
func toCardFlower(val int) int {
	if val == 53 {
		return flowerKing
	}
	if val == 54 {
		return flowerKing
	}
	return ((val - 1) / 13) + 1
}

// 从牌值获取点数
func toCardPoint(val int) int {
	if val == 53 {
		return pointX // 小王
	}
	if val == 54 {
		return pointY // 大王
	}
	if val%13 > 10 {
		return point10
	}
	return val % 13
}

func main() {
	for _, v := range NewCards() {
		fmt.Println(*v)
	}
	fmt.Println(len(NewCards()))
}
