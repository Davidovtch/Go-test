package arrsli

func Sum(numbers []int) int {
	sum := 0
	/*when using range to iterate over arrays,slices and maps it
	returns two values:the first is the index/key the second is the value
	if you only need the first value a single variable works but if you only
	need the second make sure to use _(blank identifier) to invalidate the first
	and only get the second*/
	for _, i := range numbers {
		sum += i
	}
	return sum
}

// variadic function takes a variable number of arguments
func SumAll(numbersSlice ...[]int) []int {
	length := len(numbersSlice)
	/*when using make to create a slice you NEED to pass a length and optionally
	a capacity,length values are initialized with the zeroed value of the type*/
	sums := make([]int, length)

	for i, numbers := range numbersSlice {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllv2(numbersSlice ...[]int) []int {
	var sums []int
	for _, number := range numbersSlice {
		sums = append(sums, Sum(number))
	}
	return sums
}

func SumAllTails(numbersSlice ...[]int) []int {
	var sums []int
	for _, number := range numbersSlice {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}

	}
	return sums
}
