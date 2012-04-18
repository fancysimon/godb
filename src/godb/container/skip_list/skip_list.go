// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package skip_list

import (
	"fmt"
	"math/rand"
	"time"
	"godb/util/comparator"
)

const (
	kMaxHeight = 12
	kBranching = 2 	/* 4 */
)

type Node struct {
	key string
	next []*Node
}

type SkipList struct {
	head *Node
	height int
	cmp comparator.Comparator
}

func init() {
	t := time.Now()
	second := t.UnixNano()
	println("DEBUG: Seed:", second)
	rand.Seed(second)
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
func (s *SkipList) Init(cmp comparator.Comparator) {
	s.height = kMaxHeight
	s.head = NewNode(s.height)
	s.cmp = cmp
}

// Search implication
func (s *SkipList) searchImpl(key string) (bool, []*Node) {
	var update = make([]*Node, s.height)
	var node = s.head
	for i := s.height - 1; i >= 0; i-- {
		//for node.next[i] != nil && node.next[i].key.Compare(key) == true {
		for node.next[i] != nil && s.cmp.Compare(node.next[i].key, key) < 0 {
			node = node.next[i]
		}
		update[i] = node
	}
	node = node.next[0]
	if node != nil && node.key == key {
		return true, update
	}
	return false, update
}

// Search key
func (s *SkipList) Search(key string) bool {
	has_key, _ := s.searchImpl(key)
	return has_key
}

func RandomHeight() int {
	height := 1
	for rand.Intn(kBranching) == 0 && height < kMaxHeight {
		height++
	}
	println("DEBUG: Random height:", height)
	return height
}

// Insert key to skip list
func (s *SkipList) Insert(key string) {
	has_key, update := s.searchImpl(key)
	if !has_key {
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

// Delete key in skip list
func (s *SkipList) Delete(key string) {
	has_key, update := s.searchImpl(key)
	if has_key {
		node := update[0].next[0]
		for i := 0; i < len(node.next); i++ {
			update[i].next[i] = node.next[i]
		}
	}
}

func (s *SkipList) Print() {
	var node = s.head
	for node != nil {
		if node == s.head {
			fmt.Print("List head: ")
		}
		fmt.Print("Key:", node.key, " height:", len(node.next))
		fmt.Print("  Next keys:")
		for _, next := range node.next {
			if next != nil {
				fmt.Print(next.key, " ")
			} else {
				fmt.Print(nil, " ")
			}
		}
		fmt.Println()
		node = node.next[0]
	}
	fmt.Println()
}
