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
				if rand.Int()%2 == 1 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			return result

		} else if level == "med" {
			password += randomInt

			slice := make([]string, len(password))
			for i := range password {
				if rand.Int()%2 == 1 {
					slice[i] = strings.ToLower(string(password[i]))
				} else {
					slice[i] = strings.ToUpper(string(password[i]))
				}
			}
			result := strings.Join(slice, "")

			return result

		} else if level == "strong" {
			slice := make([]string, len(password))
			for i := range password {
				if rand.Int()%2 == 1 {
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
			return "invalid level, please use low, medium, or strong"
		}
	} else {
		return "password too weak, at least 6 characters"
	}
}

func countDuration(n int) string {
	data := []int{2, 3, 1, 8, 5, 4, 6, 7}
	result := ""
	for i := 0; i < 1000; i++ {
		x := data[rand.Intn(len(data))]
		y := data[rand.Intn(len(data))]
		if x+y == n {
			result = strconv.Itoa(x) + " and " + strconv.Itoa(y)
			break
		}
	}
	if result != "" {
		return result
	} else {
		return "data not found"
	}
}
