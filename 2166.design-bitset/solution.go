package designbitset

type Bitset struct {
	size            int
	bits            []uint8
	oneCnt, zeroCnt int
	flip            bool
}

func Constructor(size int) Bitset {
	return Bitset{bits: make([]uint8, size), size: size, zeroCnt: size}
}

func (this *Bitset) Fix(idx int) {
	if this.flip {
		if this.bits[idx] == 1 {
			this.bits[idx] = 0
			this.oneCnt--
			this.zeroCnt++
		}
		return
	}
	if this.bits[idx] == 0 {
		this.bits[idx] = 1
		this.oneCnt++
		this.zeroCnt--
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.flip {
		if this.bits[idx] == 0 {
			this.bits[idx] = 1
			this.oneCnt++
			this.zeroCnt--
		}
		return
	}
	if this.bits[idx] == 1 {
		this.bits[idx] = 0
		this.oneCnt--
		this.zeroCnt++
	}
}

func (this *Bitset) Flip() {
	this.flip = !this.flip
}

func (this *Bitset) All() bool {
	if this.flip {
		return this.zeroCnt == this.size
	}
	return this.oneCnt == this.size
}

func (this *Bitset) One() bool {
	if this.flip {
		return this.zeroCnt > 0
	}
	return this.oneCnt > 0
}

func (this *Bitset) Count() int {
	if this.flip {
		return this.zeroCnt
	}
	return this.oneCnt
}

func (this *Bitset) ToString() string {
	buf := make([]byte, this.size)
	for i := range this.bits {
		if this.flip {
			if this.bits[i] == 0 {
				buf[i] = '1'
			} else {
				buf[i] = '0'
			}
		} else {
			buf[i] = '0' + this.bits[i]
		}
	}
	return string(buf)
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
