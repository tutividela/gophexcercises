package main

import (
	"fmt"
	"strings"
	"testing"
)



func TestGetCharactersSet(t *testing.T) {
	set := getCharactersSet("beabeefeab")
	answers := "beaf"
	for _,v := range set {
		if !strings.Contains(answers,v) {
			t.Errorf("Expected %s , got %v",answers,set)
		}
	}
}

func TestAlternat(t *testing.T) {
	l := alternate("beabeefeab")
	if l != 5 {
		t.Errorf("Expected 5, got %d",l)
	}
}

func TestGetRestOfCharacters(t *testing.T) {
	filter := getRestOfCharacters(3,5,[]string{"a","b","c","d","e","f","g",})

	for _,v := range filter {
		if v == "d" || v == "f" {
			t.Errorf("%s is on filter",v)
		}
	}
}

func Test01(t *testing.T){
	s := []string{"a","b","c","d","e"}
	res := getRestOfCharacters(2,3,s)
	fmt.Println(res)
}