package main

func main() {
	stack := []int{1, 2, 3}
	println("Length before push:", len(stack))

	Push(&stack, 4, 5, 6, 7, 8, 9)
	println("Pusing values:")
	for i, v := range stack {
		println(i, v)
	}

	Pop(&stack, 3)
	println("Poping values:")
	for i, v := range stack {
		println(i, v)
	}
}

// LIFO Pop values based on Index
func Pop(stack *[]int, index int) {
	// Resize stack
	*stack = (*stack)[:index]
}

// Push
func Push(stack *[]int, values ...int) {
	// Push value at the end of stack
	for _, v := range values {
		*stack = append(*stack, v)
	}
}
