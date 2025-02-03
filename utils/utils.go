package utils

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

// checks if n number is prime
func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5

	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

// return the sum of all digitsof a number
func DigitSum(n int) int {
	sum := 0

	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// isPerfectNumber checks if a number is a perfect number
func IsPerfectNumber(n int) bool {
	if n < 2 {
		return false
	}

	sum := 1 // Start with 1 since 1 is a divisor of every number
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i { // Avoid adding the same divisor twice
				sum += n / i
			}
		}
	}
	return sum == n
}

// isArmstrong checks if a number is an Armstrong number
func IsArmstrong(n int) bool {
	original := n
	sum := 0
	numDigits := int(math.Log10(float64(n))) + 1 // Count number of digits

	for temp := n; temp > 0; temp /= 10 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits))) // Raise to power of numDigits
	}

	return sum == original
}

// this gets fun fact from numbersapi.com
func GetFunFact(n int) *string {
	res, err := http.Get(fmt.Sprintf("http://numbersapi.com/%d/math", n))
	if err != nil {
		log.Println("error getting fun fact", err)
		return nil
	}
	defer res.Body.Close()

	byt, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil
	}
	text := string(byt)
	return &text
}
