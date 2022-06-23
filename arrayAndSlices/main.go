package main

func Adder(numbers [4]int) int {
	var sum int
	for i := 0; i < 4; i++ {
		sum += numbers[i]
	}
	return sum
}

func main() {

}
