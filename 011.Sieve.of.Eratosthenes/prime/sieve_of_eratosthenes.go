package prime

// SieveOfErastosthenes finds all prime numbers up to n using the Sieve of Erastosthenes algorithm
func SieveOfErastosthenes(n int) []int {
	notPrime := make([]bool, n)

	for i := 2; i <= n; i++ {
		for p := 2 * i; p <= n; p += i {
			notPrime[p-1] = true
		}
	}

	return extractPrimes(notPrime, n)
}

func extractPrimes(numbers []bool, n int) []int {
	var primes []int

	for i, notPrime := range numbers[1:] {
		if !notPrime {
			primes = append(primes, i+2)
		}
	}

	return primes
}
