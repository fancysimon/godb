// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package skip_list

import (
	"testing"
)

func TestExample(t *testing.T) {
	s := new(SkipList)
	s.Init()
	s.Insert("4")
	s.Insert("2")
	s.Insert("3")
	s.Insert("1")
	
	var key Key
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
}