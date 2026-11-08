package main

import "fmt"

func GuessMyNumber(game Game) string {
	left, right := 1, 360

	for left <= right {
		mid := left + (right-left)/2

		result := game.CheckNumber(mid)

		switch result {
		case "CORRECT":
			return fmt.Sprintf("Your Number was %d", mid)

		case "My Number is Greater":
			left = mid + 1

		case "My Number is Lower":
			right = mid - 1
		}
	}

	return ""
}
