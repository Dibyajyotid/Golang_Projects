package main

import "fmt"

func main() {
	fmt.Println(greatestCommonDivisor(2, 4))
}

func greatestCommonDivisor(a int, b int) int {
	minimum := min(a, b)

	for i := minimum; i > 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}

	return 1
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//using euclidean algorithm
// func greatestCommonDivisor(a int, b int) int {
//     if b == 0 {
//         return a
//     }
//     return greatestCommonDivisor(b, a%b)
// }
