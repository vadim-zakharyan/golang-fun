package main

import (
	"math/rand"
	"time"
)

type BaseList struct {
	m map[int]interface{}
}

func (t *BaseList) Create(e interface{}) int {
	t.init()
	id := t.getNewId()
	t.m[id] = e
	return id
}
func (t *BaseList) Update(id int, e interface{}) {
	t.m[id] = e
}
func (t *BaseList) Get(id int) interface{} {
	return t.m[id]
}
func (t *BaseList) GetAll() interface{} {
	return t.m
}
func (t *BaseList) Delete(id int) {
	delete(t.m, id)
}
func (t *BaseList) init() {
	if t.m == nil {
		t.m = make(map[int]interface{})
	}
}
func (t *BaseList) getNewId() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(100)
}
