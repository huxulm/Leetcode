package simplebanksystem

type Bank struct {
	balance []int64
	n       int
}

func Constructor(balance []int64) Bank {
	return Bank{balance: balance, n: len(balance)}
}

func (this *Bank) Transfer(a1 int, a2 int, money int64) (r bool) {
	r = a1 >= 1 && a1 <= this.n && a2 >= 1 && a2 <= this.n && this.balance[a1-1] >= money
	if r {
		this.balance[a1-1] -= money
		this.balance[a2-1] += money
	}
	return
}

func (this *Bank) Deposit(a int, money int64) (r bool) {
	r = a >= 1 && a <= this.n
	if r {
		this.balance[a-1] += money
	}
	return
}

func (this *Bank) Withdraw(a int, money int64) (r bool) {
	r = a >= 1 && a <= this.n && this.balance[a-1] >= money
	if r {
		this.balance[a-1] -= money
	}
	return
}
