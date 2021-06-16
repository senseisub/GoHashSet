package main

import  (
	"fmt"
)

var exists = struct{}{}

type HashSet struct {
	internalMap map[interface{}] struct{}
}

func (set *HashSet) Insert(word interface{}) {
	if set.internalMap == nil {
		set.internalMap = make(map[interface{}] struct{})
	}
	set.internalMap[word] = exists
}

func (set *HashSet) Exists(word interface{}) bool{
	if set.internalMap == nil {
		return false
	}
	_, has := set.internalMap[word]
	return has
}

func (set *HashSet) Get(word interface{}) interface{}{
	if set.internalMap == nil {
		return false
	}
	_, has := set.internalMap[word]
	if has {
		return set.internalMap[word]
	}
	return nil
}

func (set *HashSet) Remove(word interface{}) {
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
	set.Insert("new")
	fmt.Println(set.Exists("new"))
}