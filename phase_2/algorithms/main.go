package main

import (
	"fmt"
	"strconv"
)

func main() {
	// list := []int{3, 5, 13, 4, 11, 6}
	//list := []int{-1, -2, 0, 3, 1, 2, 5}
	//threeSum2(list, 0)

	fmt.Println(addOne([]int{1, 2, 9, 9}))

	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func threeSum(list []int, target int) (int, int, int) {
	for x, xVal := range list {
		for y, yVal := range list {
			nextPossibleIndex := y + 1

			if nextPossibleIndex == len(list) {
				nextPossibleIndex = (nextPossibleIndex + 1) % len(list)
			}

			if xVal+yVal+list[nextPossibleIndex] == 18 {
				return x, y, nextPossibleIndex
			}
		}
	}
	return 0, 0, 0
}

func threeSum2(list []int, target int) {
	hash := make(map[int]string)

	for i, val := range list {
		nextIndice := i + 1
		if nextIndice < len(list) {
			hash[val+list[i+1]] = strconv.Itoa(i) + "," + strconv.Itoa(nextIndice)
		}
	}
	for i, v := range list {
		wantedSum := target - v
		if val, ok := hash[wantedSum]; ok {
			fmt.Printf("(%v,%v)", i, val)
		}
	}
}

func addOne(givenArray []int) []int {
	carry := 1
	result := make([]int, len(givenArray))
	index := len(givenArray) - 1

	for i := index; i >= 0; i-- {
		total := givenArray[i] + carry
		if total == 10 {
			carry = 1
		} else {
			carry = 0
		}
		result[i] = total % 10
	}
	if carry == 1 {
		result = make([]int, len(givenArray)+1)
		result[0] = 1
	}
	return result
}

/* //lst = [3, 5, 13, 4, 11, 6]
lst = [-1, -2, 0, 3, 1, 2, 5]
target = 0

console.log(lst)

lstSize = lst.length

hash = {} // hash com pares e somas
for (i = 0; i < lstSize; i++) { // O(N)
  if (i <= lstSize - 2) {
    hash[lst[i] + lst[i + 1]] = i + "," + (i + 1)
  }
}

console.log(hash)

for (i = 0; i < lstSize; i++) { // O(N)
  sumIwant = target - lst[i]
  if (hash[sumIwant] != undefined && !hash[sumIwant].includes(i)) {
    console.log("achei " + i.toString() + "," + hash[sumIwant])
  }
}
/*
/* Given a list of unique integers and a target integer, find the indices of 3 elements that sum up to the target.

e.g.
lst = [3,5,13,4,11,6]
target = 18

threeSum(lst, target) should return => (0, 3, 4)

       0    1  2  3  4  5  6
lst = [-1, -2, 0, 3, 1, 2, 5]
target = 0
*/
