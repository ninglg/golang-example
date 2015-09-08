package main

import "fmt"

func main() {
	i := [...]int{3: 4, 6: 9}
	fmt.Println(i)

	var j []int
	j = i[:]
	fmt.Println(j)

	j = i[3:]
	fmt.Println(j)

	j = i[:6]
	fmt.Println(j)

	j = i[3:6]
	fmt.Println(j)

	j[2] = 7
	fmt.Println(j)
	fmt.Println(i)

	fmt.Println("len is:", len(j))
	fmt.Println("cap is:", cap(j))

	j = append(j, 8)
	fmt.Println(j)
	fmt.Println(i)

	j = append(j, 6)
	fmt.Println(j)
	fmt.Println(i)
}
