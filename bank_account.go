package account

import "sync"

const testVersion = 1

type Account struct {
	balance int64
	ok      bool
}

var mutex = &sync.RWMutex{}

func Open(initalDeposit int64) *Account {
	if initalDeposit < 0 {
		return nil
	}
	acc := new(Account)

	acc.balance = initalDeposit
	acc.ok = true
	return acc
}

func (acc *Account) Close() (int64, bool) {
	mutex.Lock()
	if !acc.ok {
		mutex.Unlock()
		return 0, false
	}
	acc.ok = false
	mutex.Unlock()
	return acc.balance, true
}

func (acc *Account) Balance() (int64, bool) {
	return acc.balance, acc.ok
}

func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	if !acc.ok {
		return acc.balance, false
	}
	mutex.Lock()
	newBalance = acc.balance + amount
	if newBalance < 0 {
		mutex.Unlock()
		return acc.balance, false
	}
	acc.balance = newBalance
	mutex.Unlock()
	return acc.balance, true
}
