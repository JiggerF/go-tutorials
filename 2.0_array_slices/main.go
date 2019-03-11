package main

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s = s[1:4] // output: 2,3,4 (Slice)
	s = s[1:4] // output: 3,4,5 (Slice)
	s = s[:6]  // output: 3,4,5,6,7,8 (Slice)

	s = append(s, 9, 10) // 3,4,5 .... 9,10

	println("Print Slices of Slices")
	for i, val := range s {
		println(i, val)
	}

	// Contatinate 2 slices using ...
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...) // a = [1,2,3,4,5,6]
	println("Concatinate slices\n")
	for i, v := range a {
		println(i, v)
	}

	// Contatinate without explicit var dec
	z := append([]int{1, 2, 3}, []int{4, 5, 6}...)
	println("Concatinate without slice declaration")
	for i, v := range z {
		println(i, v)
	}
	// Get Slize capacity
	println("Cap size", cap(z))

	// Copy slice to another slice
	x := []int{0, 0}
	xx := copy(x, z[1:6]) // 2,3 copied. dest max cap is 2
	println("Total size copied:", xx)
	for i, v := range x {
		println("New copied slice x", i, v)
	}
}
