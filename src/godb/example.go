// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package main

import (
	"fmt"
	"godb/container/skip_list"
)

func main() {
	s := new(skip_list.SkipList)
	s.Init()
	s.Insert("4")
	s.Insert("2")
	s.Insert("3")
	s.Insert("1")
	
	s.Print()
	
	var key skip_list.Key
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
}