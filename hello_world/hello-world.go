package main

import "fmt" 

func add(a, b int) int {
	return a + b
}

func swap(a, b string) (string, int) {
	return a, 0
}

func some(x string) (y, z int) {
	y = 10
	z = 11
	return
}

func main() {
	 sum := 10

	 for sum < 20 {
		 fmt.Println(sum)
		 sum++
	 }
}