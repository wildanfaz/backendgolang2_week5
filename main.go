package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	printInvertedTriangle(5)

	fmt.Println(generatePassword("tes", "low"))
	fmt.Println(generatePassword("muhamadwildanfaz", "low"))
	fmt.Println(generatePassword("muhamadwildanfaz", "med"))
	fmt.Println(generatePassword("muhamadwildanfaz", "strong"))
	fmt.Println(generatePassword("muhamadwildanfaz", "very strong"))

	fmt.Println(countDuration(9))
	fmt.Println(countDuration(10))
	fmt.Println(countDuration(11))
	fmt.Println(countDuration(100))
}

func printInvertedTriangle(num int) {
	result := ""
	for i := num; i > 0; i-- {
		for j := 1; j < i*2; j++ {
			result += "*"
		}
		result += "\n"
		for k := num - i; k >= 0; k-- {
			result += " "
		}
	}
	fmt.Println(result)
}

func generatePassword(password string, level string) string {
	var randomInt = strconv.Itoa(rand.Intn(900) + 100)

	if len(password) >= 6 {
		if level == "low" {
			slice := make([]string, len(password))
			for i := range password {
				if rand.Intn(10)%2 == 1 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			return result

		} else if level == "med" {
			slice := make([]string, len(password))
			for i := range password {
				if rand.Intn(10)%2 == 1 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			result += randomInt

			return result

		} else if level == "strong" {
			slice := make([]string, len(password))
			for i := range password {
				if rand.Intn(10)%2 == 1 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			specialChars := "~!@$%*_&#^()"
			for i := 0; i < 3; i++ {
				result += string(specialChars[rand.Intn(len(specialChars))])
			}

			result += randomInt

			return result

		} else {
			return "invalid level, please use low, med, or strong"
		}
	} else {
		return "password too weak, at least 6 characters"
	}
}

func countDuration(n int) string {
	data := []int{2, 3, 1, 8, 5, 4, 6, 7, 1, 4, 2}
	result := ""
	for i, first := range data {
		for j, second := range data {
			if first+second == n && i != j {
				result = strconv.Itoa(first) + " and " + strconv.Itoa(second)
				break
			}
		}
	}
	if result != "" {
		return result
	} else {
		return "data not found"
	}
}
