package domain

type DataPayment struct {
	Number    string
	CVV       int
	Quota     int
	Email     string
	Bank      string
	MethodPay string
}

func MethodPayFactory(dataPayment DataPayment) MethodPay {
	if dataPayment.MethodPay == "TC" {
		return NewCreditCard(dataPayment.Number, dataPayment.CVV, dataPayment.Quota)
	}

	return NewPSE(dataPayment.Email, dataPayment.Bank)
}
