package domain

type CreditCard struct {
	number string
	cvv    int
	quota  int
}

func (cc *CreditCard) Number() string {
	return cc.number
}
