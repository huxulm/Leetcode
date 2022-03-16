```javascript []
var convertToBase7 = function(num) {
  if (num == 0) {
    return "0"
  }
  var sign = false
  if (num < 0) {
    sign = true
    num = Math.abs(num)
  }

  var res = []
  while (num > 0) {
    res.push(num%7)
    num = Math.floor(num/7)
  }
  if (sign) {
    res.push('-')
  }
  return res.reverse().join('')
}

console.log(convertToBase7(8))
```

```golang []

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	var sign bool
	if num < 0 {
		sign = true
		num = -num
	}

	ans := []byte{}
	for num > 0 {
		ans = append(ans, byte(num%7+'0'))
		num /= 7
	}

	if sign {
		ans = append(ans, '-')
	}

	for i, j := 0, len(ans); i < j; i, j = i+1, j-1 {
		ans[i], ans[j] = ans[j], ans[i]
	}

	return string(ans)
}
```