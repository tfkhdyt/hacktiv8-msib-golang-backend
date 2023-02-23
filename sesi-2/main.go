package main

import "fmt"

func main() {
	// point := 10
	//
	// if point == 10 {
	// 	fmt.Println("lulus dengan nilai sempurna")
	// } else if point > 5 {
	// 	fmt.Println("lulus")
	// } else if point == 4 {
	// 	fmt.Println("hampir lulus")
	// } else {
	// 	fmt.Printf("tidak lulus. nilai anda %d\n", point)
	// }

	// point := 8840.0
	//
	// if percent := point / 100; percent >= 100 {
	// 	fmt.Printf("%.1f%% perfect!\n", percent)
	// } else if percent >= 70 {
	// 	fmt.Printf("%.1f%% good\n", percent)
	// } else {
	// 	fmt.Printf("%.1f%% not bad\n", percent)
	// }

	// point := 6
	//
	// switch {
	// case point == 8:
	// 	fmt.Println("perfect")
	// case (point < 8) && (point > 3):
	// 	fmt.Println("awesome")
	// 	fallthrough
	// case point < 5:
	// 	fmt.Println("you need to learn more")
	// default:
	// 	{
	// 		fmt.Println("not bad")
	// 		fmt.Println("you can be better!")
	// 	}
	// }

	point := 10

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("try harder!")
			}
		}
	}
}
