package arrayslices

func Sum(s []int) int {
	var sum int
	for _, v := range s {
		sum += v
	}
	return sum
}

//func SumAll(s ...[]int) []int {
//	var sums []int
//	for _, v := range s {
//		sums = append(sums, Sum(v))
//	}
//	return sums
//}

func SumAllTails(s ...[]int) []int {
	var sums []int
	for _, v := range s {
		if len(v) == 0 {
			sums = append(sums, 0)
			continue
		}
		tail := v[1:]
		sums = append(sums, Sum(tail))
	}
	return sums
}
