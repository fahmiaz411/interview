package prime

import (
	"fmt"
	"math"
	"strings"
)

func GetPrime(minB, maxB int) []int {
	lists := []int{}

	for i := minB; i < maxB; i++ {
		if i < 2 {
			continue
		}

		isPrime := true

		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime == true {
			lists = append(lists, i)
		}
	}

	return lists
}

func CalculatePrime(lists []int, min, max int) string {
	count := make([]int, len(lists))

	sum := 0

	for {
		if sum <= min {
			sum += lists[0]
			count[0]++
		} else {
			break
		}
	}

	for i := 1; i < len(lists); i++ {
		max := math.Max(float64(lists[0]), float64(lists[i]))
		min := math.Min(float64(lists[0]), float64(lists[i]))

		substract := int(math.Floor(max / min))

		sum -= lists[0] * substract
		count[0] -= substract
		count[i] += substract
	}

	text := []string{}
	total := 0

	for i, c := range count {
		fmt.Printf("Bottle %d : %d liter\n", i+1, lists[i])

		text = append(text, fmt.Sprintf(`Bottle %d = %d bottles`, i+1, c))
		total += c
	}

	text = append(text, fmt.Sprintf(`Total = %d bottles`, total))

	return strings.Join(text, ", ")

}

func Prime() {
	minBetween := 0
	maxBetween := 30

	x := 100
	xMax := 10000000

	numbers := GetPrime(minBetween, maxBetween)

	fmt.Println(CalculatePrime(numbers[0:3], x, xMax))

}