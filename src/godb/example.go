// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package main

import (
	"fmt"
	"godb/container/skip_list"
	"godb/util/comparator"
)

func main() {
	s := new(skip_list.SkipList)
	cmp := new(comparator.BytewiseComparator)
	s.Init(cmp)
	
	s.Print()
	
	var key string
	key	= "1"
	if s.Search(key) {
		fmt.Println("Error: key(%d) in skip list", key)
	}
	
	s.Insert("4")
	s.Insert("2")
	s.Insert("3")
	s.Insert("1")
	
	s.Insert("6")
	s.Insert("7")
	s.Insert("8")
	s.Insert("5")
	
	s.Insert("4")
	s.Insert("2")
	s.Insert("3")
	s.Insert("1")
	
	s.Print()
	
	key	= "1"
	if !s.Search(key) {
		fmt.Println("Error: key(%d) in skip list", key)
	}
	key = "4"
	if !s.Search(key) {
		fmt.Println("Error: key(%d) in skip list", key)
	}
	key = "5"
	if s.Search(key) {
		fmt.Println("Error: key(%d) not in skip list", key)
	}
	
	s.Delete("4")
	s.Print()
	key = "4"
	if s.Search(key) {
		fmt.Println("Error: key(%d) not in skip list", key)
	}
}