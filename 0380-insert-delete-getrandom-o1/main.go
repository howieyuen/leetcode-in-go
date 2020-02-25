package _380_insert_delete_getrandom_o1

import (
	"math/rand"
)

type RandomizedSet struct {
	maps map[int]bool
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{maps: map[int]bool{}}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if this.maps[val] {
		return false
	}
	this.maps[val] = true
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if this.maps[val] {
		delete(this.maps, val)
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	random := rand.Int() % len(this.maps)
	count := 0
	for k := range this.maps {
		if count == random {
			return k
		}
		count++
	}
	return -1
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
