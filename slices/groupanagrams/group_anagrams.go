package groupanagrams

import (
	"go-interview/datastructures/maps/hashmultimaps"
	"go-interview/strings/sorts"
)

// GroupAnagrams groups anagrams
func GroupAnagrams(list []string) *hashmultimaps.HashMultiMap {
	multimap := hashmultimaps.New()
	for _, value := range list {
		sortedValue := sorts.Sort(value)
		multimap.Put(sortedValue, value)
	}
	return multimap
}
