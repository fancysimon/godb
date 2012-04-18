// Copyright (c) 2012 The GoDB Authors. All rights reserved.
// Author: fancysimon@gmail.com

package comparator

import (
	"testing"
)

func TestCompare(t *testing.T) {
	cmp := new(BytewiseComparator)
	if cmp.Compare("123", "456") >= 0 {
		t.Errorf("123 < 456")
	}
	if cmp.Compare("123", "123") != 0 {
		t.Errorf("123 == 123")
	}
	if cmp.Compare("456", "123") <= 0 {
		t.Errorf("456 > 123")
	}
}

func TestName(t *testing.T) {
	cmp := new(BytewiseComparator)
	if cmp.Name() != "godb.BytewiseComparator" {
		t.Errorf("name != godb.BytewiseComparator")
	}
}