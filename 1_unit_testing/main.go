package main

func main() {
	bar()
}

func bar() string {
	return "bar"
}

// call foo(name) returns known lastname
func foo(name string) string {
	if name == "jigger" {
		return "fantonial"
	} else if name == "bryan" {
		return "adams"
	}
	return "default"
}
