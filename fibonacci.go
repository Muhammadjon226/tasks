package tasks


func Fibonacci (limit int) []int {

	var fibonacciNums []int = []int{}

	var a, b, c int

	b = 1

	for i := 0; i < limit; i++ {

		fibonacciNums = append(fibonacciNums, c + b)
		a = c + b
		c = b
		b = a

	}
	return fibonacciNums
}