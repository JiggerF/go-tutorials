package main

import "testing"

func TestBar1(t *testing.T) {
	result := bar()
	if result != "bar" {
		t.Errorf("expecting bar, got %s", result)
	}
}

func TestFoo(t *testing.T) {
	result := foo("jigger")
	if result != "fantonial" {
		t.Errorf("expecting fantonial, got %s", result)
	}

	result = foo("bryan")
	if result != "adams" {
		t.Errorf("expecting adams, got %s", result)
	}

	result = foo("nothing")
	if result != "default" {
		t.Errorf("expecting default, got %s", result)
	}

	result = foo("just random 123123 @#$@#$")
	if result != "default" {
		t.Errorf("expecting default, got %s", result)
	}

	result = foo("")
	if result != "default" {
		t.Errorf("expecting default, got %s", result)
	}
}
