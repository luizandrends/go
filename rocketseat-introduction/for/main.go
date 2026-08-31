package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum +=1
	}

	fmt.Println(sum)

	anotherSum := 0
	for anotherSum < 20 {
		anotherSum += 20
	}

	for anotherSum := 0; anotherSum < 20; anotherSum++ {
		fmt.Println(anotherSum)
	}

	iterator := 0
	for 2 > 1 {
		if iterator <=100 {
			iterator++
		} else {
			break
		}
	}

	fmt.Println(anotherSum)

	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
