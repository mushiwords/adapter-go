package main

import "./bank"

func main() {
	cbc := bank.CBCBank{Money: 100}
	cbcAdapter := bank.CBCBankAdapter{MyBank: cbc}

	myDeposit(cbcAdapter, 10)
	myDeposit(cbcAdapter, 20)

	icbc := bank.ICBCBank{Money: 200}
	icbcAdapter := bank.ICBCBankAdapter{MyBank: icbc}

	myDeposit(icbcAdapter, 20)
	myDeposit(icbcAdapter, 50)
}

func myDeposit(b bank.Bank, m int) {
	b.Deposit(m)
}

func myWithdraw(b bank.Bank, m int) {
	b.Withdraw(m)
}
