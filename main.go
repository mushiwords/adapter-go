package main

import (
	"./bank"
	"fmt"
)

func main() {
	cbc := bank.CBCBank{CBCMoney: 100} // 初始CBC 100$
	cbcAdapter := bank.CBCBankAdapter{MyBank: cbc}
	fmt.Printf("CBCBank Init Money Balance: %v$\n", cbc.CBCMoney)

	myDeposit(cbcAdapter, 10)  // CBC 存 10$
	myWithdraw(cbcAdapter, 20) // CBC 取 20$

	fmt.Println()
	icbc := bank.ICBCBank{ICBCMoney: 200} // 初始ICBC 200$
	icbcAdapter := bank.ICBCBankAdapter{MyBank: icbc}
	fmt.Printf("ICBCBank Init Money Balance: %v$\n", icbc.ICBCMoney)

	myDeposit(icbcAdapter, 20)  // ICBC 存 20$
	myWithdraw(icbcAdapter, 50) // ICBC 取 50$
}

func myDeposit(b bank.Bank, m int) {
	b.Deposit(m)
}

func myWithdraw(b bank.Bank, m int) {
	b.Withdraw(m)
}
