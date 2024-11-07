package pointers

import (
	"errors"
	"fmt"
)

type Wallet struct {
	balance BitCoin
}

type BitCoin int

// BitCoin implement Stringer interface
// fmt Printf will call this method to print the value
func (b BitCoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() BitCoin {
	return w.balance
}

func (w *Wallet) Deposite(amount BitCoin) {
	w.balance += amount
}

var ErrInsufficientFunds error = errors.New("cannot withdraw, insufficient funds")

// 提領錢包正常流程是回成功，如果錢包餘額不足，則回傳錯誤
// 預設是成功，所以失敗是偶然的，所以列為檢驗項目
func (w *Wallet) Withdraw(amount BitCoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return ErrInsufficientFunds
}
