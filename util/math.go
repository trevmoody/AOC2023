package util

// https://en.wikipedia.org/wiki/Least_common_multiple
func GetLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		lcm = lcm * (num2 / GCD(num1, num2))
	}
	return lcm
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
// https://en.wikipedia.org/wiki/Euclidean_algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
