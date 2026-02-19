// 练习3：银行账户系统
// 实现带有账户余额和转账功能的 BankAccount

package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("===== 练习3：银行账户系统 =====")

	// TODO: 创建两个账户
	acc1 := NewBankAccount("Alice", 1000)
	acc2 := NewBankAccount("Bob", 500)

	// TODO: 查看账户信息
	acc1.Info()
	acc2.Info()

	// TODO: 存款
	acc1.Deposit(500)
	acc1.Info()

	// TODO: 取款
	err := acc1.Withdraw(300)
	if err != nil { fmt.Println("取款失败:", err) }
	acc1.Info()

	// TODO: 转账
	err = acc1.Transfer(acc2, 200)
	if err != nil { fmt.Println("转账失败:", err) }

	// TODO: 查看最终余额
	acc1.Info()
	acc2.Info()
}

// BankAccount 银行账户结构体
type BankAccount struct {
	// TODO: 定义 AccountNumber（账号）、Owner（户主）、Balance（余额）字段
	AccountNumber int
	Owner         string
	Balance       float64
}

// 全局变量：用于生成唯一账号
var accountCounter = 1000

// NewBankAccount 创建新银行账户
func NewBankAccount(owner string, initialBalance float64) *BankAccount {
	// TODO: accountCounter++ 生成新账号
	accountCounter += 1
	acc := BankAccount{
		AccountNumber: accountCounter,
		Owner:         owner,
		Balance:       initialBalance,
	}
	// 返回 BankAccount 指针
	return &acc
}

// Deposit 存款
func (a *BankAccount) Deposit(amount float64) error {
	// TODO: 检查 amount > 0
	// 增加余额
	if amount > 0 {
		a.Balance += amount
		return nil
	}
	// 返回 nil 或错误
	return fmt.Errorf("存入数值不能小于等于0")
}

// Withdraw 取款
func (a *BankAccount) Withdraw(amount float64) error {
	// TODO: 检查 amount > 0
	// 检查余额是否充足
	// 减少余额
	// 返回 nil 或错误 errors.New("余额不足")
	if amount > 0 && (a.Balance-amount) >= 0 {
		a.Balance = a.Balance - amount
		return nil
	}
	return errors.New("余额不足")
}

// Transfer 转账给另一个账户
func (a *BankAccount) Transfer(target *BankAccount, amount float64) error {
	// TODO: 先从当前账户取款
	// 然后存入目标账户
	// 注意：如果取款失败，不应该执行存款
	if err := a.Withdraw(amount); err != nil {
		return err
	}

	if err := target.Deposit(amount); err != nil {
		return err
	}
	return nil

}

// GetBalance 获取当前余额
func (a BankAccount) GetBalance() float64 {
	// TODO: 返回余额
	return a.Balance
}

// Info 打印账户信息
func (a BankAccount) Info() {
	// TODO: 打印账号、户主、余额信息
	fmt.Printf("账号: %d, 户主: %s, 余额: %.2f\n", a.AccountNumber, a.Owner, a.Balance)
}
