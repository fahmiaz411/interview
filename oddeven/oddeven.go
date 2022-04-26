package oddeven

import "fmt"

func OddEven() {
	odd := 0
	even := 0

	for i := 3; i < 1000000; i++ {
		if i%2 == 0 {
			odd++
		}
	}

	for i := 4; i < 1000000; i++ {
		if i%2 == 1 {
			even++
		}
	}

	fmt.Println("Odd Number :", odd)
	fmt.Println("Even Number :", even)
}