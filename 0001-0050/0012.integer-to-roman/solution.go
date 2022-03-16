package integertoroman

// I             1
// IV						 4
// V             5
// IX						 9
// X             10
// XL						 40
// L             50
// XC						 90
// C             100
// CD						 400
// D             500
// CM						 900
// M             1000
func intToRoman(num int) (r string) {
	if num >= 1000 {
		return "M" + intToRoman(num-1000)
	} else if num >= 900 {
		return "CM" + intToRoman(num-900)
	} else if num >= 500 {
		return "D" + intToRoman(num-500)
	} else if num >= 400 {
		return "CD" + intToRoman(num-400)
	} else if num >= 100 {
		return "C" + intToRoman(num-100)
	} else if num >= 90 {
		return "XC" + intToRoman(num-90)
	} else if num >= 50 {
		return "L" + intToRoman(num-50)
	} else if num >= 40 {
		return "XL" + intToRoman(num-40)
	} else if num >= 10 {
		return "X" + intToRoman(num-10)
	} else if num >= 9 {
		return "IX" + intToRoman(num-9)
	} else if num >= 5 {
		return "V" + intToRoman(num-5)
	} else if num >= 4 {
		return "IV" + intToRoman(num-4)
	} else if num >= 1 {
		return "I" + intToRoman(num-1)
	}
	return
}

func intToRoman1(num int) (r string) {
	pairs := []struct {
		Val int
		Sig string
	}{
		{Sig: "M", Val: 1000},
		{Sig: "CM", Val: 900},
		{Sig: "D", Val: 500},
		{Sig: "CD", Val: 400},
		{Sig: "C", Val: 100},
		{Sig: "XC", Val: 90},
		{Sig: "L", Val: 50},
		{Sig: "XL", Val: 40},
		{Sig: "IX", Val: 9},
		{Sig: "X", Val: 10},
		{Sig: "V", Val: 5},
		{Sig: "IV", Val: 4},
		{Sig: "I", Val: 1},
	}

	for _, p := range pairs {
		for num >= p.Val {
			r += p.Sig
			num -= p.Val
			if num == 0 {
				return
			}
		}
	}
	return
}
