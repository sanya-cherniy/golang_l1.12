package main

import "fmt"

type mySet struct {
	elements map[string]bool
}

func getSet() *mySet {
	return &mySet{make(map[string]bool)}
}

func (s *mySet) insert(elem string) {
	s.elements[elem] = true
}

func (s *mySet) contains(elem string) bool {
	if s.elements[elem] {
		return true
	}
	return false
}

func main() {
	strs := []string{"cat", "cat", "dog", "cat", "tree"}
	set := getSet()

	for i := range strs {
		set.insert(strs[i])
	}

	if set.contains("dog") {
		fmt.Println("dog is contains")
	}

	fmt.Println(set.elements)
}
