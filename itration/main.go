package main

func Repeat(charactors string, n int) string {

	for i := 1; i < n; i++ {
		charactors += "a"
	}
	return charactors
}
func main() {

}
