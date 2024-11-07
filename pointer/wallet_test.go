package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("A Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertError(t, err, "cannot withdraw, insufficient funds")
		//因為餘額不足，所以不會扣款，所以餘額應該是 20
		assertBalance(t, wallet, startingBalance)
	})

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposite(BitCoin(10))
		assertBalance(t, wallet, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: BitCoin(20)}
		err := wallet.Withdraw(BitCoin(10))

		assertNonError(t, err)
		assertBalance(t, wallet, BitCoin(30))
	})

}

func assertNonError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want BitCoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

// 這個輔助函式假定是有錯誤的
func assertError(t testing.TB, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	//如果 got == null，再去執行 .Error() 會出現 panic
	//所以上面的 got == nil 必須 Fatal
	if got.Error() != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}
