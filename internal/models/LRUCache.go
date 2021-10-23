package models

import (
	"container/list"
	"errors"
)

type LRUCache interface {
	Set(key string, value interface{}) bool
	deleteLast()
	Get(key string) interface{}
}

type Item struct {
	Key   string
	Value interface{}
}

type LRU struct {
	Capacity int
	Items    map[string]*list.Element
	Queue    *list.List
}

func CreatLRUCache(cap int) (*LRU, error) {

	if cap < 0 {
		return nil, errors.New("cap have to be > 0")
	}

	return &LRU{
		Capacity: cap,
		Items:    make(map[string]*list.Element),
		Queue:    list.New(),
	}, nil
}

func (c *LRU) Set(key string, value interface{}) bool {
	if element, exists := c.Items[key]; exists == true {
		c.Queue.MoveToFront(element)
		element.Value.(*Item).Value = value
		return true
	}

	if c.Queue.Len() == c.Capacity {
		c.deleteLast()
	}

	item := &Item{
		Key:   key,
		Value: value,
	}

	element := c.Queue.PushFront(item)
	c.Items[item.Key] = element

	return true
}

func (c *LRU) deleteLast() {
	if element := c.Queue.Back(); element != nil {
		item := c.Queue.Remove(element).(*Item)
		delete(c.Items, item.Key)
	}
}

func (c *LRU) Get(key string) interface{} {
	element, exists := c.Items[key]
	if exists == false {
		return nil
	}

	c.Queue.MoveToFront(element)
	return element.Value.(*Item).Value
}
