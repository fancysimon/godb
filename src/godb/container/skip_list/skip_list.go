// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package skip_list

import (
	"fmt"
	"math/rand"
)

const (
	kMaxHeight = 12
	kBranching = 4
)

type Key string

func (key Key) Compare(other Key) bool {
	return key < other
}

type Node struct {
	key Key
	next []*Node
}

type SkipList struct {
	head *Node
	height int
}

func NewNode(height int) *Node {
	var node = new(Node)
	node.next = make([]*Node, height)
	for i := 0; i < height; i++ {
		node.next[i] = nil
	}
	return node
}

// Init skip list
func (s *SkipList) Init() {
	s.height = kMaxHeight
	s.head = NewNode(s.height)
}

// Search key
func (s *SkipList) Search(key Key) bool {
	var node = s.head
	for i := s.height - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].key.Compare(key) == true {
			node = node.next[i]
		}
	}
	node = node.next[0]
	if node != nil && node.key == key {
		return true
	}
	return false
}

func RandomHeight() int {
	height := 1
	for rand.Intn(kBranching) == 0 && height < kMaxHeight {
		height++
	}
	return height
}

// Insert key to skip list
func (s *SkipList) Insert(key Key) {
	var update = make([]*Node, s.height)
	var node = s.head
	for i := s.height - 1; i >= 0; i-- {
		for node.next[i] != nil && node.next[i].key.Compare(key) {
			node = node.next[i]
		}
		update[i] = node
	}

	if node == s.head || node.key != key {
		height := RandomHeight()
		if height > s.height {
			for i := s.height; i < height; i++ {
				update[i] = s.head
			}
			s.height = height
		}
		new_node := NewNode(height)
		new_node.key = key
		for i := 0; i < height; i++ {
			new_node.next[i] = update[i].next[i]
			update[i].next[i] = new_node
		}
	}
}

func (s *SkipList) Print() {
	var node = s.head
	for node != nil {
		fmt.Print(node.key, " ")
		node = node.next[0]
	}
	fmt.Println()
}
