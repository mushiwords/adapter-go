# adapter-go
以银行取款/存款为例

工商银行（ICBC）与建设银行（CBC）

通用接口：
Deposit(m int)
Withdraw(m int)

----
Code:interface

type Bank interface {
    Deposit(m int)
    Withdraw(m int)
}

----
CBC接口: CBC Bank


type CBCBank struct {
    Money int
}

func (c CBCBank) CBCDeposit(m int) {
    fmt.Println("CBCBank Deposit %v $", m)
}

func (c CBCBank) CBCWithdraw(m int) {
    fmt.Println("CBCBank Withdraw %v $", m)
}

----
CBC适配器: CBC Bank Adapter
 
type CBCBankAdapter struct {
    MyBank CBCBank
}

func (a CBCBankAdapter) Deposit(m int) {
    a.MyBank.CBCDeposit(m)
}
func (a CBCBankAdapter) Withdraw(m int) {
    a.MyBank.CBCWithdraw(m)
}

-----
ICBC接口: ICBC Bank

type ICBCBank struct {
    Money int
}

func (i ICBCBank) ICBCDeposit(m int) {
    fmt.Println("ICBCBank Deposit %v $", m)
}

func (i ICBCBank) ICBCWithdraw(m int) {
    fmt.Println("ICBCBank Withdraw %v $", m)
}

----
ICBC适配器: ICBC Bank Adapter

type ICBCBankAdapter struct {
    MyBank ICBCBank
}

func (a ICBCBankAdapter) Deposit(m int) {
    a.MyBank.ICBCDeposit(m)
}
func (a ICBCBankAdapter) Withdraw(m int) {
    a.MyBank.ICBCWithdraw(m)
}

----
测试代码：

    cbc := bank.CBCBank{Money: 100} // 初始化CBC 100
    cbcAdapter := bank.CBCBankAdapter{MyBank: cbc}

    myDeposit(cbcAdapter, 10) // 存款
    myWithdraw(cbcAdapter, 20) // 取款

    icbc := bank.ICBCBank{Money: 200} // 初始化ICBC 100
    icbcAdapter := bank.ICBCBankAdapter{MyBank: icbc}

    myDeposit(icbcAdapter, 20)  // 存款
    myWithdraw(icbcAdapter, 50) // 取款


