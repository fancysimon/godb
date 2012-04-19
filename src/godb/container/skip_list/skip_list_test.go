// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package skip_list

import (
	"testing"
	"godb/util/comparator"
)

func TestSkipList(t *testing.T) {
	s := new(SkipList)
	cmp := new(comparator.BytewiseComparator)
	s.Init(cmp)
	s.Insert("4")
	s.Insert("2")
	s.Insert("3")
	s.Insert("1")
	
	var key string
	key	= "1"
	if !s.Search(key) {
		t.Errorf("key(%d) in skip list", key)
	}
	key = "4"
	if !s.Search(key) {
		t.Errorf("key(%d) in skip list", key)
	}
	key = "5"
	if s.Search(key) {
		t.Errorf("key(%d) not in skip list", key)
	}
	
	key = "4"
	s.Delete("4")
	if s.Search(key) {
		t.Errorf("key(%d) in skip list", key)
	}
}
