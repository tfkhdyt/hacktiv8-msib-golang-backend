package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	// names := []string{"Taufik", "Hidayat"}
	// printMessage("Halo", names)

	// rand.Seed(time.Now().Unix())
	//
	// randomValue := randomWithRange(2, 10)
	// fmt.Println("Random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("Random number:", randomValue)
	// randomValue = randomWithRange(2, 10)
	// fmt.Println("Random number:", randomValue)

	// divideNumber(10, 2)
	// divideNumber(4, 0)
	// divideNumber(8, -4)

	// diameter := 15.0
	// area, circumference := calculate(diameter)
	//
	// fmt.Printf("Luas lingkaran\t\t: %.2f\n", area)
	// fmt.Printf("Keliling lingkaran\t: %.2f\n", circumference)

	// numbers := []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
	// avg := calculate(numbers...)
	// fmt.Printf("Rata-rata: %.2f\n", avg)

	hobbies := []string{"Coding", "Listening to music"}
	yourHobbies("Taufik", hobbies...)
}

func printMessage(message string, arr []string) {
	nameString := strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

func randomWithRange(min, max int) int {
	value := rand.Int()%(max-min+1) + min
	return value
}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	res := m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

// func calculate(d float64) (area float64, circumference float64) {
// 	// hitung luas
// 	area = math.Pi * math.Pow(d/2, 2)
//
// 	// hitung keliling
// 	circumference = math.Pi * d
//
// 	// kembalikan 2 nilai
// 	return
// }

func calculate(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}

	avg := float64(total) / float64(len(numbers))
	return avg
}

func yourHobbies(name string, hobbies ...string) {
	hobbiesAsString := strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}
