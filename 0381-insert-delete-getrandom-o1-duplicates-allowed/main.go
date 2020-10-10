package _381_insert_delete_getrandom_o1_duplicates_allowedimport "math/rand"type RandomizedCollection struct {	dict   map[int]map[int]bool	values []int}/** Initialize your data structure here. */func Constructor() RandomizedCollection {	return RandomizedCollection{		dict: map[int]map[int]bool{},	}}/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */func (this *RandomizedCollection) Insert(val int) bool {	this.values = append(this.values, val)	if _, ok := this.dict[val]; !ok {		this.dict[val] = map[int]bool{len(this.values) - 1: true}		return true	}	this.dict[val][len(this.values)-1] = true	return false}/** Removes a value from the collection. Returns true if the collection contained the specified element. */func (this *RandomizedCollection) Remove(val int) bool {	if len(this.dict[val]) == 0 {		return false	}	// 随机选择一个待删除下标	removedIndex := -1	for removedIndex = range this.dict[val] {		break	}	// 与队尾元素交换	this.values[removedIndex] = this.values[len(this.values)-1]	// 删除val与removedIndex组合	delete(this.dict[val], removedIndex)	this.dict[this.values[removedIndex]][removedIndex] = true	// 删除原队尾元素与下标组合	delete(this.dict[this.values[removedIndex]], len(this.values)-1)	// 删除队尾元素	this.values = this.values[:len(this.values)-1]	return true}/** Get a random element from the collection. */func (this *RandomizedCollection) GetRandom() int {	if len(this.values) > 0 {		index := rand.Intn(len(this.values))		return this.values[index]	}	return -1}/** * Your RandomizedCollection object will be instantiated and called as such: * obj := Constructor(); * param_1 := obj.Insert(val); * param_2 := obj.Remove(val); * param_3 := obj.GetRandom(); */