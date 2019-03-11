package main

import "fmt"

type Postcode struct {
	ip, port, id string
}

var m map[string]Postcode

func main() {
	m = make(map[string]Postcode)
	m["Jigger"] = Postcode{
		"10.0.0.1", "8080", "43325",
	}
	m["Joshua"] = Postcode{
		"10.0.1.43", "4000", "54433",
	}
	m["Arabella"] = Postcode{
		"20.32.4", "4390", "32144",
	}

	for _, v := range m {
		fmt.Println(v)
	}

}
