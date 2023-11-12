package main

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenOrOdd(s)
}

func evenOrOdd(s []int) []string {
	var results []string
	for _, v := range s {
		if v%2 == 0 {
			results = append(results, "even")
		} else {
			results = append(results, "odd")
		}
	}
	return results
}