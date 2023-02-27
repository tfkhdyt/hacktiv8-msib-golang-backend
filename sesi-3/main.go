package main

import (
	"fmt"
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

	// hobbies := []string{"Coding", "Listening to music"}
	// yourHobbies("Taufik", hobbies...)

	// getMinMax := func(n []int) (int, int) {
	// 	var min, max int
	// 	for i, e := range n {
	// 		switch {
	// 		case i == 0:
	// 			max, min = e, e
	// 		case e > max:
	// 			max = e
	// 		case e < min:
	// 			min = e
	// 		}
	// 	}
	//
	// 	return min, max
	// }
	//
	// numbers := []int{2, 3, 4, 3, 4, 2, 3}
	// min, max := getMinMax(numbers)
	// fmt.Printf("data: %v\nmin: %v\nmax: %v\n", numbers, min, max)

	// numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	// newNumbers := func(min int) []int {
	// 	var r []int
	// 	for _, e := range numbers {
	// 		if e < min {
	// 			continue
	// 		}
	// 		r = append(r, e)
	// 	}
	// 	return r
	// }(3)
	//
	// fmt.Println("Original numbers: ", numbers)
	// fmt.Println("Filtered numbers: ", newNumbers)

	// max := 3
	// numbers := []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	// howMany, getNumbers := findMax(numbers, max)
	// theNumbers := getNumbers()
	//
	// fmt.Println("numbers\t:", numbers)
	// fmt.Printf("find\t: %d\n\n", max)
	//
	// fmt.Println("found\t:", howMany)
	// fmt.Println("value\t:", theNumbers)

	data := []string{"wick", "jason", "ethan"}
	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	dataLength5 := filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli\t\t:", data)
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
}

func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result
}

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}
}

// func printMessage(message string, arr []string) {
// 	nameString := strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }
//
// func randomWithRange(min, max int) int {
// 	value := rand.Int()%(max-min+1) + min
// 	return value
// }
//
// func divideNumber(m, n int) {
// 	if n == 0 {
// 		fmt.Printf("Invalid divider. %d cannot divided by %d\n", m, n)
// 		return
// 	}
//
// 	res := m / n
// 	fmt.Printf("%d / %d = %d\n", m, n, res)
// }

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

// func calculate(numbers ...int) float64 {
// 	total := 0
// 	for _, number := range numbers {
// 		total += number
// 	}
//
// 	avg := float64(total) / float64(len(numbers))
// 	return avg
// }
//
// func yourHobbies(name string, hobbies ...string) {
// 	hobbiesAsString := strings.Join(hobbies, ", ")
//
// 	fmt.Printf("Hello, my name is: %s\n", name)
// 	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
// }
