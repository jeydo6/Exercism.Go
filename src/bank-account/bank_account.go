package account

import "sync"

type Account struct {
	sync.RWMutex
	balance  int64
	isOpened bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		balance:  amount,
		isOpened: true,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.RLock()
	balance, ok := a.balance, a.isOpened
	a.RUnlock()
	return balance, ok
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	defer a.Unlock()

	if !a.isOpened {
		return a.balance, false
	}
	if amount < 0 && a.balance+amount < 0 {
		return a.balance, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	defer a.Unlock()
	if !a.isOpened {
		return 0, false
	}

	balance := a.balance
	a.isOpened = false
	a.balance = 0

	return balance, true
}
