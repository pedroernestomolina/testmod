package strings

func CountEvenOdd(s string) (odd, even int) {
	odd = 0
	even = 0
	for _, v := range s {
		if v%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	return
}
