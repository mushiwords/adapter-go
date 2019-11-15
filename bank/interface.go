package bank

import "fmt"

/***********************************
 * interface
 **********************************/
type Bank interface {
	Deposit(m int)
	Withdraw(m int)
}

/***********************************
 * CBC Bank
 **********************************/

type CBCBank struct {
	Money int
}

func (c CBCBank) CBCDeposit(m int) {
	fmt.Println("CBCBank Deposit %v $", m)
}

func (c CBCBank) CBCWithdraw(m int) {
	fmt.Println("CBCBank Withdraw %v $", m)
}

/***********************************
 * CBC Bank Adapter
 **********************************/
type CBCBankAdapter struct {
	MyBank CBCBank
}

func (a CBCBankAdapter) Deposit(m int) {
	a.MyBank.CBCDeposit(m)
}
func (a CBCBankAdapter) Withdraw(m int) {
	a.MyBank.CBCWithdraw(m)
}

/***********************************
 * ICBC Bank
 **********************************/

type ICBCBank struct {
	Money int
}

func (i ICBCBank) ICBCDeposit(m int) {
	fmt.Println("ICBCBank Deposit %v $", m)
}

func (i ICBCBank) ICBCWithdraw(m int) {
	fmt.Println("ICBCBank Withdraw %v $", m)
}

/***********************************
 * ICBC Bank Adapter
 **********************************/
type ICBCBankAdapter struct {
	MyBank ICBCBank
}

func (a ICBCBankAdapter) Deposit(m int) {
	a.MyBank.ICBCDeposit(m)
}
func (a ICBCBankAdapter) Withdraw(m int) {
	a.MyBank.ICBCWithdraw(m)
}
