package _2_strategy

import (
	"fmt"
	"testing"
)

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}
type PaymentContext struct {
	Name, CardId string
	Money        int
}

func NewPayment(name string, cardId string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardId: cardId,
			Money:  money,
		},
		strategy: strategy,
	}
}

// 支付策略接口
type PaymentStrategy interface {
	Pay(*PaymentContext)
}

// 支付策略1 现金支付
type Cash struct{}

func (c Cash) Pay(context *PaymentContext) {
	fmt.Printf("%s支付了%d元 by cash \n", context.Name, context.Money)
}

// 支付策略2 刷卡
type Bank struct{}

func (b Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("%s支付了%d元 by bank \n", ctx.Name, ctx.Money)
}
func TestPay(t *testing.T) {
	pay := NewPayment("张三", "123456789", 100, Bank{})
	pay.strategy.Pay(pay.context)
	pay = NewPayment("张三", "123456789", 100, Cash{})
	pay.strategy.Pay(pay.context)
}
