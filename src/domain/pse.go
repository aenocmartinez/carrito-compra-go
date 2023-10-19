package domain

import "fmt"

type PSE struct {
	email string
	bank  string
}

func NewPSE(email, bank string) *PSE {
	return &PSE{
		email: email,
		bank:  bank,
	}
}

func (p *PSE) Pay(total int) {
	fmt.Println("PAGO POR PSE")
	fmt.Println("Total: ", total)
	fmt.Println("Email: ", p.email)
	fmt.Println("Banco: ", p.bank)
}
