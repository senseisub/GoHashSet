package main

import  (
	"fmt"
)

var exists = struct{}{}

type HashSet struct {
	internalMap map[interface{}] struct{}
}

func (set *HashSet) insert(word string) {
	if set.internalMap == nil {
		set.internalMap = make(map[interface{}] struct{})
	}
	set.internalMap[word] = exists
}

func (set *HashSet) exists(word string) bool{
	if set.internalMap == nil {
		return false
	}
	_, has := set.internalMap[word]
	return has
}

func (set *HashSet) remove(word string) {
	if set.internalMap == nil {
		return
	}
	_, has := set.internalMap[word]
	if has { 
		delete(set.internalMap, word);
	}
}

func main() {
	set := HashSet{}
	set.insert("new")
	fmt.Println(set.exists("new"))
}