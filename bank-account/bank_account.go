package account

import (
	"sync"
)

// Account represents a bank account
type Account struct {
	balance   int64
	open      bool
	mutexLock sync.Mutex
}

// Open is used to open a new account
// Does not work with negative opening balance
func Open(openingBalance int64) *Account {
	if openingBalance < 0 {
		return nil
	}
	return &Account{balance: openingBalance, open: true}
}

// Close is used to close the account and get the money back
// Does not work for already closed account
func (acc *Account) Close() (int64, bool) {
	acc.mutexLock.Lock()
	defer acc.mutexLock.Unlock()
	if !acc.open {
		return 0, false
	}
	payout := acc.balance
	acc.balance = 0
	acc.open = false
	return payout, true
}

// Balance is used to know the account balance
// Does not work for closed account
func (acc *Account) Balance() (int64, bool) {
	acc.mutexLock.Lock()
	defer acc.mutexLock.Unlock()
	if !acc.open {
		return 0, false
	}
	payout := acc.balance
	return payout, true
}

// Deposit is used to deposit/widraw money and get new balance. Use negative amount for deposit
// Does not work for closed account or if transaction will result in negative balance
func (acc *Account) Deposit(amount int64) (int64, bool) {
	acc.mutexLock.Lock()
	defer acc.mutexLock.Unlock()
	if !acc.open {
		return 0, false
	}
	if acc.balance+amount < 0 {
		return 0, false
	}
	acc.balance += amount
	newBalance := acc.balance
	return newBalance, true
}
