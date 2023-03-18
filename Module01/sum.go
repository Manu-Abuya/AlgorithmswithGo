package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	total := 0
	for _, val := range numbers {
		total = total + val
	}
	return total
}


// Recursive
if len(numbers) == {
	return 0
}
return numbers[0] + Sum(numbers[1:])