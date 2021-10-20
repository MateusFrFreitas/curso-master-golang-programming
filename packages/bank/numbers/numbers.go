package numbers

func Even(n uint) bool {
	return n%2 == 0
}

func odd(n int) bool {
	return n%2 != 0
}

func isPrime(num int) bool {
	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func OddAndPrime(num int) bool {
	return odd(num) && isPrime(num)
}
