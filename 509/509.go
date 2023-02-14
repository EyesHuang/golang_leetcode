package fibonacci_number

func FibBruteforce(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return FibBruteforce(n-1) + FibBruteforce(n-2)
}
