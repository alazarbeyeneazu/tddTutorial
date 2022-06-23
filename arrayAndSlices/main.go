package main

func Adder(numbers [4]int) int {
	var sum int
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}
func Sum(numbers []int) (sum int) {
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func SumAll(slice ...[]int) []int {

	var res []int
	for i := range slice {

		res = append(res, Sum(slice[i]))
	}
	return res

}
func SumAllTail(slices ...[]int) []int {
	slice := make([]int, 0)
	for _, v := range slices {
		slice = append(slice, v[len(v)-1])
	}
	return slice
}
func main() {

}
