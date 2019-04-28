package main

import "fmt"

func helper(i int, j int, carry int) (int, int) {
	total := i + j + carry

	if total >= 10 {
		return total % 10, 1
	}

	return total, 0
}

// Add two slices
func add(a []int, b []int) (sum []int) {
	var rslt []int
	var carry int
	sizeA, sizeB := len(a), len(b)
	switch {
	case sizeA > sizeB:
		for i := range a {
			if i > sizeB-1 {
				x, y := helper(a[i], 0, carry)
				rslt = append(rslt, x)
				carry = y
			} else {
				x, y := helper(a[i], b[i], carry)
				rslt = append(rslt, x)
				carry = y
			}
		}
	case sizeB > sizeA:
		for i := range b {
			if i > sizeA-1 {
				x, y := helper(0, b[i], carry)
				rslt = append(rslt, x)
				carry = y
			} else {
				x, y := helper(a[i], b[i], carry)
				rslt = append(rslt, x)
				carry = y
			}
		}
	default:
		for i := range a {
			x, y := helper(a[i], b[i], carry)
			rslt = append(rslt, x)
			carry = y
		}
	}

	return rslt
}

func main() {
	fmt.Printf("%v\n", add([]int{2, 4, 3}, []int{5, 6, 4}))

	fmt.Printf("%v\n", add([]int{2, 4, 3}, []int{5}))

	fmt.Printf("%v\n", add([]int{3, 2, 2, 2, 5, 0, 4, 8, 6}, []int{6, 0, 4, 8, 4, 6, 7, 7, 5, 1, 4, 5, 6, 7, 8}))
}
