package main

import "fmt"

func main() {
	i := []int{3: 4, 6: 9}
	fmt.Println(i)

	var j []int
	j = i[:]
	fmt.Println("j is:", j)

	j = i[3:]
	fmt.Println(j)

	j = i[:6]
	fmt.Println(j)

	j = i[3:6]
	fmt.Println(j)

	j[2] = 7
	fmt.Println(j)

	fmt.Println("i is:", i)

	//len
	fmt.Println("len is:", len(j))
	//cap
	fmt.Println("cap is:", cap(j))

	//apend
	j = append(j, 8)
	fmt.Println(j)
	//fmt.Println(i)

	j = append(j, 6)
	fmt.Println(j)
	//fmt.Println(i)

	//copy
	s := [5]int{1, 2, 3, 4, 5}

	s1 := s[1:4]
	s2 := s[0:2]
	fmt.Println("s1 is:", s1, "s2 is:", s2)

	copy(s2, s1)
	fmt.Println("s is:", s, "s1 is:", s1, "s2 is:", s2)

	s1[0] = 9
	fmt.Println("after:", s, s1, s2)

	//copy
	s3 := s[1:3]

	s4 := make([]int, 2)
	fmt.Println("s4 is:", s4)

	copy(s4, s3)
	fmt.Println("after s4:", s4)

	s4[0] = 9
	s3[0] = 99

	fmt.Println(s3, s4)
}
