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
	CBCMoney int
}

func (c CBCBank) CBCDeposit(m int) {
	c.CBCMoney += m
	fmt.Printf("CBCBank Deposit %v $ Balance: %v\n", m, c.CBCMoney)
}

func (c CBCBank) CBCWithdraw(m int) {
	c.CBCMoney -= m
	fmt.Printf("CBCBank Withdraw %v $ Balance: %v\n", m, c.CBCMoney)
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
	ICBCMoney int
}

func (i ICBCBank) ICBCDeposit(m int) {
	i.ICBCMoney += m
	fmt.Printf("ICBCBank Deposit %v $ Balance: %v\n", m, i.ICBCMoney)
}

func (i ICBCBank) ICBCWithdraw(m int) {
	i.ICBCMoney -= m
	fmt.Printf("ICBCBank Withdraw %v$ Balance: %v$\n", m, i.ICBCMoney)
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
