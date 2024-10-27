package arrayslices

func Sum(s [5]int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}
