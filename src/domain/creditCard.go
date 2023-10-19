package domain

import (
	"fmt"
)

type CreditCard struct {
	number string
	cvv    int
	quota  int
}

func NewCreditCard(number string, cvv, quota int) *CreditCard {
	return &CreditCard{
		number: number,
		cvv:    cvv,
		quota:  quota,
	}
}

func (cc *CreditCard) Number() string {
	return cc.number
}

func (cc *CreditCard) Cvv() int {
	return cc.cvv
}

func (cc *CreditCard) Quota() int {
	return cc.quota
}

func (cc *CreditCard) Pay(total int) {
	fmt.Println("PAGO CON TARJETA CRÉDITO")
	fmt.Println("Total: ", total)
	fmt.Println("Cutoas: ", cc.quota)
}
