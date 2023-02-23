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

	// point := 10
	//
	// if point > 7 {
	// 	switch point {
	// 	case 10:
	// 		fmt.Println("perfect!")
	// 	default:
	// 		fmt.Println("nice!")
	// 	}
	// } else {
	// 	if point == 5 {
	// 		fmt.Println("not bad")
	// 	} else if point == 3 {
	// 		fmt.Println("keep trying")
	// 	} else {
	// 		fmt.Println("you can do it")
	// 		if point == 0 {
	// 			fmt.Println("try harder!")
	// 		}
	// 	}
	// }

	//	for i := 0; i < 5; i++ {
	//		fmt.Println("Angka", i)
	//	}
	//
	//	i := 0
	//	for i < 5 {
	//		fmt.Println("Angka", i)
	//		i++
	//	}
	//
	//	j := 0
	//	for {
	//		fmt.Println("Angka", j)
	//		j++
	//		if j == 5 {
	//			break
	//		}
	//	}
	//
	//	for k := 1; k <= 10; k++ {
	//		if k%2 == 1 {
	//			continue
	//		}
	//
	//		if k > 8 {
	//			break
	//		}
	//
	//		fmt.Println("Angka", k)
	//	}
	//
	//	for i := 0; i < 5; i++ {
	//		for j := i; j < 5; j++ {
	//			fmt.Print(j, " ")
	//		}
	//		fmt.Println()
	//	}
	//
	// outerloop:
	//
	//	for i := 0; i < 5; i++ {
	//		for j := 0; j < 5; j++ {
	//			if i == 3 {
	//				break outerloop
	//			}
	//			fmt.Print("matriks [", i, "][", j, "]", "\n")
	//		}
	//	}

	// var names [4]string
	// names[0] = "trafalgar"
	// names[1] = "d"
	// names[2] = "water"
	// names[3] = "law"
	//
	// fmt.Println(names)
	//
	// fruits := [4]string{"apple", "grape", "banana", "melon"}
	// fmt.Println("Jumlah element \t\t", len(fruits))
	// fmt.Println("Isi semua element \t", fruits)
	//
	// numbers := [...]int{2, 3, 2, 4, 3}
	// fmt.Println("data array \t:", numbers)
	// fmt.Println("jumlah elemen \t:", len(numbers))
	//
	// numbers1 := [2][3]int{
	// 	{3, 2, 3},
	// 	{3, 4, 5},
	// }
	// fmt.Println("numbers1:", numbers1)
	//
	// for _, fruit := range fruits {
	// 	fmt.Printf("nama buah: %s\n", fruit)
	// }

	// fruits := make([]string, 2)
	// fruits[0] = "apple"
	// fruits[1] = "mango"
	// fmt.Println(fruits)

	// fruits := []string{"apple", "grape", "banana", "melon"}
	// fmt.Println(len(fruits))
	// fmt.Println(cap(fruits))
	//
	// aFruits := fruits[0:3]
	// fmt.Println(len(aFruits))
	// fmt.Println(cap(aFruits))
	//
	// bFruits := fruits[1:4]
	// fmt.Println(len(bFruits))
	// fmt.Println(cap(bFruits))

	// aaFruits := aFruits[1:2]
	// baFruits := bFruits[0:1]
	//
	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)
	//
	// baFruits[0] = "pineapple"
	//
	// fmt.Println(fruits)
	// fmt.Println(aFruits)
	// fmt.Println(bFruits)
	// fmt.Println(aaFruits)
	// fmt.Println(baFruits)

	// cFruits := append(fruits, "papaya")
	// fmt.Println(fruits)
	// fmt.Println(cFruits)

	// fruits := []string{"apple", "grape", "banana"}
	// bFruits := fruits[0:2]
	// fmt.Println(cap(bFruits))
	// fmt.Println(len(bFruits))
	//
	// fmt.Println(fruits)
	// fmt.Println(bFruits)
	//
	// cFruits := append(bFruits, "papaya")
	//
	// fmt.Println(fruits)
	// fmt.Println(bFruits)
	// fmt.Println(cFruits)

	// dst := make([]string, 3)
	// src := []string{"watermelon", "pineapple", "apple", "orange"}
	// n := copy(dst, src)
	//
	// fmt.Println(dst)
	// fmt.Println(src)
	// fmt.Println(n)

	// var chicken map[string]int
	// chicken = map[string]int{}
	//
	// chicken["januari"] = 50
	// chicken["februari"] = 40
	//
	// fmt.Println("januari", chicken["januari"])
	// fmt.Println("mei", chicken["mei"])

	// chicken := map[string]int{
	// 	"januari":  50,
	// 	"februari": 40,
	// 	"maret":    34,
	// 	"april":    67,
	// }
	//
	// for key, val := range chicken {
	// 	fmt.Println(key, " \t", val)
	// }
	//
	// delete(chicken, "januari")
	//
	// fmt.Println(len(chicken))
	// fmt.Println(chicken)
	//
	// value, isExists := chicken["mei"]
	//
	// if isExists {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("item is not exists")
	// }

	chickens := []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}
}
